package terminal

import (
  "fmt"
  "../lib"
  "github.com/wsxiaoys/terminal"
)

func InitTerminalStatusSearchProject(language string) *TerminalStatusSearchProject {

  status := &TerminalStatusSearchProject{
    Language: language,
  }
  return status
}

type TerminalStatusSearchProject struct {
  Language string
}

func (me *TerminalStatusSearchProject) Show(app *lib.App) TerminalStatusInterface {

  terminal.Stdout.Clear().Move(0, 0).
    Nl().
    Color("y").
    Print("Search repositories on ").
    Color("w").
    Print(me.Language).
    Color("y").
    Print(": ").Color("w")

  var search string
  fmt.Scanf("%s", &search)

  return InitTerminalStatusSelectProject(me.Language, search)
}
