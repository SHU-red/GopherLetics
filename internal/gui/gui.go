package gui

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/SHU-red/GopherLetics.git/internal/global"
)

// Whole Content
var content fyne.Container

// Declare Variable contents
var playbutton widget.ToolbarAction
var toolbar widget.Toolbar
var timer canvas.Text

func Main() {

	// Initialize global variables
	global.Gui_initval()

	// Fyne App
	a := app.New()
	w := a.NewWindow("GopherLetics")

	// URLs
	url_gopherletics, err := url.Parse("https://github.com/SHU-red/GopherLetics")
	if err != nil {
		fmt.Println(err)
		return
	}
	url_fyne, err := url.Parse("https://fyne.io/")
	if err != nil {
		fmt.Println(err)
		return
	}
	url_golang, err := url.Parse("https://go.dev/")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Top
	top_center := container.NewCenter(widget.NewLabel("GopherLetics"))
	top_right := widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), func() {})
	top := container.NewBorder(nil, nil, nil, top_right, top_center)

	// Bottom
	bottom_left := container.New(layout.NewHBoxLayout(),
		widget.NewHyperlink("GopherLetics v0.1", url_gopherletics),
		widget.NewHyperlink("Fyne v0", url_fyne),
		widget.NewHyperlink("Golang v0", url_golang))
	bottom_right := widget.NewLabel("")
	bottom := container.NewBorder(nil, nil, bottom_left, bottom_right, nil)

	// Toolbar
	playbutton = *widget.NewToolbarAction(theme.MediaPlayIcon(), toggleplay)
	toolbar = *widget.NewToolbar(widget.NewToolbarAction(theme.MediaSkipPreviousIcon(), func() {}), &playbutton, widget.NewToolbarAction(theme.MediaSkipNextIcon(), func() {}))

	// Menu
	menu := container.NewBorder(nil, nil, widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), refresh), widget.NewButtonWithIcon("", theme.AccountIcon(), func() {}), container.NewCenter(&toolbar))

	// Progress bar
	progbar := widget.NewProgressBarWithData(global.Gui.Progress)

	// Exercise
	exercise := widget.NewLabel("Exercise")

	// Timer
	timer.Text = "0004"
	timer.TextSize = 100

	// List
	list := widget.NewLabel("List")

	// Excercise Levels
	lv4 := container.NewVSplit(container.NewCenter(&timer), list)
	lv3 := container.NewHSplit(lv4, exercise)
	lv2 := container.NewBorder(nil, progbar, nil, nil, lv3)
	lv1 := container.NewBorder(nil, menu, nil, nil, lv2)

	// Main Content
	content = *container.NewBorder(top, bottom, nil, nil, lv1)

	// Set Content
	w.SetContent(&content)

	// Runloop
	go runloop()

	// Show App
	w.ShowAndRun()
}

func refresh() {

	// Refresh values
	global.Gui.Timer.Set(10)

	// Update content
	update_all()

}
