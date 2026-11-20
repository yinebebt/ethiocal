package dateconverter

import "time"

func startDayOfEthiopian(year int) int {
	// newYearDay is the start of Ethiopian new year in the Gregorian calendar.
	// The century terms are the Julian-Gregorian gap, so they come from the
	// Gregorian year in which Meskerem 1 falls (year + 7), not the Ethiopian year.
	gregorianYear := year + 7
	newYearDay := (gregorianYear / 100) - (gregorianYear / 400) - 4

	// If the previous Ethiopian year is a leap year, new year occurs on the 12th.
	if (year-1)%4 == 3 {
		newYearDay++
	}
	return newYearDay
}

func isValidGregorian(year, month, date int) bool {
	if year <= 0 || month < 1 || month > 12 || date < 1 {
		return false
	}

	t := time.Date(year, time.Month(month), date, 0, 0, 0, 0, time.UTC)
	return t.Year() == year && int(t.Month()) == month && t.Day() == date
}

func isValidEthiopian(year, month, date int) bool {
	if year <= 0 || month < 1 || month > 13 || date < 1 {
		return false
	}
	if month <= 12 {
		return date <= 30
	}
	if year%4 == 3 { // Ethiopian leap year
		return date <= 6
	}
	return date <= 5
}
