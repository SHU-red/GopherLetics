package gui

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2/theme"
	"github.com/SHU-red/GopherLetics.git/internal/global"
)

// Create Channel
var stop = make(chan bool)

// Permanently running calculations
func runloop() {

	// Update all GUI elements
	updatevals()

}

// Execute Actions for Toggling Play button
func toggleplay() {

	// Toggle play
	global.Gui.Play = !global.Gui.Play

	// Debug
	fmt.Println("Toggled Play to " + strconv.FormatBool(global.Gui.Play))
	// Change Button icon
	if global.Gui.Play {
		playbutton.SetIcon(theme.MediaPlayIcon())
	} else {
		playbutton.SetIcon(theme.MediaPauseIcon())
	}

	prog, _ := global.Gui.Progress.Get()
	global.Gui.Progress.Set(prog + 0.1)

	// Debug
	fmt.Println("New playbutton icon: " + playbutton.Icon.Name())

	// Set Content
	//w.SetContent(content)
	//playbutton.ToolbarObject().Refresh()
	//toolbar.Refresh()

	// Count Timer if active
	if global.Gui.Play {

		// Debug
		fmt.Println("before false")

		// Debug
		fmt.Println("before channel")

		// Debug
		fmt.Println(global.Gui.Timer_Str)

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
			ti, _ = global.Gui.Timer.Get()
			global.Gui.Timer.Set(ti - 1)
			update_timer_str()

			// Debug
			fmt.Println(global.Gui.Timer_Str)

			// Acousitc Coutnter
			//TODO
		}
	}

}
