package glob

import (
	"errors"
	"fmt"
	"os"

	"dario.cat/mergo"
	"github.com/BurntSushi/toml"
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
		} else {
			zap.L().Error("Error loading config file:", zap.Error(err))
		}
	}

	// Debug
	fmt.Println("Initialized Configuration:")
	fmt.Print(Conf)

	return nil

}

// Load current Config
func Conf_Load() {

	// Temporary Configuration for later Merge
	var tempconf Config

	// Load Config from File
	if _, err := toml.DecodeFile(ConfPath, &tempconf); err != nil {
		panic(err)
	}

	// (Overwriting-)Merge of temporary loaded Conf into Initialized Conf
	if err := mergo.Merge(&Conf, tempconf, mergo.WithOverride); err != nil {
		panic(err)
	}

	// Loaded Config File
	fmt.Println("Initialized Configuration:")
	fmt.Print(Conf)

}

// Write current Config
func Conf_Write() {

}

// Check if a file exists
func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}
