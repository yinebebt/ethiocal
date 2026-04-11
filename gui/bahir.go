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

	"github.com/yinebebt/ethiocal/bahirehasab"
	"github.com/yinebebt/ethiocal/dateconverter"
)

// Ethiopian month names indexed by month number (1-13).
var ethMonths = [14]string{
	"", "Meskerem", "Tikimt", "Hidar", "Tahsas", "Tir", "Yekatit",
	"Megabit", "Miyazya", "Ginbot", "Sene", "Hamle", "Nehase", "Pagume",
}

func fmtDateNamed(d bahirehasab.Date) string {
	if d.MonthOfTheYear >= 1 && d.MonthOfTheYear <= 13 {
		return fmt.Sprintf("%s %d", ethMonths[d.MonthOfTheYear], d.DateOfTheMonth)
	}
	return fmt.Sprintf("Month %d, Day %d", d.MonthOfTheYear, d.DateOfTheMonth)
}

func currentEthiopianYear() int {
	now := time.Now()
	etDate, err := dateconverter.Ethiopian(now.Year(), int(now.Month()), now.Day())
	if err != nil {
		return 2017 // safe fallback; 2017 E.C. ≈ 2024/2025 G.C.
	}
	return etDate.Year()
}

func festivalRow(name, date string, odd bool) fyne.CanvasObject {
	nameLabel := widget.NewLabel(name)
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}
	dateLabel := widget.NewLabel(date)

	row := container.New(layout.NewGridLayout(2), nameLabel, dateLabel)

	if odd {
		// Use a subtle background for alternating rows.
		bg := canvas.NewRectangle(theme.Color(theme.ColorNameHover))
		return container.NewStack(bg, row)
	}
	return row
}

func newBahirTab() fyne.CanvasObject {
	curYear := currentEthiopianYear()

	// Error label with red color.
	errorText := canvas.NewText("", theme.Color(theme.ColorNameError))
	errorText.TextSize = 14
	errorText.Hide()

	// Combined festival card — year info + festival rows in one card.
	festContent := container.NewVBox()
	festCard := widget.NewCard("Bahire-Hasab (አጽዋማትና በዓላት)", "", festContent)
	festCard.Hide()

	// Lookup function — reused by entry change and initial load.
	lookup := func(yearStr string) {
		year, err := strconv.Atoi(yearStr)
		if err != nil || year < 0 {
			festCard.Hide()
			errorText.Text = "Please enter a valid Ethiopian year."
			errorText.Show()
			errorText.Refresh()
			return
		}

		festival, err := bahirehasab.BahireHasab(year)
		if err != nil {
			festCard.Hide()
			errorText.Text = "Error: " + err.Error()
			errorText.Show()
			errorText.Refresh()
			return
		}
		errorText.Text = ""
		errorText.Hide()

		// Year info row.
		evangLbl := widget.NewLabel("Evangelist:")
		evangLbl.TextStyle = fyne.TextStyle{Bold: true}
		evangVal := widget.NewLabel(festival.Year.Evangelist)
		newYearLbl := widget.NewLabel("New Year:")
		newYearLbl.TextStyle = fyne.TextStyle{Bold: true}
		newYearVal := widget.NewLabel(festival.Year.DayOfTheWeek)
		yearInfo := container.NewGridWithColumns(4, evangLbl, evangVal, newYearLbl, newYearVal)

		// Festival header.
		headerName := widget.NewLabel("FESTIVAL")
		headerName.TextStyle = fyne.TextStyle{Bold: true}
		headerDate := widget.NewLabel("DATE")
		headerDate.TextStyle = fyne.TextStyle{Bold: true}
		headerRow := container.New(layout.NewGridLayout(2), headerName, headerDate)

		festContent.Objects = []fyne.CanvasObject{
			yearInfo,
			widget.NewSeparator(),
			headerRow,
			widget.NewSeparator(),
			festivalRow("Nenewie (ነነዌ ጾም)", fmtDateNamed(festival.Basic.Nenewie), true),
			festivalRow("Abiy Tsome (አብይ ጾም)", fmtDateNamed(festival.Fasting.Abiy), false),
			widget.NewSeparator(),
			festivalRow("Debre Zeit (ደብረ ዘይት)", fmtDateNamed(festival.Fasting.DebreZeit), true),
			festivalRow("Hosanna (ሆሳህና)", fmtDateNamed(festival.Fasting.Hosanna), false),
			widget.NewSeparator(),
			festivalRow("Siklet (ስቅለት)", fmtDateNamed(festival.Fasting.Siklet), true),
			festivalRow("Tinsaye (ፋሲካ)", fmtDateNamed(festival.Fasting.Tinsaye), false),
			widget.NewSeparator(),
			festivalRow("Rkbe Kahnat (ርክበ ካህናት)", fmtDateNamed(festival.Fasting.RkbeKahnat), true),
			festivalRow("Erget (እርገት)", fmtDateNamed(festival.Fasting.Erget), false),
			widget.NewSeparator(),
			festivalRow("Peraklitos (ጰራቅሊጦስ)", fmtDateNamed(festival.Fasting.Peraklitos), true),
			festivalRow("Hawariyat (ጾመ ሐዋሪያት)", fmtDateNamed(festival.Fasting.Hawariyat), false),
			widget.NewSeparator(),
			festivalRow("Dihnet (ጾመ ድህነት)", fmtDateNamed(festival.Fasting.Dihnet), true),
		}
		festContent.Refresh()
		festCard.Show()
	}

	yearEntry := widget.NewEntry()
	yearEntry.SetText(strconv.Itoa(curYear))
	yearEntry.OnChanged = func(s string) {
		if s != "" {
			lookup(s)
		}
	}

	decrementBtn := widget.NewButton("-", func() {
		y, err := strconv.Atoi(yearEntry.Text)
		if err != nil {
			return
		}
		yearEntry.SetText(strconv.Itoa(y - 1))
	})
	incrementBtn := widget.NewButton("+", func() {
		y, err := strconv.Atoi(yearEntry.Text)
		if err != nil {
			return
		}
		yearEntry.SetText(strconv.Itoa(y + 1))
	})

	yearLabel := widget.NewLabel("Ethiopian Year:")
	yearLabel.TextStyle = fyne.TextStyle{Bold: true}
	yearRow := container.NewBorder(nil, nil, yearLabel, container.NewHBox(decrementBtn, incrementBtn), yearEntry)

	inputCard := widget.NewCard("Bahire-Hasab", "Ethiopian religious calendar lookup", yearRow)

	content := container.NewVBox(
		inputCard,
		errorText,
		layout.NewSpacer(),
		festCard,
		layout.NewSpacer(),
		newFooter(),
	)

	// Auto-load current year on startup.
	lookup(strconv.Itoa(curYear))

	return container.NewScroll(content)
}
