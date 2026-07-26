package workout

import (
	"fmt"
	"math"

	"go.uber.org/zap"
)

// ---------- Template definitions ----------

// SlotType describes how exercises are arranged within a phase.
type SlotType string

const (
	SlotInterval SlotType = "interval" // one exercise at a time, no rest within phase
	SlotCircuit  SlotType = "circuit"  // N exercises repeated in rounds, rest between rounds
	SlotPyramid  SlotType = "pyramid"  // one exercise done at increasing durations (base, 2×, 3×)
)

// PhaseDef describes one phase of a workout template.
type PhaseDef struct {
	Name     string   // display name (Warmup, Main, Finisher, Cooldown)
	Type     string   // phase type identifier
	Ratio    float64  // fraction of total workout time (must sum to ~1.0)
	SlotType SlotType // how exercises are scheduled
	Count    int      // number of unique exercises in this phase
	Sets     int      // rounds (for circuit) or pyramid step count (for pyramid)
}

// TemplateDef defines a complete workout structure.
type TemplateDef struct {
	Name        string
	Description string
	Phases      []PhaseDef
}

// Built-in templates.
var Templates = []TemplateDef{
	{
		Name: "Classic Strength",
		Description: "Warmup → circuit sets → finisher → cooldown",
		Phases: []PhaseDef{
			{Name: "Warmup", Type: "warmup", Ratio: 0.15, SlotType: SlotInterval, Count: 5, Sets: 1},
			{Name: "Main", Type: "main", Ratio: 0.60, SlotType: SlotCircuit, Count: 6, Sets: 3},
			{Name: "Finisher", Type: "finisher", Ratio: 0.10, SlotType: SlotInterval, Count: 3, Sets: 1},
			{Name: "Cooldown", Type: "cooldown", Ratio: 0.15, SlotType: SlotInterval, Count: 4, Sets: 1},
		},
	},
	{
		Name: "Pyramid Strength",
		Description: "Warmup → pyramid sets (15s→30s→45s) → finisher → cooldown",
		Phases: []PhaseDef{
			{Name: "Warmup", Type: "warmup", Ratio: 0.15, SlotType: SlotInterval, Count: 5, Sets: 1},
			{Name: "Main", Type: "main", Ratio: 0.60, SlotType: SlotPyramid, Count: 3, Sets: 3},
			{Name: "Finisher", Type: "finisher", Ratio: 0.10, SlotType: SlotInterval, Count: 3, Sets: 1},
			{Name: "Cooldown", Type: "cooldown", Ratio: 0.15, SlotType: SlotInterval, Count: 4, Sets: 1},
		},
	},
	{
		Name: "Cardio HIIT",
		Description: "Warmup → HIIT circuits → cooldown",
		Phases: []PhaseDef{
			{Name: "Warmup", Type: "warmup", Ratio: 0.15, SlotType: SlotInterval, Count: 4, Sets: 1},
			{Name: "Main", Type: "main", Ratio: 0.70, SlotType: SlotCircuit, Count: 5, Sets: 4},
			{Name: "Cooldown", Type: "cooldown", Ratio: 0.15, SlotType: SlotInterval, Count: 4, Sets: 1},
		},
	},
}

// ---------- Generation ----------

// GenerateConfig mirrors user settings for workout generation.
type GenerateConfig struct {
	Duration  int      // total workout duration in minutes
	Type      string   // strength, cardio, mixed
	Area      string   // full, upper, lower, core
	Level     string   // beginner, intermediate, expert
	Equipment []string // selected equipment types
	Template  string   // name of the template to use
}

// Generate produces a Workouts slice from the exercise pool using the given config.
func (wo *Workouts) Generate(cfg GenerateConfig) {
	Wo = Workouts{}

	// Find template
	var tmpl *TemplateDef
	for i := range Templates {
		if Templates[i].Name == cfg.Template {
			tmpl = &Templates[i]
			break
		}
	}
	if tmpl == nil {
		tmpl = &Templates[0] // fallback
	}

	totalSec := cfg.Duration * 60
	if totalSec < 60 {
		totalSec = 60
	}

	// Build exercise pool
	poolCfg := FilterConfig{
		Type:      cfg.Type,
		Area:      cfg.Area,
		Level:     cfg.Level,
		Equipment: cfg.Equipment,
	}
	pool := FilterExercises(poolCfg)

	// Pre-warmup pool for stretching phases
	stretchPool := FilterStretches(poolCfg)

	zap.L().Debug("generating workout",
		zap.Int("duration_sec", totalSec),
		zap.String("template", tmpl.Name),
		zap.Int("pool_size", len(pool)),
	)

	for _, phase := range tmpl.Phases {
		phaseSec := int(math.Round(float64(totalSec) * phase.Ratio))
		if phaseSec < 10 {
			phaseSec = 10
		}

		// Pick exercises appropriate for the phase type
		var phasePool Exercises
		if phase.Type == "cooldown" && len(stretchPool) > 0 {
			phasePool = stretchPool
		} else {
			phasePool = pool
		}

		switch phase.SlotType {
		case SlotInterval:
			wo.generateInterval(phase, phaseSec, phasePool)
		case SlotCircuit:
			wo.generateCircuit(phase, phaseSec, phasePool)
		case SlotPyramid:
			wo.generatePyramid(phase, phaseSec, phasePool)
		}
	}

	zap.L().Debug("generated workout", zap.Int("total_entries", len(Wo)))
}

