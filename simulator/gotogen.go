package simulator

import (
	"sync"
	"time"

	"github.com/ajanata/textbuf"

	"github.com/ajanata/gotogen"
)

type Gotogen struct {
	// Face *pixbufmatrix.Matrix
	// Menu *pixbufmatrix.Matrix
	Face *Display
	Menu *Display

	lock          sync.Mutex
	pendingButton gotogen.MenuButton
	accX          int32
	accY          int32
	accZ          int32
	boop          uint8
	talk          bool
}

var _ gotogen.Driver = (*Gotogen)(nil)

func (g *Gotogen) EarlyInit() (faceDisplay gotogen.Display, err error) {
	return g.Face, nil
}

func (g *Gotogen) LateInit(_ *textbuf.Buffer) {
	time.Sleep(time.Second)
}

func (g *Gotogen) PressedButton() gotogen.MenuButton {
	g.lock.Lock()
	b := g.pendingButton
	g.pendingButton = gotogen.MenuButtonNone
	g.lock.Unlock()
	return b
}

func (g *Gotogen) ButtonPress(b gotogen.MenuButton) {
	g.lock.Lock()
	g.pendingButton = b
	g.lock.Unlock()
}

func (g *Gotogen) UpdateAccelerometerX(v float64) {
	g.lock.Lock()
	g.accX = int32(v * 1000)
	g.lock.Unlock()
}

func (g *Gotogen) UpdateAccelerometerY(v float64) {
	g.lock.Lock()
	g.accY = int32(v * 1000)
	g.lock.Unlock()
}

func (g *Gotogen) UpdateAccelerometerZ(v float64) {
	g.lock.Lock()
	g.accZ = int32(v * 1000)
	g.lock.Unlock()
}

func (g *Gotogen) UpdateBoop(v float64) {
	g.lock.Lock()
	g.boop = uint8(255 * v)
	g.lock.Unlock()
}

func (g *Gotogen) UpdateTalk(v bool) {
	g.lock.Lock()
	g.talk = v
	g.lock.Unlock()
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
	g.lock.Lock()
	defer g.lock.Unlock()
	return g.boop, gotogen.SensorStatusAvailable
}

func (g *Gotogen) Accelerometer() (int32, int32, int32, gotogen.SensorStatus) {
	g.lock.Lock()
	defer g.lock.Unlock()
	return g.accX, g.accY, g.accZ, gotogen.SensorStatusAvailable
}

func (g *Gotogen) Talking() bool {
	g.lock.Lock()
	defer g.lock.Unlock()
	return g.talk
}

func (g *Gotogen) StatusLine() string {
	return ""
}
