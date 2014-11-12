package terminal

import (
	"fmt"
	"strings"
	"../lib"
	"github.com/wsxiaoys/terminal"
	"github.com/google/go-github/github"
)

func InitTerminalStatusSelectProject(language string, search string) *TerminalStatusSelectProject {

	status := &TerminalStatusSelectProject{
		Language: language,
		Search: search,
	}
	return status
}

type TerminalStatusSelectProject struct {
	Language string
	Search string
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

	if me.Search != ""{
		repos = me.FilterRepos(repos,me.Search)
	}

	for i:=range(repos) {

	  t = t.
	        Color("y").
	        Print(" ").Print(i+1).Print(" - ").
	        Color("w").
	        Print(*repos[i].FullName).Nl()

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
		case 0: nextStep = InitTerminalStatusSearchProject(me.Language)
		default: nextStep = InitTerminalStatusShowProject(me.Language, &repos[selection-1]) 
	}

	return nextStep
}

func (me *TerminalStatusSelectProject) FilterRepos(s []github.Repository, filter string) []github.Repository {
		var p []github.Repository // == nil
		for _, v := range s {
				if strings.Index(strings.ToLower(*v.FullName), strings.ToLower(filter)) >= 0 {
						p = append(p, v)
						fmt.Println(len(p))
				}
		}
		return p
}
