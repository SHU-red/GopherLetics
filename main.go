// - [ ] Move Buttons to 1. Settings, 2. Workout, 3. Refresh
// - [ ] "Transition to Pyramid"
// - [ ] Show workout parameters left of counter
// - [ ] Fix proceeding after last workout
// - [ ] Keep last opened workout and do not open the same again
// - [ ] Keep last opened workout and open the new one on click or play/pause

package main

import (
	"os"

	"github.com/SHU-red/GopherLetics.git/internal/glob"
	"github.com/SHU-red/GopherLetics.git/internal/gui"
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
	// Overwrite Configuration from config-file
	glob.Conf_Load()

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
	for i, _ := range args {

		if args[i] == "--debug" {
			glob.Args.Debug = true
		}

	}
}
