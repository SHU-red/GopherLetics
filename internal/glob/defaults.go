package glob

import (
	"github.com/spf13/viper"
)

// Set all default config values
func conf_setDefaults() {

	// Workout
	viper.SetDefault("Conf.Workout.Duration", "30")
	viper.SetDefault("Conf.Workout.Type", "strength")
	viper.SetDefault("Conf.Workout.Area", "full")
	viper.SetDefault("Conf.Workout.Level", "beginner")

	// Audio
	viper.SetDefault("Conf.Settings.Audio.Activate", true)
	viper.SetDefault("Conf.Settings.Audio.ActivateCountdown", true)
	viper.SetDefault("Conf.Settings.Audio.ActivateExercise", true)
	viper.SetDefault("Conf.Settings.Audio.ActivatePause", true)

	// Url Trigger
	viper.SetDefault("Conf.Settings.UrlTrigger.Url", "https://www.youtube.com/results?search_query=%s")
}
