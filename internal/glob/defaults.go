package glob

import "github.com/spf13/viper"

// Set all default config values
func conf_setDefaults() {

	// Workout
	viper.SetDefault("workout.duration", "30")
	viper.SetDefault("workout.type", "strength")
	viper.SetDefault("workout.area", "full")
	viper.SetDefault("workout.level", "beginner")

	// Audio
	viper.SetDefault("settings.audio.activate", true)
	viper.SetDefault("settings.audio.activatecountdown", true)
	viper.SetDefault("settings.audio.activateexercise", true)
	viper.SetDefault("settings.audio.activatepause", true)

	// Url Trigger
	viper.SetDefault("settings.urltrigger.url", "https://www.youtube.com/results?search_query=%s")
}
