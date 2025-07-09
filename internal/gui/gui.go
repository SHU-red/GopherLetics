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
	"github.com/SHU-red/GopherLetics.git/internal/glob"
	"github.com/SHU-red/GopherLetics.git/internal/tts"
	"github.com/SHU-red/GopherLetics.git/internal/workout"
	"github.com/kr/pretty"
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

// Last shown Workout
var last_shown_workout string

// Play/Pause Button
var playbutton widget.Button

func Main() {

	// Initialize glob variables
	glob.Gui_initval()

	// Fyne App
	a := app.New()
	w = a.NewWindow("GopherLetics")

	// Show Image Window
	go ShowImageWindow(a)

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
	top := container.NewBorder(nil, nil, nil, nil, top_center)

	// Bottom
	bottom_left := container.New(layout.NewHBoxLayout(),
		widget.NewHyperlink("GopherLetics v0.1", url_gopherletics),
		widget.NewHyperlink("Fyne v0", url_fyne),
		widget.NewHyperlink("Golang v0", url_golang))
	bottom_right := widget.NewLabel("")
	bottom := container.NewBorder(nil, nil, bottom_left, bottom_right, nil)

	// Play Button
	PlayButtonPause(&playbutton)
	playbutton.OnTapped = func() { toggleplay(&playbutton) }

	// ToolBar
	toolbar := container.NewHBox(widget.NewButtonWithIcon("", theme.MediaSkipPreviousIcon(), previousExercise), &playbutton, widget.NewButtonWithIcon("", theme.MediaSkipNextIcon(), nextExercise))

	low_left := container.NewHBox(widget.NewButtonWithIcon("Workout", theme.AccountIcon(), workoutSettings), widget.NewButtonWithIcon("Refresh", theme.ViewRefreshIcon(), refresh))
	// Menu
	menu := container.NewBorder(nil, nil, low_left, widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), settings), container.NewCenter(toolbar))

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
	// lv4 = container.NewVSplit(&timercontainer, list)
	// lv3 := container.NewHSplit(lv4, exercise)
	// lv2 := container.NewBorder(nil, progbar, nil, nil, lv3)
	lv2 := container.NewBorder(&timercontainer, progbar, nil, nil, list)
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

func nextExercise() {
	currentIndex := glob.Gui.WorkoutNr
	for i := currentIndex + 1; i < len(workout.Wo); i++ {
		if workout.Wo[i].Ty == "exercise" {
			SwitchWorkout(i)
			return
		}
	}
	// If no next exercise found, loop back to the first exercise
	for i := 0; i < len(workout.Wo); i++ {
		if workout.Wo[i].Ty == "exercise" {
			SwitchWorkout(i)
			return
		}
	}
}

func previousExercise() {
	currentIndex := glob.Gui.WorkoutNr
	for i := currentIndex - 1; i >= 0; i-- {
		if workout.Wo[i].Ty == "exercise" {
			SwitchWorkout(i)
			return
		}
	}
	// If no previous exercise found, loop back to the last exercise
	for i := len(workout.Wo) - 1; i >= 0; i-- {
		if workout.Wo[i].Ty == "exercise" {
			SwitchWorkout(i)
			return
		}
	}
}

func refresh() {

	// Pull new Workout
	workout.Wo.Fetch()

	// If Workout could be feched
	if len(workout.Wo) > 0 {

		// Switch workout
		SwitchWorkout(0)

		// Update content
		update_all()

	}

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

	// if Workout lenght is 0
	if len(workout.Wo) == 0 {
		return
	}

	zap.L().Debug("Switched Workout", zap.Int("Switching to", x))
	zap.L().Debug("Switching to", zap.Int("Workout Lenght", len(workout.Wo)))

	// If Workout has finished
	if x >= len(workout.Wo) {

		// Set Indicator to false
		glob.Gui.Play = false

		// Reset PlayButton
		PlayButtonPause(&playbutton)

		// Speech feedback
		go tts.SpeakRand("done")

		// Return to start
		x = 0

		// Debug
		zap.L().Debug("Finished Workout", zap.Int("Switching to", glob.Gui.WorkoutNr))
		pretty.Print("-thats it")

		// Stop Timer countown non blocking
		go func() {
			stop <- true
		}()

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

	// Update Image
	if workout.Wo[i].Ty == "rest" || workout.Wo[i].Ty == "transition" {
		if nextExercise := findNextExercise(i); nextExercise != nil {
			go UpdateImage(nextExercise.Na)
		}
	} else {
		go UpdateImage(workout.Wo[i].Na)
	}

	// Update All
	update_all()

}

func findNextExercise(currentIndex int) *workout.Workout {
	for i := currentIndex + 1; i < len(workout.Wo); i++ {
		if workout.Wo[i].Ty == "exercise" {
			return &workout.Wo[i]
		}
	}
	return nil
}

// Show workout in Browser
func ShowWorkout(x int, force bool) {
}
