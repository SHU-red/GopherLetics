package gui

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/SHU-red/GopherLetics/internal/glob"
	"github.com/SHU-red/GopherLetics/internal/tts"
	"github.com/SHU-red/GopherLetics/internal/workout"
	"go.uber.org/zap"
)

// Signal-only channel; recreated on each play to prevent double-start
var stop chan struct{}

// Execute Actions for Toggling Play button
func toggleplay(button *widget.Button) {

	// Only if Workouts are present
	if len(workout.Wo) == 0 {
		return
	}

	// Toggle play
	glob.Gui.Play = !glob.Gui.Play

	zap.L().Debug("toggled play", zap.Bool("play", glob.Gui.Play))

	if glob.Gui.Play {

		// Format Play Button to Play (Pause icon)
		PlayButtonPlay(button)

		// Speech feedback
		go tts.Speak(workout.Wo[glob.Gui.WorkoutNr].Na)

		// Create fresh stop channel
		stop = make(chan struct{})

		// Concurrently run Counter
		go count_timer()

	} else {

		// Format Play Button to Pause (Play icon)
		PlayButtonPause(button)

		// Speech feedback
		go tts.SpeakRand("stop")

		// Stop concurrent Functions
		close(stop)

	}

}

// Mathematically count down timer
func count_timer() {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-stop:
			zap.L().Debug("Concurrent Countdown stopped")
			return
		case <-ticker.C:

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
				update_timer_str()
			}

			// Acoustic countdown for 5, 4, 3, 2 and 1
			if ti <= 6 && ti > 1 {
				go tts.Speak(strconv.Itoa(ti - 1))
			}

		}
	}

}
