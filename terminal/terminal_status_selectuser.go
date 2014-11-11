package terminal

import (
	"fmt"
	"../lib"
	"github.com/wsxiaoys/terminal"
)

func InitTerminalStatusSelectUser() *TerminalStatusSelectUser {

	status := &TerminalStatusSelectUser{}
	return status
}

type TerminalStatusSelectUser struct{}

func (me *TerminalStatusSelectUser) Show(app *lib.App) TerminalStatusInterface {

	terminal.Stdout.Clear().Move(0, 0).
		Color("y").
		Print("Username: ").
		Color("w")

	var username string
	fmt.Scanf("%s", &username)

	terminal.Stdout.
		Nl().
		Color("y").
		Print("Loading: ").
		Print(username).
		Print("...").Nl()

	app.LoadUser(username)

	return InitTerminalStatusSelectLanguage()
}
