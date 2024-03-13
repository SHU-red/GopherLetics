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
			// go tts.SpeakRand("play")
			go tts.Speak(workout.Wo[glob.Gui.WorkoutNr].Na)

			zap.L().Debug("Starting Counter")

			// Concurrently run Counter
			go count_timer()

		} else {

			// Format Play Button to Pause
			PlayButtonPause(button)

			// Speech feedback
			go tts.SpeakRand("stop")

			zap.L().Debug("Stopping Coutner")

			// Stop concurrent Functions
			stop <- true

		}

	}

}

// Mathematically count down timer
func count_timer() {

	// Execute every second
	for range time.Tick(time.Second) {

		// !! Non-blocking !! channel read
		select {
		case <-stop:

			zap.L().Debug("Concurrent Countdown stopped")

			// Cancel countdown
			return

		// If no message with true is received
		default:

			// Get current Timer
			ti, _ := glob.Gui.Timer.Get()
			glob.Gui.Timer.Set(ti - 1)
			zap.L().Debug("Counting", zap.Int("Timer", ti))

			// If done
			if ti < 1 {
				zap.L().Debug("Next Workout")
				SwitchWorkout(glob.Gui.WorkoutNr + 1)
			} else { // Proceed
				zap.L().Debug("Update Timer")
				go update_timer_str()
			}

			// Acousitc Countdown for 5, 4, 3, 2 and 1
			// Has to be sent in advantage for syncronization with shown timer
			if ti <= 6 && ti > 1 {
				go tts.Speak(strconv.Itoa(ti - 1))
			}

			// Show next workout in browser
			if ti == 10 && glob.Gui.WorkoutNr < len(workout.Wo)-1 {
				go ShowWorkout(glob.Gui.WorkoutNr, false)
			}

		}
	}

}
