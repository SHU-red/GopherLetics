package gui

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

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

// App version (set from FyneApp.toml metadata at startup)
var appVersion string

// Update check state
var (
	updateAvailable bool
	updateTag       string
	updateChecked   bool
)

func Main() {

	// Initialize glob variables
	glob.Gui_initval()

	// Fyne App
	a := app.NewWithID("com.gopherletics.app")
	a.Settings().SetTheme(&fitnessTheme{})
	w = a.NewWindow("GopherLetics")
	w.Resize(fyne.NewSize(1000, 750))

	// Store version from FyneApp.toml
	appVersion = a.Metadata().Version

	// Create shared exercise widgets
	createExerciseWidgets()

	// Start background update check
	go checkUpdate()

	// URLs
	url_gopherletics, _ := url.Parse("https://github.com/SHU-red/GopherLetics")
	url_fyne, _ := url.Parse("https://fyne.io/")
	url_golang, _ := url.Parse("https://go.dev/")
	url_bmc, _ := url.Parse("https://buymeacoffee.com/yffbptmtaa")
	url_releases, _ := url.Parse("https://github.com/SHU-red/GopherLetics/releases")

	// Top title
	top := container.NewCenter(widget.NewLabelWithStyle("GopherLetics", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))

	// Bottom footer
	footerLinks := container.NewHBox(
		widget.NewHyperlink(fmt.Sprintf("GopherLetics %s", appVersion), url_gopherletics),
		widget.NewLabel("|"),
		widget.NewHyperlink("Fyne", url_fyne),
		widget.NewLabel("|"),
		widget.NewHyperlink("Go", url_golang),
		widget.NewLabel("|"),
		widget.NewHyperlink("☕ Support", url_bmc),
	)

	// Update indicator — shown when check completes
	updateLink := widget.NewHyperlink("", url_releases)
	updateLink.Hidden = true

	footer := container.NewVBox(
		footerLinks,
		container.NewCenter(updateLink),
	)

	// Poll update status and show the link when ready
	go func() {
		for range time.Tick(500 * time.Millisecond) {
			if updateChecked {
				fyne.Do(func() {
					if updateAvailable {
						updateLink.SetText("⬆ Update available: " + updateTag)
						updateLink.Hidden = false
					}
				})
				return
			}
		}
	}()

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

// checkUpdate fetches the latest release version from GitHub.
func checkUpdate() {
	resp, err := http.Get("https://api.github.com/repos/SHU-red/GopherLetics/releases/latest")
	if err != nil {
		zap.L().Debug("update check failed", zap.Error(err))
		updateChecked = true
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		updateChecked = true
		return
	}

	var rel struct {
		TagName string `json:"tag_name"`
	}
	if err := json.Unmarshal(body, &rel); err != nil {
		updateChecked = true
		return
	}

	local := "v" + appVersion
	if compareVersions(rel.TagName, local) > 0 {
		updateAvailable = true
		updateTag = rel.TagName
	}
	updateChecked = true
}

// compareVersions returns 1 if v1 > v2, -1 if v1 < v2, 0 if equal.
// Expects "v1.2.3" format.
func compareVersions(v1, v2 string) int {
	v1 = strings.TrimPrefix(v1, "v")
	v2 = strings.TrimPrefix(v2, "v")
	p1 := strings.Split(v1, ".")
	p2 := strings.Split(v2, ".")
	for i := range 3 {
		var n1, n2 int
		if i < len(p1) {
			n1, _ = strconv.Atoi(p1[i])
		}
		if i < len(p2) {
			n2, _ = strconv.Atoi(p2[i])
		}
		if n1 > n2 {
			return 1
		}
		if n1 < n2 {
			return -1
		}
	}
	return 0
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

	activateAudio := widget.NewCheck("", func(b bool) {
		viper.Set("settings.audio.activate", b)
		glob.Conf_Write()
	})
	activateAudio.SetChecked(viper.GetBool("settings.audio.activate"))

	activateCountdown := widget.NewCheck("", func(b bool) {
		viper.Set("settings.audio.activatecountdown", b)
		glob.Conf_Write()
	})
	activateCountdown.SetChecked(viper.GetBool("settings.audio.activatecountdown"))

	activateExercise := widget.NewCheck("", func(b bool) {
		viper.Set("settings.audio.activateexercise", b)
		glob.Conf_Write()
	})
	activateExercise.SetChecked(viper.GetBool("settings.audio.activateexercise"))

	activatePause := widget.NewCheck("", func(b bool) {
		viper.Set("settings.audio.activatepause", b)
		glob.Conf_Write()
	})
	activatePause.SetChecked(viper.GetBool("settings.audio.activatepause"))

	// Row: label left, checkbox right
	row := func(label string, check *widget.Check) *fyne.Container {
		return container.NewBorder(nil, nil, widget.NewLabel(label), nil, check)
	}

	audioCard := widget.NewCard("Audio Feedback", "",
		container.NewVBox(
			row("Voice prompts", activateAudio),
			widget.NewSeparator(),
			row("Countdown numbers", activateCountdown),
			widget.NewSeparator(),
			row("Exercise names", activateExercise),
			widget.NewSeparator(),
			row("Pause announcements", activatePause),
		),
	)

	// About section
	url_bmc, _ := url.Parse("https://buymeacoffee.com/yffbptmtaa")
	url_gh, _ := url.Parse("https://github.com/SHU-red/GopherLetics")
	url_releases, _ := url.Parse("https://github.com/SHU-red/GopherLetics/releases")

	updateStatus := "Checking..."
	if updateChecked {
		if updateAvailable {
			updateStatus = "⬆ " + updateTag + " available"
		} else {
			updateStatus = "✓ Up to date"
		}
	}

	aboutContent := container.NewVBox(
		widget.NewLabel("Version: "+appVersion),
		container.NewHBox(
			widget.NewLabel("Updates:"),
			widget.NewHyperlink(updateStatus, url_releases),
		),
		widget.NewSeparator(),
		container.NewHBox(
			widget.NewHyperlink("GitHub", url_gh),
			widget.NewLabel("|"),
			widget.NewHyperlink("☕ Buy me a coffee", url_bmc),
		),
	)
	aboutCard := widget.NewCard("About", "", aboutContent)

	content := container.NewVBox(audioCard, widget.NewSeparator(), aboutCard)

	d := dialog.NewCustomConfirm("Settings", "Close", "", container.NewScroll(content), func(b bool) {}, w)
	d.Show()
	d.Resize(fyne.NewSize(640, 520))
}
func workoutSettings() {

	// Build interactive controls
	durationEntry := widget.NewEntry()
	durationEntry.SetPlaceHolder("Duration (minutes)")
	durationEntry.SetText(fmt.Sprintf("%.0f", glob.Conf.Workout.Duration))

	typeSelect := widget.NewSelect(glob.Choices_Type, nil)
	typeSelect.SetSelected(glob.Conf.Workout.Type)

	areaSelect := widget.NewSelect(glob.Choices_Area, nil)
	areaSelect.SetSelected(glob.Conf.Workout.Area)

	levelSelect := widget.NewSelect(glob.Choices_Level, nil)
	levelSelect.SetSelected(glob.Conf.Workout.Level)

	equipmentCheck := widget.NewCheckGroup(glob.Choices_Equipment, nil)
	equipmentCheck.SetSelected(glob.Conf.Workout.Equipment)

	templateSelect := widget.NewSelect(glob.Choices_Template, nil)
	templateSelect.SetSelected(glob.Conf.Workout.Template)

	// Card-based layout
	content := container.NewVBox(
		widget.NewCard("Duration", "", durationEntry),
		widget.NewSeparator(),
		widget.NewCard("Type", "", typeSelect),
		widget.NewSeparator(),
		widget.NewCard("Area", "", areaSelect),
		widget.NewSeparator(),
		widget.NewCard("Level", "", levelSelect),
		widget.NewSeparator(),
		widget.NewCard("Equipment", "Select one or more", equipmentCheck),
		widget.NewSeparator(),
		widget.NewCard("Template", "", templateSelect),
	)
	d := dialog.NewCustomConfirm("Workout Settings", "Save", "Cancel",
		container.NewScroll(content),
		func(b bool) {
			if !b {
				return
			}
			if d, err := strconv.ParseFloat(durationEntry.Text, 64); err == nil {
				viper.Set("workout.duration", d)
			}
			viper.Set("workout.type", typeSelect.Selected)
			viper.Set("workout.area", areaSelect.Selected)
			viper.Set("workout.level", levelSelect.Selected)
			viper.Set("workout.equipment", equipmentCheck.Selected)
			viper.Set("workout.template", templateSelect.Selected)
			glob.Conf_Write()
		}, w)
	d.Show()
	d.Resize(fyne.NewSize(640, 560))
}