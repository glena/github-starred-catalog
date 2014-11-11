package lib

import (
	"github.com/google/go-github/github"
)

type TerminalInterface interface {
	Init(app *App)
}

func InitApp(terminal TerminalInterface) *App {

	app := &App{
		Client:   InitGHClient(),
		Terminal: terminal,
	}

	return app
}

type App struct {
	Client   *GHClient
	Terminal TerminalInterface
}

func (me *App) Init() {

	me.Terminal.Init(me)

}

func (me *App) LoadUser(username string) {

	me.Client.Load(username)

}

func (me *App) GetLanguages() []string {

	return me.Client.GetLanguages()

}

func (me *App) GetRepositories(language string) []github.Repository {

	return me.Client.GetRepositories(language)

}
