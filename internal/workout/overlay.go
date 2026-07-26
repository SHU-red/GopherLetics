package workout

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
	"go.uber.org/zap"
)

// Overlay holds user-contributed mappings and custom exercises.
// Stored as overrides.yaml in the config directory.
type Overlay struct {
	// Aliases maps sebhulse-style exercise names to upstream exercise IDs.
	// e.g. "Hand Release Push Up" → "Push-Up"
	Aliases map[string]string `json:"aliases"`

	// CustomExercises are user-defined exercises not in the upstream DB.
	CustomExercises []Exercise `json:"custom_exercises"`

	// aliasLookup is built at load time: maps alias → resolved exercise pointer
	aliasLookup map[string]*Exercise
}

var currentOverlay *Overlay

// LoadOverlay reads the overlay file from the config directory.
func LoadOverlay() *Overlay {
	confDir := configdir.LocalConfig("gopherletics")
	path := filepath.Join(confDir, "overrides.json")

	o := &Overlay{
		Aliases:         make(map[string]string),
		CustomExercises: nil,
		aliasLookup:     make(map[string]*Exercise),
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if !os.IsNotExist(err) {
			zap.L().Error("failed to read overlay", zap.Error(err))
		}
		currentOverlay = o
		return o
	}

	if err := json.Unmarshal(data, o); err != nil {
		zap.L().Error("failed to parse overlay", zap.Error(err))
		currentOverlay = o
		return o
	}

	// Build alias lookup: resolve each alias to an exercise in AllExercises
	for alias, id := range o.Aliases {
		for i := range AllExercises {
			if AllExercises[i].ID == id {
				o.aliasLookup[alias] = &AllExercises[i]
				break
			}
		}
	}

	zap.L().Info("loaded overlay", zap.Int("aliases", len(o.Aliases)), zap.Int("custom_exercises", len(o.CustomExercises)))
	currentOverlay = o
	return o
}

// SaveOverlay writes the overlay to disk.
func SaveOverlay(o *Overlay) error {
	confDir := configdir.LocalConfig("gopherletics")
	path := filepath.Join(confDir, "overrides.json")

	data, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}

	// Rebuild lookup
	o.aliasLookup = make(map[string]*Exercise)
	for alias, id := range o.Aliases {
		for i := range AllExercises {
			if AllExercises[i].ID == id {
				o.aliasLookup[alias] = &AllExercises[i]
				break
			}
		}
	}

	return nil
}

// ResolveExercise tries the alias map first, then falls back to fuzzy matching.
// Returns nil if no match is found.
func ResolveExercise(name string) *Exercise {
	// Try exact match by name first
	for i := range AllExercises {
		if AllExercises[i].Name == name {
			return &AllExercises[i]
		}
	}

	// Try alias lookup
	if currentOverlay != nil {
		if ex, ok := currentOverlay.aliasLookup[name]; ok {
			return ex
		}
	}

	// Try custom exercises
	if currentOverlay != nil {
		for i := range currentOverlay.CustomExercises {
			if currentOverlay.CustomExercises[i].Name == name {
				return &currentOverlay.CustomExercises[i]
			}
		}
	}

	return nil
}

// AddAlias records a new alias and persists the overlay.
func AddAlias(sebhulseName, exerciseID string) error {
	if currentOverlay == nil {
		LoadOverlay()
	}
	currentOverlay.Aliases[sebhulseName] = exerciseID
	return SaveOverlay(currentOverlay)
}
