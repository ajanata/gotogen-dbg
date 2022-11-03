module github.com/ajanata/gotogen-simulator

go 1.19

replace github.com/ajanata/gotogen => ../gotogen

replace github.com/ajanata/textbuf => ../textbuf

require (
	github.com/ajanata/oled_font v1.2.0
	github.com/ajanata/textbuf v0.0.2
	github.com/gotk3/gotk3 v0.6.1
	tinygo.org/x/drivers v0.23.0
)

require github.com/ajanata/gotogen v0.0.0-00010101000000-000000000000

require golang.org/x/image v0.0.0-20210628002857-a66eb6448b8d // indirect
