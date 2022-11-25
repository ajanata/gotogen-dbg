package simulator

import (
	"sync"

	"github.com/ajanata/textbuf"
	"tinygo.org/x/drivers"

	"github.com/ajanata/gotogen"

	"github.com/ajanata/gotogen-simulator/pixbufmatrix"
)

type Gotogen struct {
	Face *pixbufmatrix.Matrix
	Menu *pixbufmatrix.Matrix

	buttonLock    sync.Mutex
	pendingButton gotogen.MenuButton
}

var _ gotogen.Driver = (*Gotogen)(nil)

func (g *Gotogen) EarlyInit() (faceDisplay drivers.Displayer, err error) {
	return g.Face, nil
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

func (g *Gotogen) MenuItems() []gotogen.Item {
	return []gotogen.Item{
		&gotogen.ActionItem{
			Name:   "No config for sim",
			Invoke: func() {},
		},
	}
}

func (g *Gotogen) BoopDistance() (uint8, gotogen.SensorStatus) {
	// TODO simulator boop sensor
	return 0, gotogen.SensorStatusUnavailable
}

func (g *Gotogen) Accelerometer() (int32, int32, int32, gotogen.SensorStatus) {
	// TODO simulator accelerometer
	return 0, 0, 0, gotogen.SensorStatusUnavailable
}
