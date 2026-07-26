package gui

import (
	"fmt"
	"net/url"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/SHU-red/GopherLetics/internal/glob"
	"github.com/SHU-red/GopherLetics/internal/tts"
	"github.com/SHU-red/GopherLetics/internal/workout"
	"github.com/spf13/viper"
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

// Workout list widget
var list fyne.Widget

// Play/Pause Button
var playbutton widget.Button

func Main() {

	// Initialize glob variables
	glob.Gui_initval()

	// Fyne App
	a := app.NewWithID("com.github.SHU-red.GopherLetics")
	w = a.NewWindow("GopherLetics")

	// Create shared exercise widgets
	createExerciseWidgets()

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

	// Top title
	top := container.NewCenter(widget.NewLabelWithStyle("GopherLetics", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))

	// Bottom footer
	footer := container.NewHBox(
		widget.NewHyperlink("GopherLetics v0.1", url_gopherletics),
		widget.NewHyperlink("Fyne v0", url_fyne),
		widget.NewHyperlink("Golang v0", url_golang),
	)

	// Play Button
	PlayButtonPause(&playbutton)
	playbutton.OnTapped = func() { toggleplay(&playbutton) }

	// Toolbar
	toolbar := container.NewHBox(
		widget.NewButtonWithIcon("", theme.MediaSkipPreviousIcon(), previousExercise),
		&playbutton,
		widget.NewButtonWithIcon("", theme.MediaSkipNextIcon(), nextExercise),
	)

	menuLeft := container.NewHBox(
		widget.NewButtonWithIcon("Workout", theme.AccountIcon(), workoutSettings),
		widget.NewButtonWithIcon("Refresh", theme.ViewRefreshIcon(), refresh),
	)
	menu := container.NewBorder(nil, nil, menuLeft, widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), settings), container.NewCenter(toolbar))

	// Progress bar
	progbar = widget.NewProgressBarWithData(glob.Gui.Progress)

	// Timer
	timer.Text = "0030"
	timer.TextSize = 72
	timercontainer = *container.NewCenter(&timer)

	// Image panel: timer on top, image fills center, name label below
	imagePanel := container.NewBorder(&timercontainer, exerciseNameLabel, nil, nil, exerciseImage)

	// Create workout list once
	create_workout_list()

	// Main split: image panel (65%) | workout list (35%)
	split := container.NewHSplit(imagePanel, list)
	split.SetOffset(0.65)

	// Stack: split above, progress bar below
	mainArea := container.NewBorder(nil, progbar, nil, nil, split)

	// Outer framing: top title, menu at bottom of main area, footer
	content = *container.NewBorder(top, footer, nil, nil,
		container.NewBorder(nil, menu, nil, nil, mainArea),
	)

	w.SetContent(&content)

	// Refresh on startup
	refresh()

	// Keyboard shortcuts
	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeySpace:
			toggleplay(&playbutton)
		case fyne.KeyLeft:
			previousExercise()
		case fyne.KeyRight:
			nextExercise()
		case fyne.KeyR:
			refresh()
		case fyne.KeyS:
			settings()
		case fyne.KeyW:
			workoutSettings()
		}
	})

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

	// Generate new workout from config
	cfg := workout.GenerateConfig{
		Duration:  int(glob.Conf.Workout.Duration),
		Type:      glob.Conf.Workout.Type,
		Area:      glob.Conf.Workout.Area,
		Level:     glob.Conf.Workout.Level,
		Equipment: glob.Conf.Workout.Equipment,
		Template:  glob.Conf.Workout.Template,
	}
	workout.Wo.Generate(cfg)

	if len(workout.Wo) > 0 {
		SwitchWorkout(0)
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
// Switch to certain workout and reset/refresh necessary components
func SwitchWorkout(x int) {

	// if Workout length is 0
	if len(workout.Wo) == 0 {
		return
	}

	zap.L().Debug("switched workout", zap.Int("to", x), zap.Int("total", len(workout.Wo)))

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

		zap.L().Debug("finished workout")

		// Stop Timer
		if stop != nil {
			close(stop)
		}

	}

	// Set workout Pointer to first non-heading starting from x
	i := x
	for workout.Wo[i].Ty == "heading" {
		i++
	}
	glob.Gui.WorkoutNr = i
	// Reset Timer to current Workout Duration
	glob.Gui.Timer.Set(workout.Wo[i].Du)

	// Update Progress
	p := float64(int(glob.Gui.WorkoutNr)) / float64(len(workout.Wo))
	zap.L().Debug("new progress", zap.Float64("progress", p))
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


func settings() {

	// Create Checkboxes
	activateAudio := widget.NewCheck("Activate Audio", func(b bool) {
		viper.Set("settings.audio.activate", b)
		glob.Conf_Write()
	})
	activateAudio.SetChecked(viper.GetBool("settings.audio.activate"))

	activateCountdown := widget.NewCheck("Activate Countdown", func(b bool) {
		viper.Set("settings.audio.activatecountdown", b)
		glob.Conf_Write()
	})
	activateCountdown.SetChecked(viper.GetBool("settings.audio.activatecountdown"))

	activateExercise := widget.NewCheck("Activate Exercise speech", func(b bool) {
		viper.Set("settings.audio.activateexercise", b)
		glob.Conf_Write()
	})
	activateExercise.SetChecked(viper.GetBool("settings.audio.activateexercise"))

	activatePause := widget.NewCheck("Activate Pause speech", func(b bool) {
		viper.Set("settings.audio.activatepause", b)
		glob.Conf_Write()
	})
	activatePause.SetChecked(viper.GetBool("settings.audio.activatepause"))

	// Create Form
	items := []*widget.FormItem{
		widget.NewFormItem("", activateAudio),
		widget.NewFormItem("", activateCountdown),
		widget.NewFormItem("", activateExercise),
		widget.NewFormItem("", activatePause),
	}
	form := widget.NewForm(items...)

	// Show Dialog
	dialog.ShowCustomConfirm("Settings", "Save", "Cancel", form, func(b bool) {
		if b {
			// Save settings (already done by checkbox callbacks)
		}
	}, w)
}
func workoutSettings() {

	// Create form elements for workout settings
	durationEntry := widget.NewEntry()
	durationEntry.SetPlaceHolder("Duration (minutes)")
	durationEntry.SetText(fmt.Sprintf("%.0f", glob.Conf.Workout.Duration))

	typeSelect := widget.NewSelect(glob.Choices_Type, func(s string) {})
	typeSelect.SetSelected(glob.Conf.Workout.Type)

	areaSelect := widget.NewSelect(glob.Choices_Area, func(s string) {})
	areaSelect.SetSelected(glob.Conf.Workout.Area)

	levelSelect := widget.NewSelect(glob.Choices_Level, func(s string) {})
	levelSelect.SetSelected(glob.Conf.Workout.Level)

	equipmentSelect := widget.NewSelect(glob.Choices_Equipment, func(s string) {})
	equipmentSelect.SetSelected(glob.Conf.Workout.Equipment)

	templateSelect := widget.NewSelect(glob.Choices_Template, func(s string) {})
	templateSelect.SetSelected(glob.Conf.Workout.Template)

	// Create a form with the workout settings
	form := widget.NewForm(
		widget.NewFormItem("Duration", durationEntry),
		widget.NewFormItem("Type", typeSelect),
		widget.NewFormItem("Area", areaSelect),
		widget.NewFormItem("Level", levelSelect),
		widget.NewFormItem("Equipment", equipmentSelect),
		widget.NewFormItem("Template", templateSelect),
	)

	// Show Dialog
	dialog.ShowCustomConfirm("Workout Settings", "Save", "Cancel", form, func(b bool) {
		if b {
			if d, err := strconv.ParseFloat(durationEntry.Text, 64); err == nil {
				viper.Set("workout.duration", d)
			}
			viper.Set("workout.type", typeSelect.Selected)
			viper.Set("workout.area", areaSelect.Selected)
			viper.Set("workout.level", levelSelect.Selected)
			viper.Set("workout.equipment", equipmentSelect.Selected)
			viper.Set("workout.template", templateSelect.Selected)

			glob.Conf_Write()
		}
	}, w)
}