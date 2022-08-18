package main

import (
	_ "embed"
	"github.com/wailsapp/wails"
	"strconv"
	"time"
)

func basic(a string) string {
	return "Hello World!" + a
}

func time2Date(times string) string {
	timestamp, err := strconv.ParseInt(times, 10, 64)
	if err != nil {
		return "秒级时间戳错误"
	}
	now := time.Unix(timestamp, 0)
	res := now.Format("2006-01-02 13:04:05")
	return res
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
	app.Run()
}
