package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/SHU-red/GopherLetics.git/internal/glob"
)

// Popover for Workout settings
func workoutSettings() {

	// Generate Widgets
	dur_bind := binding.BindFloat(&glob.Conf.Workout.Duration)
	dur := binding.FloatToStringWithFormat(dur_bind, "Duration: %0.0f min")

	label_duration := widget.NewLabelWithData(dur)

	slider_duration := widget.NewSlider(10, 60)
	slider_duration.Value = float64(glob.Conf.Workout.Duration)
	slider_duration.Bind(dur_bind)

	label_type := widget.NewLabel("Type")
	dropdown_type := widget.NewSelect(glob.Choices_Type, setType)
	dropdown_type.Selected = glob.Conf.Workout.Type

	label_level := widget.NewLabel("Level")
	dropdown_level := widget.NewSelect(glob.Choices_Level, setLevel)
	dropdown_level.Selected = glob.Conf.Workout.Level

	label_area := widget.NewLabel("Area")
	dropdown_area := widget.NewSelect(glob.Choices_Area, setArea)
	dropdown_area.Selected = glob.Conf.Workout.Area

	// Generate a Widget List
	box := container.New(layout.NewFormLayout(), label_duration, slider_duration, label_type, dropdown_type, label_level, dropdown_level, label_area, dropdown_area)

	// Show PopUp
	var modal fyne.Widget
	modal = widget.NewModalPopUp(
		container.NewVBox(
			box,
			widget.NewButton("Close", func() { modal.Hide() }),
		), w.Canvas())
	modal.Show()

}

// Set Type from Dropdown
func setType(choice string) {
	glob.Conf.Workout.Type = choice
}

// Set Type from Dropdown
func setLevel(choice string) {
	glob.Conf.Workout.Level = choice
}

// Set Area from Dropdown
func setArea(choice string) {
	glob.Conf.Workout.Area = choice
}

func settingsSettings() {

	// Generate Widgets
	dur_bind := binding.BindFloat(&glob.Conf.Workout.Duration)
	dur := binding.FloatToStringWithFormat(dur_bind, "Duration: %0.0f min")

	label_duration := widget.NewLabelWithData(dur)

	slider_duration := widget.NewSlider(10, 60)
	slider_duration.Value = float64(glob.Conf.Workout.Duration)
	slider_duration.Bind(dur_bind)

	label_type := widget.NewLabel("Type")
	dropdown_type := widget.NewSelect(glob.Choices_Type, setType)
	dropdown_type.Selected = glob.Conf.Workout.Type

	label_level := widget.NewLabel("Level")
	dropdown_level := widget.NewSelect(glob.Choices_Level, setLevel)
	dropdown_level.Selected = glob.Conf.Workout.Level

	label_area := widget.NewLabel("Area")
	dropdown_area := widget.NewSelect(glob.Choices_Area, setArea)
	dropdown_area.Selected = glob.Conf.Workout.Area

	// Generate a Widget List
	box := container.New(layout.NewFormLayout(), label_duration, slider_duration, label_type, dropdown_type, label_level, dropdown_level, label_area, dropdown_area)

	// Show PopUp
	var modal fyne.Widget
	modal = widget.NewModalPopUp(
		container.NewVBox(
			box,
			widget.NewButton("Close", func() { modal.Hide() }),
		), w.Canvas())
	modal.Show()

}
