package tts

import (
	"math/rand"
	"os"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
	"github.com/SHU-red/GopherLetics.git/internal/glob"
	"go.uber.org/zap"
)

// Re-Usable Speech
var speech htgotts.Speech

// Initialize Speech
func SpeakInit() {

	// Get Temporary Folder
	tmp_folder := os.TempDir() + "/gopherletics/audio"

	speech = htgotts.Speech{Folder: tmp_folder, Language: voices.English}
}

// Speak a stsring
func Speak(txt string) {

	if !glob.Conf.Settings.Audio.Activate || !glob.Conf.Settings.Audio.ActivateExercise {
		return
	}

	zap.L().Debug("Speaking", zap.String("Text", txt))

	// Speak
	go speech.Speak(txt)

}

// Speak with one of random words
func SpeakRand(cat string) {

	if !glob.Conf.Settings.Audio.Activate {
		return
	}

	// Prepare String
	txt := ""

	// Cases
	switch sp := cat; sp {
	case "motivation":
		if !glob.Conf.Settings.Audio.ActivateExercise {
			return
		}
		txt = randomSlice([]string{"Come on!", "Go go go!", "All you can give!", "One more rep!"})
	case "half":
		if !glob.Conf.Settings.Audio.ActivateCountdown {
			return
		}
		txt = randomSlice([]string{"Half time!", "Fifty percent already done!"})
	case "play":
		if !glob.Conf.Settings.Audio.ActivateExercise {
			return
		}
		txt = randomSlice([]string{"Go!", "Lets go!", "Start!", "All in!", "Go for it!"})
	case "stop":
		if !glob.Conf.Settings.Audio.ActivatePause {
			return
		}
		txt = randomSlice([]string{"Stop!", "Halted!", "Break!"})
	case "done":
		if !glob.Conf.Settings.Audio.ActivateExercise {
			return
		}
		txt = randomSlice([]string{"You did it!", "Excellent!", "Workout successfully finished!"})
	default:
		txt = "Voice command error"
	}

	// Loop through sp
	go speech.Speak(txt)

}

// Pick a random Slice
func randomSlice(sli []string) string {
	randomIndex := rand.Intn(len(sli))
	return sli[randomIndex]
}
