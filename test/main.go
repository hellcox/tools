package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")

	text1 := widget.NewButton("xxxxx", func() {

	})
	text2 := widget.NewButton("xx", func() {

	})
	text3 := widget.NewButton("xx", func() {

	})
	//text5 := widget.NewButton("xx", func() {
	//
	//})
	//content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3,text5)
	content := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 50)),
		text1, text2, text3)

	text4 := widget.NewButton("xx", func() {

	})
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())

	myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
	myWindow.ShowAndRun()
}
