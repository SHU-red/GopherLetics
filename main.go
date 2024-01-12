package main

import (
	"github.com/SHU-red/GopherLetics.git/internal/global"
	"github.com/SHU-red/GopherLetics.git/internal/gui"
)

func main() {

	// Initialize + Load Configuration
	global.Conf_initConf()
	global.Conf_Load()

	// Run GUI
	gui.Main()

}
