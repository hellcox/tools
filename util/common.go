package util

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
	"regexp"
	"strings"
)

// 提取字符串中的数字
func GetNum(str string) string  {
	patten := regexp.MustCompile(`\d`)
	newStr := strings.Join(patten.FindAllString(str,-1),"")
	return newStr
}


func MakeBorder(w, h float32) fyne.CanvasObject {
	rect := canvas.NewRectangle(&color.NRGBA{235, 235, 235, 0})
	rect.SetMinSize(fyne.NewSize(w, h))
	return rect
}

