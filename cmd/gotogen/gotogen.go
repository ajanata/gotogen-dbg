package main

import (
	"image/color"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/outlay"
	"github.com/ajanata/gotogen"
	font "github.com/ajanata/oled_font"

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

var faceMatrix *simulator.Display
var menuMatrix *simulator.Display

var driver *simulator.Gotogen

var (
	btnBack widget.Clickable
	btnMenu widget.Clickable
	btnUp   widget.Clickable
	btnDown widget.Clickable

	slideX    = widget.Float{Value: .5}
	slideY    = widget.Float{Value: .5}
	slideZ    = widget.Float{Value: .5}
	slideBoop widget.Float

	chkTalking widget.Bool
)

func main() {
	font.PixelOff = color.RGBA{A: 255}
	font.PixelOn = color.RGBA{R: 255, G: 255, B: 255, A: 255}

	faceMatrix = simulator.NewDisplay(faceWidth, faceHeight, faceThickness)
	menuMatrix = simulator.NewDisplay(menuWidth, menuHeight, menuThickness)

	driver = &simulator.Gotogen{
		Face: faceMatrix,
		Menu: menuMatrix,
	}

	driver.UpdateAccelerometerX(float64(2 * slideX.Value))
	driver.UpdateAccelerometerY(float64(2 * slideY.Value))
	driver.UpdateAccelerometerZ(float64(2 * slideZ.Value))

	go func() {
		window := new(app.Window)
		window.Option(app.Title("Gotogen Debug - Face"))
		window.Option(app.MinSize(faceWidth*(faceThickness+1), faceHeight*(faceThickness+1)))
		window.Option(app.MaxSize(faceWidth*(faceThickness+1), faceHeight*(faceThickness+1)))
		err := runFace(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	go func() {
		window := new(app.Window)
		window.Option(app.Title("Gotogen Debug - Menu"))
		window.Option(app.MinSize(menuWidth*(menuThickness+1), menuHeight*(menuThickness+1)))
		window.Option(app.MaxSize(menuWidth*(menuThickness+1), menuHeight*(menuThickness+1)))
		err := runMenu(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	go func() {
		window := new(app.Window)
		window.Option(app.Title("Gotogen Debug - Input"))
		window.Option(app.MinSize(600, 150))
		window.Option(app.MaxSize(600, 150))
		err := runInput(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

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

	app.Main()
}

func runFace(window *app.Window) error {
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			faceMatrix.Render(gtx.Ops)

			gtx.Execute(op.InvalidateCmd{At: time.Now().Add(time.Second / 60)})
			e.Frame(gtx.Ops)
		}
	}
}

func runMenu(window *app.Window) error {
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			menuMatrix.Render(gtx.Ops)

			gtx.Execute(op.InvalidateCmd{At: time.Now().Add(time.Second / 60)})
			e.Frame(gtx.Ops)
		}
	}
}

func runInput(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// handle events
			if btnBack.Clicked(gtx) {
				driver.ButtonPress(gotogen.MenuButtonBack)
			}
			if btnMenu.Clicked(gtx) {
				driver.ButtonPress(gotogen.MenuButtonMenu)
			}
			if btnUp.Clicked(gtx) {
				driver.ButtonPress(gotogen.MenuButtonUp)
			}
			if btnDown.Clicked(gtx) {
				driver.ButtonPress(gotogen.MenuButtonDown)
			}

			if slideX.Update(gtx) {
				driver.UpdateAccelerometerX(float64(2 * slideX.Value))
			}
			if slideY.Update(gtx) {
				driver.UpdateAccelerometerY(float64(2 * slideY.Value))
			}
			if slideZ.Update(gtx) {
				driver.UpdateAccelerometerZ(float64(2 * slideZ.Value))
			}
			if slideBoop.Update(gtx) {
				driver.UpdateBoop(float64(slideBoop.Value))
			}

			if chkTalking.Update(gtx) {
				driver.UpdateTalk(chkTalking.Value)
			}

			// ////////// draw it
			dim := func(axis layout.Axis, index, constraint int) int {
				switch axis {
				case layout.Vertical:
					return 50
				case layout.Horizontal:
					return 200
				default:
					return 0
				}
			}

			grid := &outlay.Grid{
				Horizontal: outlay.AxisPosition{
					First:     0,
					Last:      2,
					Offset:    0,
					OffsetAbs: 0,
					Length:    600,
				},
				Vertical: outlay.AxisPosition{
					First:     0,
					Last:      2,
					Offset:    0,
					OffsetAbs: 0,
					Length:    150,
				},
			}
			grid.Layout(gtx, 3, 3, dim, func(gtx layout.Context, row, col int) layout.Dimensions {
				if row == 0 && col == 0 {
					return material.Button(theme, &btnUp, "Up").Layout(gtx)
				} else if row == 0 && col == 1 {
					return material.Button(theme, &btnMenu, "Menu").Layout(gtx)
				} else if row == 0 && col == 2 {
					return accelSlider(gtx, theme, &slideBoop, "Boop")
				} else if row == 1 && col == 0 {
					return material.Button(theme, &btnDown, "Down").Layout(gtx)
				} else if row == 1 && col == 1 {
					return material.Button(theme, &btnBack, "Back").Layout(gtx)
				} else if row == 1 && col == 2 {
					return material.CheckBox(theme, &chkTalking, "Talking").Layout(gtx)
				} else if row == 2 && col == 0 {
					return accelSlider(gtx, theme, &slideX, "X")
				} else if row == 2 && col == 1 {
					return accelSlider(gtx, theme, &slideY, "Y")
				} else if row == 2 && col == 2 {
					return accelSlider(gtx, theme, &slideZ, "Z")
				}
				return layout.Dimensions{}
			})

			e.Frame(gtx.Ops)
		}
	}
}

func accelSlider(gtx layout.Context, theme *material.Theme, slider *widget.Float, label string) layout.Dimensions {
	return layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceStart,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, 16, label).Layout(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.Slider(theme, slider).Layout(gtx)
		}),
	)
}
