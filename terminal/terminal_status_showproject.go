package terminal

import (
  "fmt"
  "../lib"
  "github.com/wsxiaoys/terminal"
  "github.com/google/go-github/github"
)

func InitTerminalStatusShowProject(language string, repo *github.Repository) *TerminalStatusShowProject {

  status := &TerminalStatusShowProject{
    Language: language,
    Repo: repo,
  }
  return status
}

type TerminalStatusShowProject struct {
  Language string
  Repo *github.Repository
}

func (me *TerminalStatusShowProject) Show(app *lib.App) TerminalStatusInterface {

  t:= terminal.Stdout.Clear().Move(0, 0).Color("y").Print(*me.Repo.FullName)

  if *me.Repo.Fork {
    t.Color("y").Print("(")
    t.Color("w").Print("fork")
    t.Color("y").Print("(").Color("w")
  }
  t.Nl()
  t.Color("w").Print(*me.Repo.HTMLURL).Nl().Nl()

  t.Print(*me.Repo.Homepage).Nl()
  t.Print(*me.Repo.Description).Nl()

  t.Nl().Nl().Nl().Color("y").Print("Press enter to go back")

  var search string
  fmt.Scanf("%s", &search)

  return InitTerminalStatusSelectProject(me.Language, "")
}