// generateInterval fills a phase with one-off exercises.
func (wo *Workouts) generateInterval(phase PhaseDef, sec int, pool Exercises) {
	// Heading
	Wo = append(Wo, Workout{Ty: "heading", Na: phase.Name, Du: 0})

	if phase.Count < 1 {
		return
	}

	exercises := PickN(pool, phase.Count)
	if len(exercises) == 0 {
		return
	}

	work := sec / len(exercises)
	work = round5(work)

	for _, ex := range exercises {
		name := ex.Name
		if name == "" {
			name = fmt.Sprintf("Exercise %d", len(Wo))
		}
		Wo = append(Wo, Workout{Ty: "exercise", Na: name, Du: work})
	}
}

// generateCircuit fills a phase with circuit rounds.
func (wo *Workouts) generateCircuit(phase PhaseDef, sec int, pool Exercises) {
	Wo = append(Wo, Workout{Ty: "heading", Na: phase.Name, Du: 0})

	if phase.Count < 1 || phase.Sets < 1 {
		return
	}

	exercises := PickN(pool, phase.Count)
	if len(exercises) == 0 {
		return
	}

	// Circuit: total = sets × count × work + (sets-1) × rest
	// Solve for work and rest
	slots := phase.Sets * len(exercises)
	restSlots := phase.Sets - 1

	// Guess work duration, compute rest as leftover
	work := sec / slots
	if work > 60 {
		work = 60
	}
	work = round5(work)

	rest := 20
	// Check if rest fits; if not, reduce work
	for rest > 0 && (slots*work+restSlots*rest) > sec && work > 10 {
		work -= 5
		rest = max(10, (sec-slots*work)/max(restSlots, 1))
		rest = round5(rest)
	}
	if rest < 10 {
		rest = 10
	}

	for set := range phase.Sets {
		for _, ex := range exercises {
			name := ex.Name
			if name == "" {
				name = fmt.Sprintf("Exercise %d", len(Wo))
			}
			Wo = append(Wo, Workout{Ty: "exercise", Na: name, Du: work})
		}
		// Rest between rounds (not after the last)
		if set < phase.Sets-1 {
			Wo = append(Wo, Workout{Ty: "rest", Na: "Rest", Du: rest})
		}
	}
}

// generatePyramid fills a phase with pyramid sets (base, 2×base, 3×base).
func (wo *Workouts) generatePyramid(phase PhaseDef, sec int, pool Exercises) {
	Wo = append(Wo, Workout{Ty: "heading", Na: phase.Name, Du: 0})

	if phase.Count < 1 {
		return
	}

	exercises := PickN(pool, phase.Count)
	if len(exercises) == 0 {
		return
	}

	// Pyramid multiplier sum: 1 + 2 + 3 = 6 per exercise
	const pyramidMultiplierSum = 6
	totalUnits := phase.Count * pyramidMultiplierSum
	base := sec / totalUnits
	if base < 10 {
		base = 10
	}
	base = round5(base)

	rest := 20

	for i, ex := range exercises {
		name := ex.Name
		if name == "" {
			name = fmt.Sprintf("Exercise %d", len(Wo))
		}

		// Add transition heading between exercises in a pyramid
		if i > 0 {
			// small rest before next pyramid block
		}

		depths := []int{1, 2, 3}
		for di, mult := range depths {
			dur := base * mult
			if di == 0 {
				Wo = append(Wo, Workout{Ty: "exercise", Na: name, Du: dur})
			} else {
				Wo = append(Wo, Workout{Ty: "exercise", Na: name, Du: dur})
			}
			if di < len(depths)-1 {
				Wo = append(Wo, Workout{Ty: "rest", Na: "Rest", Du: rest})
			}
		}

		// Rest between different exercises (shorter)
		if i < len(exercises)-1 {
			Wo = append(Wo, Workout{Ty: "rest", Na: "Rest", Du: rest / 2})
		}
	}
}

func round5(n int) int {
	if n < 5 {
		return max(n, 5)
	}
	return ((n + 2) / 5) * 5
}
