package gui

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/yinebebt/ethiocal/bahirehasab"
	"github.com/yinebebt/ethiocal/dateconverter"
)

// ethMonths maps month index (0-based) to Amharic month name.
// Use MonthOfTheYear-1 to look up since MonthOfTheYear is 1-based.
var ethMonths = [13]string{
	"መስከረም", "ጥቅምት", "ኅዳር", "ታኅሣሥ", "ጥር", "የካቲት",
	"መጋቢት", "ሚያዝያ", "ግንቦት", "ሰኔ", "ሐምሌ", "ነሐሴ", "ጳጉሜ",
}

func fmtDateNamed(d bahirehasab.Date) string {
	if d.MonthOfTheYear >= 1 && d.MonthOfTheYear <= 13 {
		return fmt.Sprintf("%s %d (%02d-%02d)", ethMonths[d.MonthOfTheYear-1], d.DateOfTheMonth, d.MonthOfTheYear, d.DateOfTheMonth)
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

// festEntry is one row in the festival list: a name and its formatted date.
type festEntry struct{ name, date string }

// festivalEntries flattens a Festival into the ordered rows shown in the UI.
func festivalEntries(f bahirehasab.Festival) []festEntry {
	return []festEntry{
		{"Nenewie (ነነዌ ጾም)", fmtDateNamed(f.Basic.Nenewie)},
		{"Abiy Tsome (አብይ ጾም)", fmtDateNamed(f.Fasting.Abiy)},
		{"Debre Zeit (ደብረ ዘይት)", fmtDateNamed(f.Fasting.DebreZeit)},
		{"Hosanna (ሆሳህና)", fmtDateNamed(f.Fasting.Hosanna)},
		{"Siklet (ስቅለት)", fmtDateNamed(f.Fasting.Siklet)},
		{"Tinsaye (ፋሲካ)", fmtDateNamed(f.Fasting.Tinsaye)},
		{"Rkbe Kahnat (ርክበ ካህናት)", fmtDateNamed(f.Fasting.RkbeKahnat)},
		{"Erget (እርገት)", fmtDateNamed(f.Fasting.Erget)},
		{"Peraklitos (ጰራቅሊጦስ)", fmtDateNamed(f.Fasting.Peraklitos)},
		{"Hawariyat (ጾመ ሐዋሪያት)", fmtDateNamed(f.Fasting.Hawariyat)},
		{"Dihnet (ጾመ ድህነት)", fmtDateNamed(f.Fasting.Dihnet)},
	}
}

func newBahirTab() fyne.CanvasObject {
	curYear := currentEthiopianYear()

	errorLabel := widget.NewLabel("")
	errorLabel.Importance = widget.DangerImportance
	errorLabel.Wrapping = fyne.TextWrapWord
	errorLabel.Hide()

	// Year info (Evangelist, New Year weekday), populated on lookup.
	evangVal := widget.NewLabel("")
	newYearVal := widget.NewLabel("")
	evangLbl := widget.NewLabel("Evangelist:")
	evangLbl.TextStyle = fyne.TextStyle{Bold: true}
	newYearLbl := widget.NewLabel("New Year:")
	newYearLbl.TextStyle = fyne.TextStyle{Bold: true}
	yearInfo := container.NewGridWithColumns(4, evangLbl, evangVal, newYearLbl, newYearVal)
	infoCard := widget.NewCard("", "", container.NewVBox(
		accentHeading("አጽዋማትና በዓላት"),
		yearInfo,
	))
	infoCard.Hide()

	// Scrollable festival list. List recycles rows, so updateItem repopulates by index.
	var entries []festEntry
	list := widget.NewList(
		func() int { return len(entries) },
		func() fyne.CanvasObject {
			name := widget.NewLabel("")
			name.TextStyle = fyne.TextStyle{Bold: true}
			name.Truncation = fyne.TextTruncateEllipsis
			date := widget.NewLabel("")
			date.Alignment = fyne.TextAlignTrailing
			// Date pinned right (always shown); name fills the rest, truncated if long.
			return container.NewBorder(nil, nil, nil, date, name)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			row := o.(*fyne.Container)
			row.Objects[0].(*widget.Label).SetText(entries[i].name)
			row.Objects[1].(*widget.Label).SetText(entries[i].date)
		},
	)

	clearResults := func() {
		infoCard.Hide()
		entries = nil
		list.Refresh()
	}

	showErr := func(msg string) {
		clearResults()
		errorLabel.SetText(msg)
		errorLabel.Show()
	}

	lookup := func(yearStr string) {
		year, err := strconv.Atoi(yearStr)
		if err != nil || year < 0 {
			showErr("Please enter a valid Ethiopian year.")
			return
		}
		festival, err := bahirehasab.BahireHasab(year)
		if err != nil {
			showErr("Error: " + err.Error())
			return
		}
		errorLabel.Hide()
		evangVal.SetText(festival.Year.Evangelist)
		newYearVal.SetText(festival.Year.DayOfTheWeek)
		infoCard.Show()
		entries = festivalEntries(festival)
		list.Refresh()
	}

	yearEntry := widget.NewEntry()
	yearEntry.SetText(strconv.Itoa(curYear))
	yearEntry.OnChanged = func(s string) {
		if s == "" {
			clearResults()
			errorLabel.Hide()
			return
		}
		lookup(s)
	}

	step := func(delta int) func() {
		return func() {
			if y, err := strconv.Atoi(yearEntry.Text); err == nil {
				yearEntry.SetText(strconv.Itoa(y + delta))
			}
		}
	}
	yearLabel := widget.NewLabel("Ethiopian Year:")
	yearLabel.TextStyle = fyne.TextStyle{Bold: true}
	stepper := container.NewHBox(widget.NewButton("-", step(-1)), widget.NewButton("+", step(1)))
	yearRow := container.NewBorder(nil, nil, yearLabel, stepper, yearEntry)
	inputCard := widget.NewCard("", "", container.NewVBox(
		accentHeading("Bahire-Hasab"),
		widget.NewLabel("Ethiopian religious calendar lookup"),
		yearRow,
	))

	// Fixed header and footer; the scrolling list fills the middle (mobile-friendly).
	header := container.NewVBox(inputCard, errorLabel, infoCard)
	content := container.NewBorder(header, newFooter(), nil, nil, list)

	lookup(strconv.Itoa(curYear))

	return centered(content)
}
