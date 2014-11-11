package terminal

import (
	"fmt"
	"../lib"
	"github.com/wsxiaoys/terminal"
)

func InitTerminalStatusSelectLanguage() *TerminalStatusSelectLanguage {

	status := &TerminalStatusSelectLanguage{}
	return status
}

type TerminalStatusSelectLanguage struct{}

func (me *TerminalStatusSelectLanguage) Show(app *lib.App) TerminalStatusInterface {

	t := terminal.Stdout.
		Clear().Move(0, 0).
		Color("y").
		Print("Languages: ").Nl()
		
	t = t.Color("y").
		Print("-1 - ").
		Color("w").
		Print("Change user").Nl()

	t = t.Color("y").
		Print(" 0 - ").
		Color("w").
		Print("Exit").Nl()

	languages := app.GetLanguages()

	for i := range languages {

		t = t.Color("y").
			Print(" ").Print(i+1).Print(" - ").
			Color("w").Print(languages[i]).Nl()

	}

	terminal.Stdout.
		Color("y").
		Print("Language: ").
		Color("w")

	var selection int
	fmt.Scanf("%d", &selection)
	
	var nextStep TerminalStatusInterface;

	switch selection {
		case -1: nextStep = InitTerminalStatusSelectUser()
		case 0: nextStep = nil
		default: nextStep = InitTerminalStatusSelectProject(languages[selection-1])
	}

	return nextStep

}
