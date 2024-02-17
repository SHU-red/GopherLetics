package glob

import (
	"fmt"

	"fyne.io/fyne/v2/data/binding"
)

// Declare Runtime Struct
type Gui_Struct struct {
	Timer     binding.Int
	Timer_Str string
	Play      bool
	WorkoutNr int
	Progress  binding.Float
	Duration  int
}

// Default Values
var Gui Gui_Struct

// Initialize all Gui Values
func Gui_initval() {

	Gui.Timer = binding.NewInt()
	Gui.Timer.Set(0)

	Gui.Timer_Str = "No Data"

	Gui.Play = false

	Gui.Progress = binding.NewFloat()
	Gui.Progress.Set(0)

	fmt.Println("Init done")

}

// To Store Args
var Args struct {
	Debug bool
}
