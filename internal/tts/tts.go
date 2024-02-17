package tts

import (
	"math/rand"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
	"go.uber.org/zap"
)

// Re-Usable Speech
var speech = htgotts.Speech{Folder: "audio", Language: voices.English}

// Speak a stsring
func Speak(txt string) {

	zap.L().Debug("Speaking", zap.String("Text", txt))

	// Speak
	go speech.Speak(txt)

}

// Speak with one of random words
func SpeakRand(cat string) {

	// Prepare String
	txt := ""

	// Cases
	switch sp := cat; sp {
	case "motivation":
		txt = randomSlice([]string{"Come on!", "Go go go!", "All you can give!", "One more rep!"})
	case "half":
		txt = randomSlice([]string{"Half time!", "Fifty percent already done!"})
	case "play":
		txt = randomSlice([]string{"Go!", "Lets go!", "Start!", "All in!", "Go for it!"})
	case "stop":
		txt = randomSlice([]string{"Stop!", "Halted!", "Break!"})
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
