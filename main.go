package main

import (
	_ "embed"
	"github.com/wailsapp/wails"
)

func basic(a string) string {
	return "Hello: " + a
}

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1024,
		Height:    600,
		Title:     "tools",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
		Resizable: true,
	})
	app.Bind(basic)
	app.Bind(time2Date)
	app.Bind(date2Time)
	app.Bind(getNowTimestamp)
	app.Run()
}
