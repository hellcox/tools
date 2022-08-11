package jsontool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"tools/util"
)

func makeButtonList(count int) []fyne.CanvasObject {
	var items []fyne.CanvasObject
	for i := 1; i <= count; i++ {
		index := i // capture
		items = append(items, widget.NewButton("Button "+strconv.Itoa(index), func() {
			fmt.Println("Tapped", index)
		}))
	}

	return items
}

func makeScrollBothTab() fyne.CanvasObject {
	logo := canvas.NewImageFromResource(theme.FyneLogo())
	logo.SetMinSize(fyne.NewSize(100, 100))

	scroll := container.NewScroll(logo)
	scroll.Resize(fyne.NewSize(50, 50))

	return scroll
}


func MakeLayout(_ fyne.Window) fyne.CanvasObject {
	inVal ,outVal := `{"test":"1"	}`,""
	in :=widget.NewMultiLineEntry()
	in.Wrapping = fyne.TextWrapBreak
	in.Bind(binding.BindString(&inVal))
	out :=widget.NewMultiLineEntry()
	out.Bind(binding.BindString(&outVal))
	out.Wrapping = fyne.TextWrapBreak

	b1 := widget.NewButton("格式化", func() {
		var output bytes.Buffer
		err := json.Indent(&output, []byte(in.Text), "", "\t")
		if err!=nil{
			out.SetText(err.Error())
			return
		}
		out.SetText(output.String())
	})
	b3 := widget.NewButton("压 缩", func() {
		var data interface{}
		_ = json.Unmarshal([]byte(in.Text),&data)
		fmt.Println(data)
		str ,_ := json.Marshal(data)
		out.SetText(string(str))
	})
	b2 := widget.NewButton("转 义", func() {

	})


	horiz := container.NewHScroll(container.NewBorder(nil,nil,util.MakeBorder(2,2),nil,container.NewHBox(b1,b2,b3)))
	inWidget := container.NewVScroll(container.NewHScroll(in))
	outwidget := container.NewVScroll(container.NewHScroll(out))

	return container.NewBorder(horiz,nil,util.MakeBorder(2,2),nil,container.NewAdaptiveGrid(2,inWidget,outwidget))
}
