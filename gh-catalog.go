package main

import (
	"./lib"
	"./terminal"
)

func main() {
	terminal := terminal.InitTerminal()

	app := lib.InitApp(terminal)
	app.Init()

}
