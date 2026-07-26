package workout

import (
	"math/rand/v2"
	"strings"
)

// areaMuscles maps user-facing area to primary muscle groups.
var areaMuscles = map[string][]string{
	"full":  {"chest", "shoulders", "biceps", "triceps", "abdominals", "quadriceps", "hamstrings", "glutes", "calves", "lats", "traps", "lower back", "middle back", "forearms", "neck", "adductors", "abductors"},
	"upper": {"chest", "shoulders", "biceps", "triceps", "lats", "traps", "lower back", "middle back", "forearms", "neck"},
	"lower": {"quadriceps", "hamstrings", "glutes", "calves", "adductors", "abductors"},
	"core":  {"abdominals", "lower back", "neck"},
}

// FilterConfig mirrors user settings for exercise selection.
type FilterConfig struct {
	Type      string   // strength, cardio, mixed
	Area      string   // full, upper, lower, core
	Level     string   // beginner, intermediate, expert
	Equipment []string // selected equipment types; ["any"] or empty means no filter
}

// FilterExercises returns a shuffled pool of exercises matching the given criteria.
// Unmatched criteria are silently loosened (e.g. no beginner exercises → include intermediate).
func FilterExercises(cfg FilterConfig) Exercises {
	// Build inclusion sets
	areaSet := make(map[string]bool)
	if muscles, ok := areaMuscles[cfg.Area]; ok {
		for _, m := range muscles {
			areaSet[m] = true
		}
	}

	// Level fallback ladder
	levelRank := map[string]int{"beginner": 0, "intermediate": 1, "expert": 2}
	maxLevel := levelRank[cfg.Level]

	var matched Exercises
	for _, ex := range AllExercises {
		// Level: accept if equal or lower rank (loosen downward)
		if exRank, ok := levelRank[ex.Level]; !ok || exRank > maxLevel {
			continue
		}

		// Equipment: if specific list (not ["any"]), check membership
		if len(cfg.Equipment) > 0 && cfg.Equipment[0] != "any" {
			if !contains(cfg.Equipment, ex.Equipment) {
				continue
			}
		}

		// Area: at least one primary muscle must match the target area
		if cfg.Area != "" && cfg.Area != "full" && !muscleOverlaps(ex.PrimaryMuscles, areaSet) {
			continue
		}

		// Category filter for type
		if cfg.Type == "strength" && ex.Category != "strength" && ex.Category != "powerlifting" && ex.Category != "olympic weightlifting" && ex.Category != "strongman" {
			continue
		}
		if cfg.Type == "cardio" && ex.Category != "cardio" && ex.Category != "plyometrics" {
			continue
		}

		matched = append(matched, ex)
	}

	// Shuffle for variety
	rand.Shuffle(len(matched), func(i, j int) {
		matched[i], matched[j] = matched[j], matched[i]
	})

	return matched
}

// FilterStretches returns exercises suitable for cooldown/stretching.
func FilterStretches(cfg FilterConfig) Exercises {
	var matched Exercises
	levelRank := map[string]int{"beginner": 0, "intermediate": 1, "expert": 2}
	maxLevel := levelRank[cfg.Level]

	for _, ex := range AllExercises {
		if exRank, ok := levelRank[ex.Level]; !ok || exRank > maxLevel {
			continue
		}
		if ex.Category != "stretching" {
			continue
		}
		matched = append(matched, ex)
	}

	rand.Shuffle(len(matched), func(i, j int) {
		matched[i], matched[j] = matched[j], matched[i]
	})
	return matched
}

// PickN returns n random unique exercises from the pool (or as many as available).
func PickN(pool Exercises, n int) Exercises {
	if n >= len(pool) {
		return pool
	}
	return pool[:n]
}

func muscleOverlaps(muscles []string, areaSet map[string]bool) bool {
	for _, m := range muscles {
		if areaSet[strings.ToLower(m)] {
			return true
		}
	}
	return false
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
