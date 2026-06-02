package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// accent is a warm gold for primary actions, links, focus and
// selection; reads on both light and dark backgrounds.
var (
	accent          = color.NRGBA{R: 0xc8, G: 0x88, B: 0x1a, A: 0xff}
	accentSelection = color.NRGBA{R: 0xc8, G: 0x88, B: 0x1a, A: 0x33}
	accentFocus     = color.NRGBA{R: 0xc8, G: 0x88, B: 0x1a, A: 0x88}
)

// ethioTheme renders Amharic via a bundled Ethiopic font and applies the accent color.
type ethioTheme struct{}

func (e ethioTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNamePrimary, theme.ColorNameHyperlink:
		return accent
	case theme.ColorNameSelection:
		return accentSelection
	case theme.ColorNameFocus:
		return accentFocus
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (e ethioTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (e ethioTheme) Font(style fyne.TextStyle) fyne.Resource {
	// Fall back to default for monospace since the bundled font is proportional.
	if style.Monospace {
		return theme.DefaultTheme().Font(style)
	}
	// NotoSansEthiopic is a variable font covering regular and bold weights.
	return resourceNotoSansEthiopicTtf
}

func (e ethioTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
