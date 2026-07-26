package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type fitnessTheme struct{}

var _ fyne.Theme = (*fitnessTheme)(nil)

func (t *fitnessTheme) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch c {
	case theme.ColorNamePrimary, theme.ColorNameHyperlink, theme.ColorNameButton:
		return color.RGBA{0xE8, 0x4D, 0x0E, 0xFF} // energetic orange
	case theme.ColorNameFocus, theme.ColorNameSelection:
		return color.RGBA{0xFF, 0x7B, 0x3F, 0xFF}
	case theme.ColorNameHeaderBackground:
		return color.RGBA{0xE8, 0x4D, 0x0E, 0xFF}
	case theme.ColorNameBackground:
		if v == theme.VariantDark {
			return color.RGBA{0x1A, 0x1A, 0x2E, 0xFF}
		}
		return color.RGBA{0xF8, 0xF9, 0xFA, 0xFF}
	case theme.ColorNameError:
		return color.RGBA{0xE7, 0x4C, 0x3C, 0xFF}
	default:
		return theme.DefaultTheme().Color(c, v)
	}
}

func (t *fitnessTheme) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(s)
}

func (t *fitnessTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (t *fitnessTheme) Size(s fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(s)
}
