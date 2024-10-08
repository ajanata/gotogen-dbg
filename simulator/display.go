package simulator

import (
	"image"
	"image/color"
	"sync"

	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"github.com/ajanata/gotogen"
)

type Display struct {
	lock      *sync.RWMutex
	width     int16
	height    int16
	thickness int16
	data      [][]color.RGBA
}

var _ gotogen.Display = (*Display)(nil)

func NewDisplay(width, height, thickness int16) *Display {
	d := &Display{
		lock:      &sync.RWMutex{},
		width:     width,
		height:    height,
		thickness: thickness,
		data:      make([][]color.RGBA, width),
	}

	for x := range width {
		d.data[x] = make([]color.RGBA, height)
	}

	return d
}

func (d *Display) Size() (x, y int16) {
	return d.width, d.height
}

func (d *Display) SetPixel(x, y int16, c color.RGBA) {
	c.A = 255
	d.lock.Lock()
	d.data[x][y] = c
	d.lock.Unlock()
}

func (d *Display) Display() error {
	// no-op
	return nil
}

func (d *Display) CanUpdateNow() bool {
	return true
}

func (d *Display) Render(ops *op.Ops) {
	d.lock.RLock()
	for x := range d.width {
		for y := range d.height {
			c := d.data[x][y]
			func() {
				pt := d.thickness + 1
				px := int(x * pt)
				py := int(y * pt)
				defer op.Offset(image.Pt(px, py)).Push(ops).Pop()
				defer clip.Rect{Max: image.Pt(int(pt), int(pt))}.Push(ops).Pop()
				paint.ColorOp{Color: color.NRGBA(c)}.Add(ops)
				paint.PaintOp{}.Add(ops)
			}()
		}
	}
	d.lock.RUnlock()
}
