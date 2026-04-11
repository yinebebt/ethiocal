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

func isValid(year, month, date int) bool {
	if year <= 0 || month <= 0 || date <= 0 {
		return false
	}
	if date > 31 || month > 13 {
		return false
	}
	return true
}
