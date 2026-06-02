package gui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/yinebebt/ethiocal/dateconverter"
)

// ethDaysInMonth returns the number of days in the given Ethiopian month.
// Months 1-12 have 30 days; Pagume (13) has 5, or 6 in a leap year.
func ethDaysInMonth(month, year int) int {
	if month == 13 {
		if year%4 == 3 {
			return 6
		}
		return 5
	}
	return 30
}

func newConverterTab() fyne.CanvasObject {
	now := time.Now()

	// Gregorian date state.
	gregYear := now.Year()
	gregMonth := int(now.Month())
	gregDay := now.Day()

	dateBtn := widget.NewButtonWithIcon(
		fmt.Sprintf("%04d-%02d-%02d", gregYear, gregMonth, gregDay),
		theme.CalendarIcon(),
		nil,
	)

	cal := widget.NewCalendar(now, nil)
	cal.OnChanged = func(t time.Time) {
		gregYear = t.Year()
		gregMonth = int(t.Month())
		gregDay = t.Day()
		dateBtn.SetText(fmt.Sprintf("%04d-%02d-%02d", gregYear, gregMonth, gregDay))
	}

	calBox := container.NewStack(cal)

	// Year quick-jump: the calendar's arrows step months; these step years.
	shiftYear := func(delta int) {
		gregYear += delta
		gregDay = 1
		cal = widget.NewCalendar(
			time.Date(gregYear, time.Month(gregMonth), 1, 0, 0, 0, 0, time.Local),
			cal.OnChanged,
		)
		calBox.Objects = []fyne.CanvasObject{cal}
		calBox.Refresh()
		dateBtn.SetText(fmt.Sprintf("%04d-%02d-%02d", gregYear, gregMonth, gregDay))
	}

	prevYearBtn := widget.NewButtonWithIcon("", theme.MediaSkipPreviousIcon(), func() { shiftYear(-1) })
	nextYearBtn := widget.NewButtonWithIcon("", theme.MediaSkipNextIcon(), func() { shiftYear(1) })
	prevYearBtn.Importance = widget.LowImportance
	nextYearBtn.Importance = widget.LowImportance
	yearNav := container.NewBorder(nil, nil, prevYearBtn, nextYearBtn,
		container.NewCenter(widget.NewLabel("Year")))

	calContainer := container.NewVBox(yearNav, calBox)
	calContainer.Hide()

	dateBtn.OnTapped = func() {
		if calContainer.Visible() {
			calContainer.Hide()
		} else {
			calContainer.Show()
		}
	}

	gregPickerRow := container.NewVBox(dateBtn, calContainer)

	// Error display.
	errorLabel := widget.NewLabel("")
	errorLabel.Importance = widget.DangerImportance
	errorLabel.Wrapping = fyne.TextWrapWord
	errorLabel.Hide()

	// Ethiopian date input — a tappable month grid (see ethcalendar.go).

	// Default to the current Ethiopian date; fall back to year 1 if conversion fails.
	etY, etM, etD := 1, 1, 1
	if etDate, err := dateconverter.Ethiopian(now.Year(), int(now.Month()), now.Day()); err == nil {
		etY, etM, etD = etDate.Year(), int(etDate.Month()), etDate.Day()
	}

	ethCal := newEthCalendar(etY, etM, etD, func(_, _, _ int) { errorLabel.Hide() })

	ethContainer := container.NewVBox(ethCal.object())
	ethContainer.Hide()

	// Result and error display.

	resultTitle := widget.NewLabel("")
	resultTitle.Alignment = fyne.TextAlignCenter
	resultTitle.TextStyle = fyne.TextStyle{Bold: true}
	resultValue := widget.NewRichText()
	resultBox := container.NewCenter(container.NewVBox(resultTitle, resultValue))
	resultCard := widget.NewCard("", "", resultBox)
	resultCard.Hide()

	// showResult displays the converted date as a large accent-colored value.
	showResult := func(title, date string) {
		resultTitle.SetText(title)
		resultValue.Segments = []widget.RichTextSegment{&widget.TextSegment{
			Text: date,
			Style: widget.RichTextStyle{
				Alignment: fyne.TextAlignCenter,
				ColorName: theme.ColorNamePrimary,
				SizeName:  theme.SizeNameHeadingText,
				TextStyle: fyne.TextStyle{Bold: true},
			},
		}}
		resultValue.Refresh()
		resultCard.Show()
	}

	// Direction toggle — a compact label plus swap button, not a full-width
	// dropdown for two choices.
	gregToEth := true
	dirLabel := widget.NewLabel("Gregorian → Ethiopian")
	dirLabel.TextStyle = fyne.TextStyle{Bold: true}

	applyDirection := func() {
		resultCard.Hide()
		errorLabel.Hide()
		calContainer.Hide()
		if gregToEth {
			dirLabel.SetText("Gregorian → Ethiopian")
			ethContainer.Hide()
			gregPickerRow.Show()
		} else {
			dirLabel.SetText("Ethiopian → Gregorian")
			gregPickerRow.Hide()
			ethContainer.Show()
		}
	}

	swapBtn := widget.NewButtonWithIcon("Swap", theme.ViewRefreshIcon(), func() {
		gregToEth = !gregToEth
		applyDirection()
	})
	swapBtn.Importance = widget.LowImportance

	directionRow := container.NewHBox(dirLabel, layout.NewSpacer(), swapBtn)

	showErr := func(msg string) {
		resultCard.Hide()
		errorLabel.SetText(msg)
		errorLabel.Show()
	}

	convertBtn := widget.NewButtonWithIcon("Convert", theme.ConfirmIcon(), func() {
		calContainer.Hide()
		errorLabel.Hide()

		if gregToEth {
			etDate, err := dateconverter.Ethiopian(gregYear, gregMonth, gregDay)
			if err != nil {
				showErr("Error: " + err.Error())
				return
			}
			showResult("Ethiopian Date", etDate.Format("2006-01-02"))
			return
		}

		gregDate, err := dateconverter.Gregorian(ethCal.year, ethCal.month, ethCal.day)
		if err != nil {
			showErr("Error: " + err.Error())
			return
		}
		showResult("Gregorian Date", gregDate.Format("2006-01-02"))
	})
	convertBtn.Importance = widget.HighImportance

	form := container.NewVBox(
		widget.NewCard("", "",
			container.NewVBox(
				accentHeading("Date Converter"),
				directionRow,
				widget.NewSeparator(),
				gregPickerRow,
				ethContainer,
				convertBtn,
			),
		),
		errorLabel,
		resultCard,
		newFooter(),
	)

	return centered(container.NewScroll(form))
}
