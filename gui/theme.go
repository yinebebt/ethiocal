package gui

// bundled.go was generated with: fyne bundle -o bundled.go NotoSansEthiopic.ttf

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// ethioTheme overrides the default Fyne theme to use a font with Ethiopic
// (Ge'ez) script support so that Amharic text renders correctly.
type ethioTheme struct{}

func (e ethioTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (e ethioTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (e ethioTheme) Font(_ fyne.TextStyle) fyne.Resource {
	return resourceNotoSansEthiopicTtf
}

func (e ethioTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
