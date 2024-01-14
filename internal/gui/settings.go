package gui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Show settings overlay
func settings() {

	// Create Settings menu
	settings_c := container.NewBorder()

	// Show PopUp
	widget.ShowModalPopUp(settings_c, w.Canvas())

}
