package gui

import (
	"fmt"
	"net/url"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/SHU-red/GopherLetics.git/internal/glob"
	"github.com/SHU-red/GopherLetics.git/internal/tts"
	"github.com/SHU-red/GopherLetics.git/internal/workout"
	"go.uber.org/zap"
)

// Window
var w fyne.Window

// Whole Content of Base Window
var content fyne.Container

// Declare Variable contents
var timer canvas.Text

// Declare global bar
var progbar fyne.Widget

// Timercontainer
var timercontainer fyne.Container

// For Rorkout list
var list fyne.Widget
var lv4 *container.Split

func Main() {

	// Initialize glob variables
	glob.Gui_initval()

	// Fyne App
	a := app.New()
	w = a.NewWindow("GopherLetics")

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
	top_right := widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), settings)
	top := container.NewBorder(nil, nil, nil, top_right, top_center)

	// Bottom
	bottom_left := container.New(layout.NewHBoxLayout(),
		widget.NewHyperlink("GopherLetics v0.1", url_gopherletics),
		widget.NewHyperlink("Fyne v0", url_fyne),
		widget.NewHyperlink("Golang v0", url_golang))
	bottom_right := widget.NewLabel("")
	bottom := container.NewBorder(nil, nil, bottom_left, bottom_right, nil)

	// Play Button
	var playbutton widget.Button
	PlayButtonPause(&playbutton)
	playbutton.OnTapped = func() { toggleplay(&playbutton) }

	// ToolBar
	toolbar := container.NewHBox(widget.NewButtonWithIcon("", theme.MediaSkipPreviousIcon(), func() {}), &playbutton, widget.NewButtonWithIcon("", theme.MediaSkipNextIcon(), func() {}))

	// Menu
	menu := container.NewBorder(nil, nil, widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), refresh), widget.NewButtonWithIcon("", theme.AccountIcon(), func() {}), container.NewCenter(toolbar))

	// Progress bar
	progbar = widget.NewProgressBarWithData(glob.Gui.Progress)

	// Exercise
	// exercise := widget.NewLabel("Exercise")

	// Timer
	timer.Text = "NO DATA"
	timer.TextSize = 100
	timercontainer = *container.NewCenter(&timer)

	// List
	update_workout_list()

	// Excercise Levels
	lv4 = container.NewVSplit(&timercontainer, list)
	// lv3 := container.NewHSplit(lv4, exercise)
	// lv2 := container.NewBorder(nil, progbar, nil, nil, lv3)
	lv2 := container.NewBorder(nil, progbar, nil, nil, lv4)
	lv1 := container.NewBorder(nil, menu, nil, nil, lv2)

	// Main Content
	content = *container.NewBorder(top, bottom, nil, nil, lv1)

	// Set Content
	w.SetContent(&content)

	// Catch Keyboard strokes
	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		fmt.Println("Key pressed: " + k.Name)
		switch k.Name {

		// Space = Play/Pause
		case fyne.KeySpace:

			go toggleplay(&playbutton)
		}

		// Left = Previous
		//TODO

		// Right = Next
		// TODO

		// W = Workout
		//TODO

		// S = Settings
		//TODO

		// R = Refresh
		//TODO
	})

	// // Runloop
	// go runloop()

	// Show App
	w.ShowAndRun()
}

func refresh() {

	// Pull new Workout
	workout.Wo.Fetch()

	// Switch workout
	SwitchWorkout(0)

	// Load first exercise in browser
	go ShowWorkout(glob.Gui.WorkoutNr, true)

	// Update content
	update_all()

}

// Set all Play/Pause Buttons Attributes for Play State (Showing Pause Icon)
func PlayButtonPlay(button *widget.Button) {

	button.SetText("")
	button.SetIcon(theme.MediaPauseIcon())
	button.Importance = widget.DangerImportance
	button.Refresh()

}

// Set all Play/Pause Buttons Attributes for Pause State (Showing Play Icon)
func PlayButtonPause(button *widget.Button) {

	button.SetText("")
	button.SetIcon(theme.MediaPlayIcon())
	button.Importance = widget.MediumImportance
	button.Refresh()

}

// Switch to certain workout and restet/refresh necessary components
func SwitchWorkout(x int) {

	zap.L().Debug("Switched Workout", zap.Int("Switching to", x))

	// If Workout has finished
	if x >= len(workout.Wo) {
		glob.Gui.Play = false
		go tts.SpeakRand("done")
		SwitchWorkout(0)
		return
	}

	//Set workout Pointer to first non-heading starting from x
	i := x
	for workout.Wo[i].Ty == "heading" {
		i++
	}
	glob.Gui.WorkoutNr = i
	// Reset Timer to current Workout Duration
	glob.Gui.Timer.Set(workout.Wo[i].Du)

	// Update Progress
	p := float64(int(glob.Gui.WorkoutNr)) / float64(len(workout.Wo))
	zap.L().Debug("New Progress calculated", zap.Float64("Progress", p))
	glob.Gui.Progress.Set(p)

	// If Workout switch was done during Play
	if glob.Gui.Play {
		go tts.Speak(workout.Wo[i].Na)
	}

	// Update All
	update_all()

}

// Show workout in Browser
func ShowWorkout(x int, force bool) {

	// If force is true
	if force {
		_ = exec.Command("xdg-open", "https://www.youtube.com/results?search_query="+workout.Wo[x].Na+" Exercise").Start()
		return
	}

	// Start analyzing beginning from the next exercise
	x++

	// Ignore headings
	for workout.Wo[x].Ty == "heading" && x < len(workout.Wo) {
		x++
	}

	// If end is not already reached
	if x < len(workout.Wo) && workout.Wo[x].Ty == "exercise" {
		// Concurrently open browser for showing the workout
		_ = exec.Command("xdg-open", "https://www.youtube.com/results?search_query="+workout.Wo[x].Na+" Exercise").Start()
	}

}
