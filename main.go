package main

import (
	"github.com/SHU-red/GopherLetics.git/internal/global"
	"github.com/SHU-red/GopherLetics.git/internal/gui"
)

func main() {

	// Initialize Configuration with default values
	global.Conf_initConf()
	// Overwrite Configuration from config-file
	global.Conf_Load()

	// Run GUI
	gui.Main()

}
