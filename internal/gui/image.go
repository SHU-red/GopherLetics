package gui

import (
	"bytes"
	"image"
	"io"
	"net/http"
	"sync"
	"time"

	_ "image/jpeg"
	_ "image/png"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/SHU-red/GopherLetics/internal/workout"
	"github.com/adrg/strutil/metrics"
	"go.uber.org/zap"
	"strings"
)

const freeExerciseDBBase = "https://raw.githubusercontent.com/yuhonas/free-exercise-db/main/exercises/"

// Shared widgets placed in the main layout
var exerciseImage *canvas.Image
var exerciseNameLabel *widget.Label

var placeholderImage image.Image

// Animation state — two frames per exercise cycled as a flipbook
var (
	animMu    sync.Mutex
	animStop  chan struct{}  // close to stop current animation goroutine
	frame0    image.Image    // start position
	frame1    image.Image    // end position
	hasFrames bool           // true when both frames are ready
)

func init() {
	placeholderImage = image.NewRGBA(image.Rect(0, 0, 1, 1))
}

func createExerciseWidgets() {
	exerciseNameLabel = widget.NewLabel("Ready")
	exerciseImage = canvas.NewImageFromImage(placeholderImage)
	exerciseImage.FillMode = canvas.ImageFillContain
}

// stopAnimation halts any running flipbook goroutine.
func stopAnimation() {
	animMu.Lock()
	defer animMu.Unlock()
	if animStop != nil {
		close(animStop)
		animStop = nil
	}
	hasFrames = false
	frame0 = nil
	frame1 = nil
}

// startAnimation begins cycling between frame0 / frame1 on an 800 ms tick.
func startAnimation() {
	animMu.Lock()
	ch := make(chan struct{})
	animStop = ch
	f0 := frame0
	f1 := frame1
	animMu.Unlock()

	go func() {
		ticker := time.NewTicker(800 * time.Millisecond)
		defer ticker.Stop()

		showFirst := true
		fyne.Do(func() {
			exerciseImage.Image = f0
			exerciseImage.Refresh()
		})

		for {
			select {
			case <-ch:
				return
			case <-ticker.C:
				showFirst = !showFirst
				frame := f0
				if !showFirst {
					frame = f1
				}
				fyne.Do(func() {
					exerciseImage.Image = frame
					exerciseImage.Refresh()
				})
			}
		}
	}()
}

func UpdateImage(exerciseName string) {
	bestMatch := findBestMatch(exerciseName)

	// Always update the label
	fyne.Do(func() {
		exerciseNameLabel.SetText(exerciseName)
	})

	// Stop any running animation
	stopAnimation()

	if bestMatch == nil || len(bestMatch.Images) == 0 {
		fyne.Do(func() {
			exerciseImage.Image = placeholderImage
			exerciseImage.Refresh()
		})
		return
	}

	// Download frames in background
	go func() {
		urls := make([]string, len(bestMatch.Images))
		for i, p := range bestMatch.Images {
			urls[i] = freeExerciseDBBase + p
		}

		decoded := make([]image.Image, 0, len(urls))
		for _, url := range urls {
			resp, err := http.Get(url)
			if err != nil {
				zap.L().Error("failed to download image", zap.Error(err), zap.String("url", url))
				return
			}

			data, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				zap.L().Error("failed to read image", zap.Error(err), zap.String("url", url))
				return
			}

			img, _, err := image.Decode(bytes.NewReader(data))
			if err != nil {
				zap.L().Error("failed to decode image", zap.Error(err), zap.String("url", url))
				return
			}
			decoded = append(decoded, img)
		}

		if len(decoded) == 0 {
			return
		}

		// Store frames under lock
		animMu.Lock()
		frame0 = decoded[0]
		hasFrames = false
		if len(decoded) > 1 {
			frame1 = decoded[1]
			hasFrames = true
		} else {
			frame1 = decoded[0]
		}
		animMu.Unlock()

		if hasFrames {
			startAnimation()
		} else {
			fyne.Do(func() {
				exerciseImage.Image = decoded[0]
				exerciseImage.Refresh()
			})
		}
	}()
}

func findBestMatch(exerciseName string) *workout.Exercise {
	var bestMatch *workout.Exercise
	highestScore := 0.0

	jw := metrics.NewJaroWinkler()
	jw.CaseSensitive = false

	for i, exercise := range workout.AllExercises {
		muscleBoost := 0.0
		for _, m := range exercise.PrimaryMuscles {
			if strings.Contains(strings.ToLower(exerciseName), strings.ToLower(m)) {
				muscleBoost = 0.15
				break
			}
		}

		baseScore := jw.Compare(exerciseName, exercise.Name)
		score := baseScore + muscleBoost

		if score > highestScore {
			highestScore = score
			bestMatch = &workout.AllExercises[i]
		}
	}

	return bestMatch
}
