package lib

func InitApp(Username string) *App {

  app := &App{
    Client:InitGHClient(),
    Terminal: InitTerminal(),
  }

  return app
}

type App struct {

  Client *GHClient
  Terminal *Terminal

}

func (me *App) Init() {

  username := me.Terminal.ShowWelcome(me)
  me.Client.Load(username)
  me.Terminal.ShowLanguages(me.Client.GetLanguages())
}
