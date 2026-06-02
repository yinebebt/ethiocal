package gui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/yinebebt/ethiocal/dateconverter"
)

// ethWeekdayHeaders are the single-letter Amharic weekday labels, Sunday-first
// to match Gregorian time.Weekday (Sunday == 0).
var ethWeekdayHeaders = [7]string{"እ", "ሰ", "ማ", "ረ", "ሐ", "ዓ", "ቅ"}

// ethCalendar is a tappable Ethiopian month-grid date picker mirroring the
// Gregorian widget.NewCalendar: chevrons navigate month/year, tapping picks a day.
type ethCalendar struct {
	year, month, day int
	onChanged        func(year, month, day int)

	title *widget.Label
	grid  *fyne.Container // 7-column grid: weekday headers + days
}

// newEthCalendar builds a picker initialised to the given Ethiopian date.
// onChanged fires with the full date whenever the user taps a day.
func newEthCalendar(year, month, day int, onChanged func(year, month, day int)) *ethCalendar {
	c := &ethCalendar{year: year, month: month, day: day, onChanged: onChanged}

	c.title = widget.NewLabel("")
	c.title.Alignment = fyne.TextAlignCenter
	c.title.TextStyle = fyne.TextStyle{Bold: true}

	c.grid = container.NewGridWithColumns(7)
	c.rebuild()
	return c
}

// object returns the renderable picker: navigation chevrons above the day grid.
func (c *ethCalendar) object() fyne.CanvasObject {
	prevYear := widget.NewButtonWithIcon("", theme.MediaSkipPreviousIcon(), func() { c.shiftYear(-1) })
	prevMonth := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() { c.shiftMonth(-1) })
	nextMonth := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() { c.shiftMonth(1) })
	nextYear := widget.NewButtonWithIcon("", theme.MediaSkipNextIcon(), func() { c.shiftYear(1) })
	for _, b := range []*widget.Button{prevYear, prevMonth, nextMonth, nextYear} {
		b.Importance = widget.LowImportance
	}

	header := container.NewBorder(nil, nil,
		container.NewHBox(prevYear, prevMonth),
		container.NewHBox(nextMonth, nextYear),
		c.title,
	)

	return container.NewVBox(header, c.grid)
}

// firstWeekday returns the weekday column (Sunday == 0) of day 1, used to
// left-pad the grid so dates line up under their weekday.
func (c *ethCalendar) firstWeekday() int {
	g, err := dateconverter.Gregorian(c.year, c.month, 1)
	if err != nil {
		return 0
	}
	return int(g.Weekday())
}

// rebuild regenerates the grid for the current year/month and highlights the
// selected day; cheap enough to call on every navigation or selection.
func (c *ethCalendar) rebuild() {
	monthName := "?"
	if c.month >= 1 && c.month <= 13 {
		monthName = ethMonths[c.month-1]
	}
	c.title.SetText(fmt.Sprintf("%s %d", monthName, c.year))

	objs := make([]fyne.CanvasObject, 0, 49)
	for _, h := range ethWeekdayHeaders {
		lbl := widget.NewLabel(h)
		lbl.Alignment = fyne.TextAlignCenter
		lbl.TextStyle = fyne.TextStyle{Bold: true}
		objs = append(objs, lbl)
	}

	for range c.firstWeekday() {
		objs = append(objs, widget.NewLabel(""))
	}

	maxDay := ethDaysInMonth(c.month, c.year)
	if c.day > maxDay {
		c.day = maxDay
	}
	for d := 1; d <= maxDay; d++ {
		btn := widget.NewButton(strconv.Itoa(d), func() {
			c.day = d
			c.rebuild()
			if c.onChanged != nil {
				c.onChanged(c.year, c.month, c.day)
			}
		})
		if d == c.day {
			btn.Importance = widget.HighImportance
		}
		objs = append(objs, btn)
	}

	c.grid.Objects = objs
	c.grid.Refresh()
}

// shiftMonth moves by delta months, wrapping across the 13-month year.
func (c *ethCalendar) shiftMonth(delta int) {
	c.month += delta
	switch {
	case c.month > 13:
		c.month = 1
		c.year++
	case c.month < 1:
		c.month = 13
		c.year--
	}
	c.rebuild()
}

// shiftYear moves by delta years, keeping the selected month and day.
func (c *ethCalendar) shiftYear(delta int) {
	c.year += delta
	c.rebuild()
}
