package main

import (
	"image/color"
	"math/rand"
	"os"

	"github.com/ajanata/gotogen"
	font "github.com/ajanata/oled_font"
	"github.com/ajanata/textbuf"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"tinygo.org/x/drivers"

	"github.com/ajanata/gotogen-simulator/pixbufmatrix"
)

const (
	faceWidth     = 128
	faceHeight    = 32
	faceThickness = 10

	menuWidth     = 128
	menuHeight    = 64
	menuThickness = 3
)

func main() {
	app, err := gtk.ApplicationNew("at.elbrarc.demo", glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		panic(err)
	}

	app.Connect("activate", func() {
		faceWindow, err := gtk.ApplicationWindowNew(app)
		if err != nil {
			panic(err)
		}
		faceWindow.SetTitle("Gotogen Debug - Face")
		faceWindow.SetDefaultSize(faceWidth*(faceThickness+1), faceHeight*(faceThickness+1))
		faceWindow.SetResizable(false)
		faceWindow.Show()
		faceWindow.Connect("destroy", func() {
			os.Exit(0)
		})

		faceMatrix, err := pixbufmatrix.New(faceWidth, faceHeight, faceThickness)
		if err != nil {
			panic(err)
		}
		faceWindow.Add(faceMatrix.Widget())
		faceMatrix.Show()

		for x := int16(0); x < faceWidth; x++ {
			for y := int16(0); y < faceHeight; y++ {
				faceMatrix.SetPixel(x, y, color.RGBA{uint8(x * 2), uint8(y * 8), 0, 0})
			}
		}

		faceMatrix.SetPixel(0, 0, color.RGBA{0xff, 0, 0, 0})
		faceMatrix.SetPixel(1, 0, color.RGBA{0, 0xff, 0, 0})
		faceMatrix.SetPixel(2, 0, color.RGBA{0, 0, 0xff, 0})

		font.PixelOff = color.RGBA{0, 0, 0, 255}
		font.PixelOn = color.RGBA{255, 255, 255, 255}
		buf, err := textbuf.New(faceMatrix, textbuf.FontSize7x10)
		if err != nil {
			panic(err)
		}
		buf.Println("hello, world!")

		// menu screen
		menuWindow, err := gtk.ApplicationWindowNew(app)
		if err != nil {
			panic(err)
		}
		menuWindow.SetTitle("Gotogen Debug - Menu")
		menuWindow.SetDefaultSize(menuWidth*(menuThickness+1), menuHeight*(menuThickness+1))
		menuWindow.Show()
		menuWindow.SetResizable(false)
		menuWindow.Connect("destroy", func() {
			os.Exit(0)
		})

		menuMatrix, err := pixbufmatrix.New(menuWidth, menuHeight, menuThickness)
		if err != nil {
			panic(err)
		}
		menuWindow.Add(menuMatrix.Widget())
		menuMatrix.Show()

		for x := int16(0); x < menuWidth; x++ {
			for y := int16(0); y < menuHeight; y++ {
				menuMatrix.SetPixel(x, y, color.RGBA{uint8(x * 2), uint8(y * 4), 0, 0})
			}
		}

		// input buttons
		inputWindow, err := gtk.ApplicationWindowNew(app)
		if err != nil {
			panic(err)
		}
		inputWindow.SetTitle("Gotogen Debug - Input")
		inputWindow.SetDefaultSize(200, 100)
		inputWindow.Connect("destroy", func() {
			os.Exit(0)
		})
		inputGrid, err := gtk.GridNew()
		if err != nil {
			panic(err)
		}
		buttonBack := mustMakeButton("Back")
		buttonUp := mustMakeButton("Up")
		buttonDown := mustMakeButton("Down")
		buttonLeft := mustMakeButton("Left")
		buttonRight := mustMakeButton("Right")
		buttonEnter := mustMakeButton("Enter")
		inputGrid.Attach(buttonUp, 1, 0, 1, 1)
		inputGrid.Attach(buttonDown, 1, 2, 1, 1)
		inputGrid.Attach(buttonLeft, 0, 1, 1, 1)
		inputGrid.Attach(buttonRight, 2, 1, 1, 1)
		inputGrid.Attach(buttonBack, 0, 0, 1, 1)
		inputGrid.Attach(buttonEnter, 1, 1, 1, 1)
		inputWindow.Add(inputGrid)
		inputWindow.ShowAll()

		// f, err := os.Open("Elbrarmemestickerscant.png")
		// if err != nil {
		// 	panic(err)
		// }
		// b := make([]uint16, 64*32)
		// var w, h int16
		// png.SetCallback(b, func(data []uint16, x, y, w, h, width, height int16) {
		// 	w, h = width, height
		// })
		// p, err := png.Decode(f)
		// if err != nil {
		// 	panic(err)
		// }
		// // r := p.Bounds()
		// r := image.Rect(0, 0, int(w), int(h))
		// for x := r.Min.X; x < r.Max.X; x++ {
		// 	xx := int16(x - r.Min.X)
		// 	if xx > 128 {
		// 		break
		// 	}
		// 	for y := r.Min.Y; y < r.Max.Y; y++ {
		// 		yy := int16(y - r.Min.Y)
		// 		if yy > 32 {
		// 			break
		// 		}
		// 		r, g, b, a := p.At(x, y).RGBA()
		// 		faceMatrix.SetPixel(xx, yy, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		// 	}
		// }
		// faceMatrix.Display()
		// f, err := os.Open("Untitled.png")
		// 24-bit non-transparent
		// f, err := os.Open("Elbrarmemestickerscant3.png")
		// if err != nil {
		// 	panic(err)
		// }
		// // row, col
		// b := make([][]uint16, 32)
		// for i := 0; i < 32; i++ {
		// 	b[i] = make([]uint16, 64)
		// }
		// var bb [64]uint16
		// w, h := uint16(64), uint16(32)
		// png.SetCallback(bb[:], func(data []uint16, x, y, w, h, width, height int16) {
		// 	// w, h = width, height
		// 	// b[x] = make([]uint16, 32)
		// 	// b[y] = data
		// 	copy(b[y], data[:64])
		// 	// b[x][y] = bb[0]
		// })
		// _, err = png.Decode(f)
		// if err != nil {
		// 	panic(err)
		// }
		// // r := p.Bounds()
		// r := image.Rect(0, 0, int(w), int(h))
		// for y := r.Min.Y; y < r.Max.Y; y++ {
		// 	yy := int16(y - r.Min.Y)
		// 	if yy > 32 {
		// 		break
		// 	}
		// 	for x := r.Min.X; x < r.Max.X; x++ {
		// 		xx := int16(x - r.Min.X)
		// 		if xx > 64 {
		// 			break
		// 		}
		// 		// re := ((b[yy][xx]&0b1111_1000_0000_0000)*255 + 15) / 31
		// 		// gr := ((b[yy][xx]&0b0000_0111_1110_0000)*255 + 31) / 63
		// 		// bl := ((b[yy][xx]&0b0000_0000_0001_1111)*255 + 15) / 31
		// 		// faceMatrix.SetPixel(xx, yy, color.RGBA{uint8(re), uint8(gr), uint8(bl), 255})
		// 		faceMatrix.SetPixel(xx, yy, rgb565ToRGBA(b[yy][xx]))
		// 	}
		// }
		// faceMatrix.Display()

		buttonEnter.Connect("clicked", func() {
			for x := int16(0); x < menuWidth; x++ {
				for y := int16(0); y < menuHeight; y++ {
					menuMatrix.SetPixel(x, y, color.RGBA{uint8(rand.Int()), uint8(rand.Int()), uint8(rand.Int()), 0})
				}
			}
			menuMatrix.QueueDraw()
		})

		g, err := gotogen.New(60, nil, menuMatrix, nil, func() (faceDisplay drivers.Displayer, menuInput gotogen.MenuInput, boopSensor gotogen.BoopSensor, err error) {
			return faceMatrix, nil, nil, nil
		})
		if err != nil {
			panic(err)
		}
		err = g.Init()
		if err != nil {
			panic(err)
		}
	})
	app.Run(os.Args)
}

func mustMakeButton(label string) *gtk.Button {
	b, err := gtk.ButtonNewWithLabel(label)
	if err != nil {
		panic(err)
	}
	return b
}
