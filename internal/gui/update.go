package gui

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/SHU-red/GopherLetics/internal/glob"
	"github.com/SHU-red/GopherLetics/internal/workout"
)

// Update the Shown timer String
func update_timer_str() {
	time, _ := glob.Gui.Timer.Get()
	timer.Text = fmt.Sprintf("%04d", time)

	fyne.Do(func() {
		timer.Refresh()
	})
}

// Create the workout list widget once (called from Main)
func create_workout_list() {
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
				if i == glob.Gui.WorkoutNr {
					o.(*widget.Button).Importance = widget.HighImportance
				} else {
					o.(*widget.Button).Importance = widget.WarningImportance
				}
				o.(*widget.Button).SetIcon(theme.ColorChromaticIcon())
				idx := i
				o.(*widget.Button).OnTapped = func() { SwitchWorkout(idx) }

			// Rest / Transition
			default:

				o.(*widget.Button).SetText(strconv.Itoa(workout.Wo[i].Du) + "s: " + workout.Wo[i].Na)
				if i == glob.Gui.WorkoutNr {
					o.(*widget.Button).Importance = widget.HighImportance
				} else {
					o.(*widget.Button).Importance = widget.SuccessImportance
				}
				o.(*widget.Button).SetIcon(theme.HistoryIcon())
				idx := i
				o.(*widget.Button).OnTapped = func() { SwitchWorkout(idx) }

			}

			o.(*widget.Button).Refresh()
			o.Refresh()

		})
}

// Refresh the Shown workouts list
func refresh_workout_list() {
	list.Refresh()
	fyne.Do(func() {
		content.Refresh()
		w.Content().Refresh()
		w.Canvas().Refresh(list)
	})
}

// Refresh workout and reset all necessary values
func update_all() {
	// Update Timer
	update_timer_str()

	// Refresh Workout list
	refresh_workout_list()
}
