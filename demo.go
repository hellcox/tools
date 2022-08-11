package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"tools/util"
)

func MakeLayout(_ fyne.Window) fyne.CanvasObject {
	border1 := util.MakeBorder(22, 22)
	border2 := util.MakeBorder(22, 22)
	border3 := util.MakeBorder(22, 22)
	border4 := util.MakeBorder(22, 22)

	//box1 := makeInput()
	//box2 := makeButton()
	//box3 := makeOutput()

	middle := container.NewGridWithRows(3,
		container.NewGridWithColumns(3,
			widget.NewEntry(),
			widget.NewButton("xx", func() {}),
			widget.NewButton("xx", func() {}),
		),

		widget.NewButton("dark", func() {
			a := fyne.CurrentApp()
			a.Settings().SetTheme(theme.DarkTheme())
			fmt.Println(" set dark")
		}),

		container.NewGridWrap(fyne.Size{
			Width:  100,
			Height: 0,
		}),
	)

	return container.NewBorder(border1, border2, border3, border4, middle)
	return nil
}
