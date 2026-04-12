package gui

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/yinebebt/ethiocal/dateconverter"
)

func newConverterTab() fyne.CanvasObject {
	now := time.Now()

	selectedYear := now.Year()
	selectedMonth := int(now.Month())
	selectedDay := now.Day()

	// Date display button — shows current selection, toggles calendar.
	dateBtn := widget.NewButtonWithIcon(
		fmt.Sprintf("%04d-%02d-%02d", selectedYear, selectedMonth, selectedDay),
		theme.CalendarIcon(),
		nil,
	)

	// Year/month jump options.
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

	cal.OnChanged = func(t time.Time) {
		selectedYear = t.Year()
		selectedMonth = int(t.Month())
		selectedDay = t.Day()
		dateBtn.SetText(fmt.Sprintf("%04d-%02d-%02d", selectedYear, selectedMonth, selectedDay))
		yearSelect.SetSelected(strconv.Itoa(selectedYear))
		monthSelect.SetSelected(monthOptions[selectedMonth-1])
	}

	// calBox holds the calendar widget; replaced on year/month jump
	// instead of copying struct internals via pointer dereference.
	calBox := container.NewStack(cal)

	jumpCalendar := func() {
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
		selectedYear = y
		selectedMonth = mIdx + 1
		selectedDay = 1
		dateBtn.SetText(fmt.Sprintf("%04d-%02d-%02d", selectedYear, selectedMonth, selectedDay))
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

	// Result display — title + date both centered.
	resultTitle := widget.NewLabel("")
	resultTitle.Alignment = fyne.TextAlignCenter
	resultTitle.TextStyle = fyne.TextStyle{Bold: true}
	resultLabel := widget.NewRichTextFromMarkdown("")
	resultBox := container.NewCenter(container.NewVBox(resultTitle, resultLabel))
	resultCard := widget.NewCard("", "", resultBox)
	resultCard.Hide()

	// Error label with red color for visibility.
	errorText := canvas.NewText("", theme.Color(theme.ColorNameError))
	errorText.TextSize = 14
	errorText.Hide()

	directionSelect := widget.NewSelect(
		[]string{"Gregorian to Ethiopian", "Ethiopian to Gregorian"},
		nil,
	)
	directionSelect.SetSelected("Gregorian to Ethiopian")

	convertBtn := widget.NewButtonWithIcon("Convert", theme.ConfirmIcon(), func() {
		calContainer.Hide()
		errorText.Hide()
		errorText.Text = ""

		switch directionSelect.Selected {
		case "Gregorian to Ethiopian":
			etDate, err := dateconverter.Ethiopian(selectedYear, selectedMonth, selectedDay)
			if err != nil {
				resultCard.Hide()
				errorText.Text = "Error: " + err.Error()
				errorText.Show()
				errorText.Refresh()
				return
			}
			resultTitle.SetText("Ethiopian Date")
			resultLabel.ParseMarkdown(fmt.Sprintf("## %s", etDate.Format("2006-01-02")))
			resultCard.Show()
		case "Ethiopian to Gregorian":
			gregDate, err := dateconverter.Gregorian(selectedYear, selectedMonth, selectedDay)
			if err != nil {
				resultCard.Hide()
				errorText.Text = "Error: " + err.Error()
				errorText.Show()
				errorText.Refresh()
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

	// Date picker and convert button on one row.
	actionRow := container.NewGridWithColumns(2, dateBtn, convertBtn)

	form := container.NewVBox(
		widget.NewCard("Date Converter", "Convert between Gregorian and Ethiopian calendars",
			container.NewVBox(
				directionRow,
				layout.NewSpacer(),
				actionRow,
				calContainer,
			),
		),
		errorText,
		resultCard,
		layout.NewSpacer(),
		newFooter(),
	)

	return container.NewScroll(form)
}
