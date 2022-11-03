package simulator

import (
	"sync"

	"github.com/ajanata/gotogen"
	"github.com/ajanata/textbuf"
	"tinygo.org/x/drivers"

	"github.com/ajanata/gotogen-simulator/pixbufmatrix"
)

type Gotogen struct {
	Face *pixbufmatrix.Matrix
	Menu *pixbufmatrix.Matrix

	buttonLock    sync.Mutex
	pendingButton gotogen.MenuButton
}

var _ gotogen.Driver = (*Gotogen)(nil)
var _ gotogen.MenuInput = (*Gotogen)(nil)

func (g *Gotogen) EarlyInit() (faceDisplay drivers.Displayer, menuInput gotogen.MenuInput, boopSensor gotogen.BoopSensor, err error) {
	return g.Face, g, nil, nil
}

func (g *Gotogen) LateInit(_ *textbuf.Buffer) error {
	return nil
}

func (g *Gotogen) PressedButton() gotogen.MenuButton {
	g.buttonLock.Lock()
	b := g.pendingButton
	g.pendingButton = gotogen.MenuButtonNone
	g.buttonLock.Unlock()
	return b
}

func (g *Gotogen) ButtonPress(b gotogen.MenuButton) {
	g.buttonLock.Lock()
	g.pendingButton = b
	g.buttonLock.Unlock()
}
