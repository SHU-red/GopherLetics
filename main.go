// - [ ] Keep last opened workout and do not open the same again
// - [ ] Keep last opened workout and open the new one on click or play/pause

package main

import (
	"os"

	"github.com/SHU-red/GopherLetics/internal/glob"
	"github.com/SHU-red/GopherLetics/internal/gui"
	"github.com/SHU-red/GopherLetics/internal/tts"
	"github.com/SHU-red/GopherLetics/internal/workout"
	"go.uber.org/zap"
)

func main() {

	// Init Zap Logger
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	// Extract Arguments
	getArgs()

	// Initialize Configuration with default values
	glob.Conf_initConf()
	// Init speech
	tts.SpeakInit()
	// Fetch all exercises
	workout.FetchAllExercises()

	// Run GUI
	gui.Main()

}

// Process Args
func getArgs() {

	// Init Args
	glob.Args.Debug = false

	// Read Args
	args := os.Args[1:]

	// Loop through Args
	for i := range args {

		if args[i] == "--debug" {
			glob.Args.Debug = true
		}

	}
}
