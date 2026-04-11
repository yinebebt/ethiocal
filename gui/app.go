package gui

import (
	"fmt"
	"net/url"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/yinebebt/ethiocal/dateconverter"
)

// Run starts the GUI application.
func Run() {
	a := app.New()
	a.Settings().SetTheme(&ethioTheme{})

	title := windowTitle()
	w := a.NewWindow(title)
	w.Resize(fyne.NewSize(600, 700))

	converterTab := newConverterTab()
	bahirTab := newBahirTab()

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Date Converter", theme.HistoryIcon(), converterTab),
		container.NewTabItemWithIcon("Bahire-Hasab", theme.ListIcon(), bahirTab),
	)

	w.SetContent(tabs)
	w.ShowAndRun()
}

// windowTitle returns the window title including the current Ethiopian year.
func windowTitle() string {
	now := time.Now()
	etDate, err := dateconverter.Ethiopian(now.Year(), int(now.Month()), now.Day())
	if err != nil {
		return "Ethiocal — Ethiopian Calendar"
	}
	return fmt.Sprintf("Ethiocal — Ethiopian Calendar (%d E.C.)", etDate.Year())
}

// newFooter creates a compact footer with source and author links.
func newFooter() fyne.CanvasObject {
	ghURL, _ := url.Parse("https://github.com/yinebebt/ethiocal")
	ghLink := widget.NewHyperlink("Source", ghURL)

	authorURL, _ := url.Parse("https://yinebebtariku.com")
	authorLink := widget.NewHyperlink("yinebebtariku.com", authorURL)

	sep := widget.NewLabel("·")

	return container.NewVBox(
		container.NewCenter(
			container.NewHBox(ghLink, sep, authorLink),
		),
		widget.NewSeparator(),
	)
}
