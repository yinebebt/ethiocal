package dateconverter

func startDayOfEthiopian(year int) int {
	// Magic formula gives start of Ethiopian new year in the Gregorian calendar.
	newYearDay := (year / 100) - (year / 400) - 4

	// If the previous Ethiopian year is a leap year, new year occurs on the 12th.
	if (year-1)%4 == 3 {
		newYearDay++
	}
	return newYearDay
}

func isValidGregorian(year, month, date int) bool {
	if year <= 0 || month <= 0 || date <= 0 {
		return false
	}
	if date > 31 || month > 12 {
		return false
	}
	return true
}

func isValidEthiopian(year, month, date int) bool {
	if year <= 0 || month <= 0 || date <= 0 {
		return false
	}
	if month > 13 {
		return false
	}
	// Pagume has 5 days, or 6 in a leap year (every 4 years). A rare 7th
	// day ("rena mealt") may occur once every 600 years; not covered here.
	if month == 13 && date > 6 {
		return false
	}
	if date > 30 {
		return false
	}
	return true
}
