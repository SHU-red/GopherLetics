package workout

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/kr/pretty"
	"github.com/spf13/viper"
)

// GopherLetics style of stepwise Workout
type Workout struct {
	Ty string // Type
	Na string // Name
	Du int    // Duration
}

// All Parsed Workouts
type Workouts []Workout

var Wo Workouts

// Fetch new Workout via HTTP
func (wo *Workouts) Fetch() {

	// Clear Workouts
	Wo = Workouts{}

	// Build URL from current config
	url := "https://api.sebhulse.com/v1/workout/?type=" + viper.GetString("workout.type") + "&area=" + viper.GetString("workout.area") + "&level=" + viper.GetString("workout.level") + "&duration=" + viper.GetString("workout.duration")

	// Debug
	println(url)

	// Get HTTP data
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	// We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Debug
	sb := string(body)
	fmt.Println(sb)

	// Get Lines
	lines := strings.Split(string(body), "\n")

	// Init counter for current workout line
	w := -1            // Workout number
	var w_buff Workout // buffer Workout
	var b int          // buffer int
	var bs string      // buffer string

	// Parse Lines into Workout
	for _, l := range lines {

		// Regex Pre-definition
		r_h, _ := regexp.Compile(".*\"(.*)\": \\[.*")    // Heading
		r_e, _ := regexp.Compile(".*\"ex\": \"(.*)\".*") // Exercise
		r_d, _ := regexp.Compile(".*\"du\": ([0-9]+).*") // Duration
		r_r, _ := regexp.Compile(".*\"re\": ([0-9]+).*") // Rest

		// If Heading
		if r_h.MatchString(l) {

			// Next Slice
			w++

			// Append
			Wo = append(Wo, w_buff)

			// Set Values
			Wo[w].Ty = "heading"
			Wo[w].Na = r_h.FindStringSubmatch(l)[1]
			Wo[w].Du = 0

		}

		// If Exercise
		if r_e.MatchString(l) {

			// Get Values
			bs = r_e.FindStringSubmatch(l)[1]

			// Next Slice
			w++

			// Append
			Wo = append(Wo, w_buff)

			// Treat Transition as Rest
			if bs == "transition" {

				Wo[w].Ty = "rest"
				Wo[w].Na = bs
				Wo[w].Du = 0

			} else {

				Wo[w].Ty = "exercise"
				Wo[w].Na = bs
				Wo[w].Du = 0

			}

		}

		// If Duration
		if r_d.MatchString(l) {

			// Set Values
			b, _ = strconv.Atoi(r_d.FindStringSubmatch(l)[1])
			Wo[w].Du = b

		}

		// If Rest
		if r_r.MatchString(l) {

			// Get Value
			b, _ = strconv.Atoi(r_r.FindStringSubmatch(l)[1])

			// If Rest not zero
			if b > 0 {

				w++

				// Append
				Wo = append(Wo, w_buff)

				// Set Values
				Wo[w].Ty = "rest"
				Wo[w].Na = "Rest"
				Wo[w].Du = b

			}

		}

	}

	//Debug
	pretty.Println("Collected Workouts: ")
	pretty.Println(Wo)

}
