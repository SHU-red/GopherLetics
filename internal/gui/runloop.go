package gui

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/SHU-red/GopherLetics.git/internal/glob"
	"github.com/SHU-red/GopherLetics.git/internal/tts"
	"github.com/SHU-red/GopherLetics.git/internal/workout"
	"go.uber.org/zap"
)

// Create Channel
var stop = make(chan bool)

// Permanently running calculations
func runloop() {

}

// Execute Actions for Toggling Play button
func toggleplay(button *widget.Button) {

	// Only if Workouts are present
	if len(workout.Wo) > 0 {

		// Toggle play
		glob.Gui.Play = !glob.Gui.Play

		// Debug
		fmt.Println("Toggled Play to " + strconv.FormatBool(glob.Gui.Play))
		// Change Button icon
		if glob.Gui.Play {

			// Format Play Button to Play
			PlayButtonPlay(button)

			// Speech feedback
			go tts.SpeakRand("play")

		} else {

			// Format Play Button to Pause
			PlayButtonPause(button)

			// Speech feedback
			go tts.SpeakRand("stop")
		}

		// Count Timer if active
		if glob.Gui.Play {

			zap.L().Debug("Starting Counter")

			// Concurrently run Counter
			go count_timer()

		} else {

			zap.L().Debug("Stopping Coutner")

			// Stop concurrent Functions
			stop <- true

		}

	}

}

// Mathematically count down timer
func count_timer() {

	// Declare
	var ti int

	// Debug
	fmt.Println("Counting")

	// Execute every second
	for range time.Tick(time.Second) {

		// !! Non-blocking !! channel read
		select {
		case <-stop:

			// Cancel countdown
			return

		// If no message with true is received
		default:

			// Decrease Counter
			ti, _ = glob.Gui.Timer.Get()
			ti = ti - 1

			// If done
			if ti < 1 {
				SwitchWorkout(glob.Gui.WorkoutNr + 1)
			} else { // Proceed
				glob.Gui.Timer.Set(ti - 1)
				go update_timer_str()
			}

			// Acousitc Coutnter
			if ti <= 5 && ti > 0 {
				go tts.Speak(strconv.Itoa(ti))
			}
			//TODO

		}
	}

}
