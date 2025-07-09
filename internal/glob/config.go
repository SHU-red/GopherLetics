package glob

import (
	"fmt"

	"github.com/kirsle/configdir"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Config constants
const conf_filename = "conf"
const conf_extension = "toml"

type Config struct {
	Workout  Workout  `json:"workout"`
	Settings Settings `json:"settings"`
}

// Add viper package to edit code

// Declare Choices
var Choices_Type = []string{"strength", "cardio"}
var Choices_Level = []string{"beginner", "intermediate", "advanced"}
var Choices_Area = []string{"full", "upper", "lower", "core"}

// All Settings from workout popup
type Workout struct {
	Duration float64 `json:"duration"`
	Type     string  `json:"type"`
	Area     string  `json:"area"`
	Level    string  `json:"level"`
}

// All Settings of settings > audio
type Settings struct {
	Audio      Audio      `json:"audio"`
	UrlTrigger UrlTrigger `json:"urltrigger"`
}

type Audio struct {
	Activate          bool `json:"activate"`
	ActivateCountdown bool `json:"activatecountdown"`
	ActivateExercise  bool `json:"activateexercise"`
	ActivatePause     bool `json:"activatepause"`
}

// All Settings of settings > URL trigger
type UrlTrigger struct {
	Url string `json:"url"`
}

// Declare Config Specific Variables
var Conf Config     // Config Struct
var ConfPath string // Config file Path

// Initially Fill config Struct
func Conf_initConf() error {

	// Innitialize defaults
	conf_setDefaults()

	// Initialize Config File if not present
	conffolder := configdir.LocalConfig("gopherletics")
	err := configdir.MakePath(conffolder) // Ensure it exists.
	if err != nil {
		zap.L().Error("Could not create config folder", zap.Error(err))
		return err
	}

	// Set Config File Path
	viper.SetConfigName(conf_filename)
	viper.SetConfigType(conf_extension)
	viper.AddConfigPath(conffolder)

	// Read Config File
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			err := viper.WriteConfigAs(conffolder + "/" + conf_filename + "." + conf_extension)
			if err != nil {
				zap.L().Error("Could not write config file", zap.Error(err))
			}
		} else {
			zap.L().Error("Error loading config file:", zap.Error(err))
		}
	}

	err = viper.Unmarshal(&Conf)
	if err != nil {
		zap.L().Error("Could not unmarshal config file", zap.Error(err))
	}

	// Debug
	fmt.Println("Initialized Configuration:")
	fmt.Print(Conf)

	return nil

}

// Write current Config
func Conf_Write() {
	err := viper.WriteConfig()
	if err != nil {
		zap.L().Error("Could not write config file", zap.Error(err))
	}
}
