package pixbufmatrix

import (
	"fmt"
	"image/color"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"tinygo.org/x/drivers"
)

type Matrix struct {
	width     int16
	height    int16
	thickness int16
	buf       *gdk.Pixbuf
	img       *gtk.Image
}

var _ drivers.Displayer = (*Matrix)(nil)

func New(width, height int16, thickness int16) (*Matrix, error) {
	buf, err := gdk.PixbufNew(gdk.COLORSPACE_RGB, false, 8, int(width)*(int(thickness)+1), int(height)*(int(thickness)+1))
	if err != nil {
		return nil, fmt.Errorf("creating buf: %w", err)
	}

	img, err := gtk.ImageNewFromPixbuf(buf)
	if err != nil {
		return nil, fmt.Errorf("creating image: %w", err)
	}

	m := &Matrix{
		width:     width,
		height:    height,
		thickness: thickness,
		buf:       buf,
		img:       img,
	}

	return m, nil
}

func (m *Matrix) Widget() gtk.IWidget {
	return m.img
}

func (m *Matrix) Show() {
	m.img.Show()
}

func (m *Matrix) Hide() {
	m.img.Hide()
}

func (m *Matrix) QueueDraw() {
	m.img.SetFromPixbuf(m.buf)
	m.img.QueueDraw()
}

// implementation of drivers.Displayer

func (m *Matrix) Size() (x, y int16) {
	return m.width, m.height
}

func (m *Matrix) SetPixel(x, y int16, c color.RGBA) {
	pix := m.buf.GetPixels()
	// number of bytes per row
	stride := m.buf.GetRowstride()
	// number of bytes per pixel
	nChan := m.buf.GetNChannels()

	for i := 0; i < int(m.thickness); i++ {
		yOffset := (int(y)*int(m.thickness+1) + i) * stride
		for j := 0; j < int(m.thickness); j++ {
			xOffset := (int(x)*int(m.thickness+1) + j) * nChan
			offset := yOffset + xOffset
			if offset > len(pix) || offset < 0 {
				continue
			}
			pix[offset] = c.R
			pix[offset+1] = c.G
			pix[offset+2] = c.B
		}
	}
}

func (m *Matrix) Display() error {
	m.QueueDraw()
	return nil
}

func (*Matrix) CanUpdateNow() bool { return true }
