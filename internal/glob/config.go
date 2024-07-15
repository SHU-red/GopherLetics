package glob

import (
	"errors"
	"fmt"
	"log"
	"os"

	"dario.cat/mergo"
	"github.com/BurntSushi/toml"
	"github.com/kirsle/configdir"
)

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
	Audio Audio      `json:"audio"`
	Url   UrlTrigger `json:"urltrigger"`
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
func Conf_initConf() {

	// Initialize Config File if not present
	conffolder := configdir.LocalConfig("gopherletics")
	err := configdir.MakePath(conffolder) // Ensure it exists.
	if err != nil {
		panic(err)
	}

	// Set Config File Path
	ConfPath = conffolder + "/conf.toml"

	// Workout
	Conf.Workout.Duration = 30
	Conf.Workout.Type = "strength"
	Conf.Workout.Area = "full"
	Conf.Workout.Level = "beginner"

	// Audio
	Conf.Settings.Audio.Activate = true
	Conf.Settings.Audio.ActivateCountdown = true
	Conf.Settings.Audio.ActivateExercise = true
	Conf.Settings.Audio.ActivatePause = true

	// Create Config file if not already existing
	if !checkFileExists(ConfPath) {
		f, err := os.Create(ConfPath)
		if err != nil {
			// failed to create/open the file
			log.Fatal(err)
		}
		if err := toml.NewEncoder(f).Encode(Conf); err != nil {
			// failed to encode
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			// failed to close the file
			log.Fatal(err)

		}
	}

	// Debug
	fmt.Println("Initialized Configuration:")
	fmt.Print(Conf)

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
