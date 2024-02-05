package gui

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/SHU-red/GopherLetics.git/internal/glob"
	"github.com/SHU-red/GopherLetics.git/internal/tts"
)

// Create Channel
var stop = make(chan bool)

// Permanently running calculations
func runloop() {

	// Update all GUI elements
	updatevals()

}

// Execute Actions for Toggling Play button
func toggleplay(button *widget.Button) {

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

	prog, _ := glob.Gui.Progress.Get()
	glob.Gui.Progress.Set(prog + 0.1)

	// Count Timer if active
	if glob.Gui.Play {

		// Debug
		fmt.Println("before false")

		// Debug
		fmt.Println("before channel")

		// Debug
		fmt.Println(glob.Gui.Timer_Str)

		// Start Timer execution
		go count_timer()

	} else {

		// Debug
		fmt.Println("else")

		// Stop concurrent Functions
		stop <- true

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
			glob.Gui.Timer.Set(ti - 1)
			update_timer_str()

			// Debug
			fmt.Println(glob.Gui.Timer_Str)

			// Acousitc Coutnter
			//TODO
		}
	}

}
