package gui

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/SHU-red/GopherLetics.git/internal/glob"
	"github.com/SHU-red/GopherLetics.git/internal/workout"
)

func updatevals() {

}

// Update the Shown timer String
func update_timer_str() {
	time, _ := glob.Gui.Timer.Get()
	timer.Text = fmt.Sprintf("%04d", time)

	// Update the shown timer
	timer.Refresh()

}

// Update Shown workouts
func update_workout_list() {

	list = widget.NewList(
		func() int {
			return len(workout.Wo)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("template", func() {})
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {

			// Generate Text based on Type
			switch ty := workout.Wo[i].Ty; ty {

			case "heading":
				o.(*widget.Button).SetText(strings.ToUpper(workout.Wo[i].Na))
				o.(*widget.Button).Importance = widget.LowImportance
				o.(*widget.Button).SetIcon(theme.InfoIcon())

			case "exercise":

				o.(*widget.Button).SetText(strconv.Itoa(workout.Wo[i].Du) + "s: " + workout.Wo[i].Na)
				o.(*widget.Button).Importance = widget.WarningImportance
				o.(*widget.Button).SetIcon(theme.ColorChromaticIcon())

			// Pause
			default:

				o.(*widget.Button).SetText(strconv.Itoa(workout.Wo[i].Du) + "s: " + workout.Wo[i].Na)
				o.(*widget.Button).Importance = widget.SuccessImportance
				o.(*widget.Button).SetIcon(theme.HistoryIcon())

			}

			o.(*widget.Button).Refresh()

		})

}

// Refresh workout and reset all necessary values
func update_all() {

	// Update Timer
	update_timer_str()

	// Update Workout
	update_workout_list()

}
