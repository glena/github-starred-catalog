package terminal

import (
	"fmt"
	"../lib"
	"github.com/wsxiaoys/terminal"
)

func InitTerminalStatusSelectProject(language string) *TerminalStatusSelectProject {

	status := &TerminalStatusSelectProject{
		Language: language,
	}
	return status
}

type TerminalStatusSelectProject struct {
	Language string
}

func (me *TerminalStatusSelectProject) Show(app *lib.App) TerminalStatusInterface {

	repos := app.GetRepositories(me.Language)

	t := terminal.Stdout.Clear().Move(0, 0).
		Nl().
		Color("w").
		Print(me.Language).
		Color("y").
		Print(" repositories: ").Nl()

	t = t.Color("y").
		Print("-1 - ").
		Color("w").
		Print("Go back").Nl()

	t = t.Color("y").
		Print(" 0 - ").
		Color("w").
		Print("Search").Nl()

	for i:=range(repos) {

	  t = t.
	        Color("y").
	        Print(" ").Print(i+1).Print(" - ").
	        Color("w").
	        Print(*repos[i].FullName).Nl().
	        Color("y").
	        Print("\t").
	        Print(*repos[i].Description).Nl().
	        Print("\t").
	        Print(*repos[i].HTMLURL).Nl()

	}
	
	terminal.Stdout.
		Color("y").
		Print("Project: ").
		Color("w")

	var selection int
	fmt.Scanf("%d", &selection)
	
	var nextStep TerminalStatusInterface;

	switch selection {
		case -1: nextStep = InitTerminalStatusSelectLanguage()
		case 0: nextStep = nil
		default: nextStep = nil
	}

	return nextStep
}
