package gui

import (
	"fmt"
	"net/http"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/SHU-red/GopherLetics/internal/workout"
	"github.com/adrg/strutil/metrics"
	"go.uber.org/zap"
)

var image_window fyne.Window

func ShowImageWindow(a fyne.App) {
	image_window = a.NewWindow("Exercise Image")

	// Placeholder image
	img := canvas.NewImageFromResource(nil)
	license := widget.NewLabel("Images provided by everkinetic (CC-BY-SA-4.0)")
	content := container.NewBorder(nil, license, nil, nil, img)
	image_window.SetContent(content)

	image_window.Show()
}

func UpdateImage(exerciseName string) {
	bestMatch := findBestMatch(exerciseName)
	if bestMatch != nil && len(bestMatch.Img) > 0 {
		go func() {
			imagePath := bestMatch.Img[0]
			imageName := imagePath[strings.LastIndex(imagePath, "/")+1:]
			url := fmt.Sprintf("https://raw.githubusercontent.com/everkinetic/data/main/src/images-web/%s", imageName)
			resp, err := http.Get(url)
			if err != nil {
				zap.L().Error("failed to download image", zap.Error(err), zap.String("url", url))
				return
			}
			defer resp.Body.Close()

			img := canvas.NewImageFromReader(resp.Body, imageName)
			img.FillMode = canvas.ImageFillContain
			currentExerciseLabel := widget.NewLabel("Current active exercise: " + exerciseName)
			foundExerciseLabel := widget.NewLabel("Closest Everkinetic match: " + bestMatch.Title)
			license := widget.NewLabel("Images provided by everkinetic (CC-BY-SA-4.0)")
			infoBox := container.NewVBox(currentExerciseLabel, foundExerciseLabel, license)
			content := container.NewBorder(nil, infoBox, nil, nil, img)
			image_window.SetContent(content)
		}()
	}
}

func findBestMatch(exerciseName string) *workout.Exercise {
	var bestMatch *workout.Exercise
	highestScore := 0.0

	jw := metrics.NewJaroWinkler()
	jw.CaseSensitive = false

	for i, exercise := range workout.AllExercises {
		// Check if any primary muscle is in the exercise name
		primaryMatch := false
		for _, primaryMuscle := range exercise.Primary {
			if strings.Contains(strings.ToLower(exerciseName), strings.ToLower(primaryMuscle)) {
				primaryMatch = true
				break
			}
		}

		if !primaryMatch && len(exercise.Primary) > 0 {
			continue // Skip if no primary muscle matches and primary muscles are defined
		}

		score := jw.Compare(exerciseName, exercise.Title)
		if score > highestScore {
			highestScore = score
			bestMatch = &workout.AllExercises[i]
		}
	}

	return bestMatch
}
