package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type myTheme struct {
}

func (t *myTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
}

func (t *myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return nil
}

func (t *myTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (t *myTheme) Size(s fyne.ThemeSizeName) float32 {
	return 14
}

