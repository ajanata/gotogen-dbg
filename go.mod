module github.com/ajanata/gotogen-simulator

go 1.22

toolchain go1.23.2

replace github.com/ajanata/gotogen => ../gotogen

replace github.com/ajanata/textbuf => ../textbuf

require (
	github.com/ajanata/oled_font v1.2.0
	github.com/ajanata/textbuf v0.0.2
	tinygo.org/x/drivers v0.23.0 // indirect
)

require (
	gioui.org v0.7.1
	gioui.org/x v0.7.1
	github.com/ajanata/gotogen v0.0.0-00010101000000-000000000000
)

require (
	gioui.org/cpu v0.0.0-20210817075930-8d6a761490d2 // indirect
	gioui.org/shader v1.0.8 // indirect
	github.com/go-text/typesetting v0.1.1 // indirect
	golang.org/x/exp v0.0.0-20240707233637-46b078467d37 // indirect
	golang.org/x/exp/shiny v0.0.0-20240707233637-46b078467d37 // indirect
	golang.org/x/image v0.18.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
)
