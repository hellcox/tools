package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"os"
	"tools/tool/jsontool"
	"tools/tool/timetool"
)

func init() {
	os.Setenv("FYNE_FONT", "assets/AlibabaPuHuiTi-2-65-Medium.ttf")
	os.Setenv("FYNE_SCALE", "1")
}

var MyWindow fyne.Window

func main() {
	myApp := app.New()
	icon, _ := fyne.LoadResourceFromPath("assets/icon1.png")
	myApp.SetIcon(icon)
	MyWindow := myApp.NewWindow("工具箱")
	MyWindow.CenterOnScreen()
	//myApp.Settings().SetTheme(theme.DarkTheme())

	disabled := widget.NewButton("Disabled", func() {})
	disabled.Disable()

	tabs := container.NewAppTabs(
		container.NewTabItem("时  间 ", timetool.MakeGridLayout(MyWindow)),
		container.NewTabItem("JSON ", jsontool.MakeLayout(MyWindow)),
		container.NewTabItem("布  局 ", MakeLayout(MyWindow)),
		container.NewTabItem("文  本 ", container.New(layout.NewVBoxLayout(), widget.NewLabel(" xxxxx"))),
		container.NewTabItem("关  于 ", container.New(layout.NewVBoxLayout(), widget.NewLabel(" Version:\t0.0.1 \n Date:\t2022-08-02 \n Author:\tHellcox"))),
	)
	tabs.SetTabLocation(container.TabLocationLeading)
	tabs.Size()

	MyWindow.Resize(fyne.Size{
		Width:  750,
		Height: 400,
	})
	MyWindow.SetMaster()
	MyWindow.SetContent(tabs)
	MyWindow.ShowAndRun()
	os.Unsetenv("FYNE_FONT")
}
