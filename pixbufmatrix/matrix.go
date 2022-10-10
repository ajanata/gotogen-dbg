package pixbufmatrix

import (
	"fmt"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Matrix struct {
	width     uint
	height    uint
	thickness uint8
	buf       *gdk.Pixbuf
	img       *gtk.Image
}

func New(width, height uint, thickness uint8) (*Matrix, error) {
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

func (m *Matrix) SetPixel(x, y uint, r, g, b uint8) {
	pix := m.buf.GetPixels()
	// number of bytes per row
	stride := m.buf.GetRowstride()
	// number of bytes per pixel
	nChan := uint(m.buf.GetNChannels())

	for i := uint(0); i < uint(m.thickness); i++ {
		yOffset := (y*(uint(m.thickness)+1) + i) * uint(stride)
		for j := uint(0); j < uint(m.thickness); j++ {
			xOffset := (x*(uint(m.thickness)+1) + j) * nChan
			pix[yOffset+xOffset] = r
			pix[yOffset+xOffset+1] = g
			pix[yOffset+xOffset+2] = b
		}
	}
}
