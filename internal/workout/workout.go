package workout

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
	"go.uber.org/zap"
)

// GopherLetics style of stepwise Workout
type Workout struct {
	Ty string // Type (heading, exercise, rest)
	Na string // Name
	Du int    // Duration
}

// All Parsed Workouts
type Workouts []Workout

var Wo Workouts

// All Exercises from the Free Exercise DB (public domain / Unlicense)
type Exercise struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Force            string   `json:"force"`
	Level            string   `json:"level"`
	Mechanic         string   `json:"mechanic"`
	Equipment        string   `json:"equipment"`
	PrimaryMuscles   []string `json:"primaryMuscles"`
	SecondaryMuscles []string `json:"secondaryMuscles"`
	Instructions     []string `json:"instructions"`
	Category         string   `json:"category"`
	Images           []string `json:"images"`
}

type Exercises []Exercise

var AllExercises Exercises

// FetchAllExercises loads the exercise DB from a local cache file.
// If the cache is missing, it downloads from GitHub and caches it.
func FetchAllExercises() {
	cachePath := cacheFilePath()

	// Try loading from cache first
	if data, err := os.ReadFile(cachePath); err == nil {
		if err := json.Unmarshal(data, &AllExercises); err == nil {
			zap.L().Info("loaded exercises from cache", zap.Int("count", len(AllExercises)))
			LoadOverlay()
			return
		}
	}

	// Cache miss — download from upstream
	url := "https://raw.githubusercontent.com/yuhonas/free-exercise-db/main/dist/exercises.json"
	resp, err := http.Get(url)
	if err != nil {
		zap.L().Error("failed to fetch exercises", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		zap.L().Error("failed to read response body", zap.Error(err))
		return
	}

	if err := json.Unmarshal(body, &AllExercises); err != nil {
		zap.L().Error("failed to unmarshal exercises", zap.Error(err))
		return
	}

	// Write cache
	if err := os.MkdirAll(filepath.Dir(cachePath), 0755); err == nil {
		_ = os.WriteFile(cachePath, body, 0644)
		zap.L().Info("cached exercises", zap.String("path", cachePath))
	}

	zap.L().Info("loaded exercises from upstream", zap.Int("count", len(AllExercises)))
	LoadOverlay()
}

func cacheFilePath() string {
	return filepath.Join(configdir.LocalConfig("gopherletics"), "exercises.json")
}