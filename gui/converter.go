package gui

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/yinebebt/ethiocal/dateconverter"
)

// ethMonthOptions for the month select dropdown, derived from the shared ethMonths array.
var ethMonthOptions = ethMonths[:]

// selectedEthMonth returns the 1-based month number for the selected option, or -1 if unset.
func selectedEthMonth(selected string) int {
	for i, opt := range ethMonthOptions {
		if opt == selected {
			return i + 1
		}
	}
	return -1
}

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

	yearOptions := make([]string, 0, 111)
	for y := now.Year() + 10; y >= now.Year()-100; y-- {
		yearOptions = append(yearOptions, strconv.Itoa(y))
	}
	monthOptions := []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}

	cal := widget.NewCalendar(now, nil)

	yearSelect := widget.NewSelect(yearOptions, nil)
	yearSelect.SetSelected(strconv.Itoa(now.Year()))

	monthSelect := widget.NewSelect(monthOptions, nil)
	monthSelect.SetSelected(monthOptions[now.Month()-1])

	updatingFromCal := false
	cal.OnChanged = func(t time.Time) {
		gregYear = t.Year()
		gregMonth = int(t.Month())
		gregDay = t.Day()
		dateBtn.SetText(fmt.Sprintf("%04d-%02d-%02d", gregYear, gregMonth, gregDay))
		updatingFromCal = true
		yearSelect.SetSelected(strconv.Itoa(gregYear))
		monthSelect.SetSelected(monthOptions[gregMonth-1])
		updatingFromCal = false
	}

	calBox := container.NewStack(cal)

	jumpCalendar := func() {
		if updatingFromCal {
			return
		}
		y, err := strconv.Atoi(yearSelect.Selected)
		if err != nil {
			return
		}
		mIdx := -1
		for i, m := range monthOptions {
			if m == monthSelect.Selected {
				mIdx = i
				break
			}
		}
		if mIdx < 0 {
			return
		}
		t := time.Date(y, time.Month(mIdx+1), 1, 0, 0, 0, 0, time.Local)
		cal = widget.NewCalendar(t, cal.OnChanged)
		calBox.Objects = []fyne.CanvasObject{cal}
		calBox.Refresh()
		gregYear = y
		gregMonth = mIdx + 1
		gregDay = 1
		dateBtn.SetText(fmt.Sprintf("%04d-%02d-%02d", gregYear, gregMonth, gregDay))
	}

	yearSelect.OnChanged = func(_ string) { jumpCalendar() }
	monthSelect.OnChanged = func(_ string) { jumpCalendar() }

	jumpRow := container.NewGridWithColumns(2, yearSelect, monthSelect)
	calContainer := container.NewVBox(jumpRow, calBox)
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

	// Ethiopian date input fields.

	etYearEntry := widget.NewEntry()
	etYearEntry.SetPlaceHolder("Year")

	etMonthSelect := widget.NewSelect(ethMonthOptions, nil)
	etMonthSelect.PlaceHolder = "Month"

	etDayOptions := make([]string, 30)
	for i := range etDayOptions {
		etDayOptions[i] = strconv.Itoa(i + 1)
	}
	etDaySelect := widget.NewSelect(etDayOptions, nil)
	etDaySelect.PlaceHolder = "Day"

	// Adjust day options when month or year changes (Pagume has 5-6 days).
	refreshEtDays := func() {
		y, _ := strconv.Atoi(etYearEntry.Text)
		mIdx := selectedEthMonth(etMonthSelect.Selected)
		if mIdx < 1 {
			return
		}
		maxDay := ethDaysInMonth(mIdx, y)
		opts := make([]string, maxDay)
		for i := range opts {
			opts[i] = strconv.Itoa(i + 1)
		}
		prevDay := etDaySelect.Selected
		etDaySelect.Options = opts
		// Keep previous day if still valid, otherwise clamp.
		if d, err := strconv.Atoi(prevDay); err == nil && d >= 1 && d <= maxDay {
			etDaySelect.SetSelected(prevDay)
		} else {
			etDaySelect.SetSelected(strconv.Itoa(maxDay))
		}
		etDaySelect.Refresh()
	}

	etMonthSelect.OnChanged = func(_ string) { errorLabel.Hide(); refreshEtDays() }
	etDaySelect.OnChanged = func(_ string) { errorLabel.Hide() }
	etYearEntry.OnChanged = func(_ string) { errorLabel.Hide(); refreshEtDays() }

	// Set defaults from approximate current Ethiopian date.
	etDate, err := dateconverter.Ethiopian(now.Year(), int(now.Month()), now.Day())
	if err == nil {
		etYearEntry.SetText(strconv.Itoa(etDate.Year()))
		etM := int(etDate.Month())
		if etM >= 1 && etM <= 13 {
			etMonthSelect.SetSelected(ethMonthOptions[etM-1])
		}
		etDaySelect.SetSelected(strconv.Itoa(etDate.Day()))
	}

	etYearLabel := widget.NewLabel("Year:")
	etYearLabel.TextStyle = fyne.TextStyle{Bold: true}
	etMonthLabel := widget.NewLabel("Month:")
	etMonthLabel.TextStyle = fyne.TextStyle{Bold: true}
	etDayLabel := widget.NewLabel("Day:")
	etDayLabel.TextStyle = fyne.TextStyle{Bold: true}

	ethContainer := container.NewVBox(
		container.NewGridWithColumns(3,
			container.NewVBox(etYearLabel, etYearEntry),
			container.NewVBox(etMonthLabel, etMonthSelect),
			container.NewVBox(etDayLabel, etDaySelect),
		),
	)
	ethContainer.Hide()

	// Result and error display.

	resultTitle := widget.NewLabel("")
	resultTitle.Alignment = fyne.TextAlignCenter
	resultTitle.TextStyle = fyne.TextStyle{Bold: true}
	resultLabel := widget.NewRichTextFromMarkdown("")
	resultBox := container.NewCenter(container.NewVBox(resultTitle, resultLabel))
	resultCard := widget.NewCard("", "", resultBox)
	resultCard.Hide()

	// Direction selector.

	directionSelect := widget.NewSelect(
		[]string{"Gregorian to Ethiopian", "Ethiopian to Gregorian"},
		nil,
	)
	directionSelect.SetSelected("Gregorian to Ethiopian")

	directionSelect.OnChanged = func(dir string) {
		resultCard.Hide()
		errorLabel.Hide()
		switch dir {
		case "Gregorian to Ethiopian":
			ethContainer.Hide()
			gregPickerRow.Show()
		case "Ethiopian to Gregorian":
			gregPickerRow.Hide()
			calContainer.Hide()
			ethContainer.Show()
		}
	}

	showErr := func(msg string) {
		resultCard.Hide()
		errorLabel.SetText(msg)
		errorLabel.Show()
	}

	hideErr := func() {
		errorLabel.SetText("")
		errorLabel.Hide()
	}

	convertBtn := widget.NewButtonWithIcon("Convert", theme.ConfirmIcon(), func() {
		calContainer.Hide()
		hideErr()

		switch directionSelect.Selected {
		case "Gregorian to Ethiopian":
			etDate, err := dateconverter.Ethiopian(gregYear, gregMonth, gregDay)
			if err != nil {
				showErr("Error: " + err.Error())
				return
			}
			resultTitle.SetText("Ethiopian Date")
			resultLabel.ParseMarkdown(fmt.Sprintf("## %s", etDate.Format("2006-01-02")))
			resultCard.Show()

		case "Ethiopian to Gregorian":
			y, err := strconv.Atoi(etYearEntry.Text)
			if err != nil || y <= 0 {
				showErr("Please enter a valid Ethiopian year.")
				return
			}
			mIdx := selectedEthMonth(etMonthSelect.Selected)
			if mIdx < 1 {
				showErr("Please select an Ethiopian month.")
				return
			}
			d, err := strconv.Atoi(etDaySelect.Selected)
			if err != nil {
				showErr("Please select a day.")
				return
			}
			gregDate, err := dateconverter.Gregorian(y, mIdx, d)
			if err != nil {
				showErr("Error: " + err.Error())
				return
			}
			resultTitle.SetText("Gregorian Date")
			resultLabel.ParseMarkdown(fmt.Sprintf("## %s", gregDate.Format("2006-01-02")))
			resultCard.Show()
		}
	})
	convertBtn.Importance = widget.HighImportance

	swapBtn := widget.NewButton("\u21c4", func() {
		if directionSelect.Selected == "Gregorian to Ethiopian" {
			directionSelect.SetSelected("Ethiopian to Gregorian")
		} else {
			directionSelect.SetSelected("Gregorian to Ethiopian")
		}
	})

	directionRow := container.NewBorder(nil, nil, nil, swapBtn, directionSelect)

	form := container.NewVBox(
		widget.NewCard("Date Converter", "Convert between Gregorian and Ethiopian calendars",
			container.NewVBox(
				directionRow,
				layout.NewSpacer(),
				gregPickerRow,
				ethContainer,
				convertBtn,
			),
		),
		errorLabel,
		resultCard,
		layout.NewSpacer(),
		newFooter(),
	)

	return container.NewScroll(form)
}
