package main

import (
	"image/color"
	"os"

	font "github.com/ajanata/oled_font"
	"github.com/ajanata/textbuf"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"github.com/ajanata/gotogen"

	"github.com/ajanata/gotogen-simulator/pixbufmatrix"
	"github.com/ajanata/gotogen-simulator/simulator"
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

		// input window
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

		// buttons
		buttonBack := mustMakeButton("Back")
		buttonUp := mustMakeButton("Up")
		buttonDown := mustMakeButton("Down")
		buttonMenu := mustMakeButton("Menu")
		// buttonDefault := mustMakeButton("Default")
		inputGrid.Attach(buttonUp, 1, 0, 1, 1)
		inputGrid.Attach(buttonDown, 1, 1, 1, 1)
		// inputGrid.Attach(buttonDefault, 0, 1, 1, 1)
		inputGrid.Attach(buttonBack, 0, 0, 1, 1)
		inputGrid.Attach(buttonMenu, 0, 1, 1, 1)

		// accelerometer
		accX, err := gtk.ScaleNewWithRange(gtk.ORIENTATION_HORIZONTAL, -1, 1, .01)
		if err != nil {
			panic(err)
		}
		accX.SetValue(0)
		accX.SetTooltipText("Accelerometer X")
		accY, err := gtk.ScaleNewWithRange(gtk.ORIENTATION_HORIZONTAL, -1, 1, .01)
		if err != nil {
			panic(err)
		}
		accY.SetValue(0)
		accY.SetTooltipText("Accelerometer Y")
		accZ, err := gtk.ScaleNewWithRange(gtk.ORIENTATION_HORIZONTAL, -1, 1, .01)
		if err != nil {
			panic(err)
		}
		accZ.SetValue(0)
		accZ.SetTooltipText("Accelerometer Z")
		inputGrid.Attach(accX, 0, 2, 2, 1)
		inputGrid.Attach(accY, 0, 3, 2, 1)
		inputGrid.Attach(accZ, 0, 4, 2, 1)

		// boop sensor
		boop, err := gtk.ScaleNewWithRange(gtk.ORIENTATION_VERTICAL, 0, 1, .1)
		if err != nil {
			panic(err)
		}
		boop.SetTooltipText("Boop sensor")
		inputGrid.Attach(boop, 2, 0, 1, 2)

		// talking
		talk, err := gtk.CheckButtonNewWithLabel("Talking")
		if err != nil {
			panic(err)
		}
		inputGrid.Attach(talk, 2, 2, 1, 1)

		inputWindow.Add(inputGrid)
		inputWindow.ShowAll()

		driver := &simulator.Gotogen{
			Face: faceMatrix,
			Menu: menuMatrix,
		}

		// action handlers
		buttonBack.Connect("clicked", func() {
			driver.ButtonPress(gotogen.MenuButtonBack)
		})
		buttonUp.Connect("clicked", func() {
			driver.ButtonPress(gotogen.MenuButtonUp)
		})
		buttonDown.Connect("clicked", func() {
			driver.ButtonPress(gotogen.MenuButtonDown)
		})
		buttonMenu.Connect("clicked", func() {
			driver.ButtonPress(gotogen.MenuButtonMenu)
		})
		// buttonDefault.Connect("clicked", func() {
		// 	driver.ButtonPress(gotogen.MenuButtonDefault)
		// })
		accX.Connect("value-changed", func() {
			driver.UpdateAccelerometerX(accX.GetValue())
		})
		accY.Connect("value-changed", func() {
			driver.UpdateAccelerometerY(accY.GetValue())
		})
		accZ.Connect("value-changed", func() {
			driver.UpdateAccelerometerZ(accZ.GetValue())
		})
		boop.Connect("value-changed", func() {
			driver.UpdateBoop(boop.GetValue())
		})
		talk.Connect("toggled", func() {
			driver.UpdateTalk(talk.GetActive())
		})

		go func() {
			g, err := gotogen.New(60, menuMatrix, nil, driver)
			if err != nil {
				panic(err)
			}
			err = g.Init()
			if err != nil {
				panic(err)
			}

			g.Run()
		}()
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
