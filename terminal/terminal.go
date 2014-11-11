package terminal

import (
	"../lib"
)

func InitTerminal() *Terminal {

	terminal := &Terminal{
		Status: InitTerminalStatusSelectUser(),
	}
	return terminal
}

type TerminalStatusInterface interface {
	Show(app *lib.App) TerminalStatusInterface
}

type Terminal struct {
	Status TerminalStatusInterface
}

func (me *Terminal) Init(app *lib.App) {

	for me.Status != nil {
		me.Status = me.Status.Show(app)
	}

}
