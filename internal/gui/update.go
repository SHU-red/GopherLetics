package gui

import (
	"fmt"

	"github.com/SHU-red/GopherLetics.git/internal/global"
)

func updatevals() {

}

// Update the Shown timer String
func update_timer_str() {
	time, _ := global.Gui.Timer.Get()
	timer.Text = fmt.Sprintf("%04d", time)

	// Update the shown timer
	timer.Refresh()

}

// Refresh workout and reset all necessary values
func update_all() {

	// Update Timer
	update_timer_str()

}
