package lib

//https://godoc.org/github.com/wsxiaoys/terminal

import (
  "fmt"
  "github.com/wsxiaoys/terminal"
  //"github.com/wsxiaoys/terminal/color"
)

func InitTerminal() *Terminal {

  terminal:= &Terminal{}
  return terminal
}

type Terminal struct {}

func (me *Terminal) ShowWelcome(app *App) string {

  terminal.Stdout.Clear().
      Color("y").
      Print("Username: ")

  var username string
  fmt.Scanf("%s", &username)

  terminal.Stdout.
      Nl().
      Color("y").
      Print("Loading: ").
      Print(username).Nl()

  return username
}


func (me *Terminal) ShowLanguages(languages []string) {

  t := terminal.Stdout.
      Nl().
      Color("y").
      Print("Languages: ").Nl()

  for i:=range(languages) {

    t = t.Print(i).Print(" - ").Print(languages[i]).Nl()

  }
}
