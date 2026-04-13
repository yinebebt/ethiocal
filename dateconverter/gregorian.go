package dateconverter

import (
	"errors"
	"fmt"
	"time"
)

// Gregorian converts an Ethiopian date to a Gregorian date, returned as a time.Time.
func Gregorian(year, month, date int) (time.Time, error) {
	var gregorianDate int

	if !isValidEthiopian(year, month, date) {
		return time.Time{}, errors.New("not a valid date")
	}
	// Ethiopian new year in Gregorian calendar
	newYearDay := startDayOfEthiopian(year)

	// September (Ethiopian) sees 7y difference
	gregorianYear := year + 7

	// Number of days in gregorian months starting with September (index 1)
	// Index 0 is reserved for leap years switches.
	gregorianMonths := []int{0, 30, 31, 30, 31, 31, 28,
		31, 30, 31, 30, 31, 31, 30}

	//if next gregorian year is leap year, February has 29 days.
	nextYear := gregorianYear + 1
	if (nextYear%4 == 0 && nextYear%100 != 0) || nextYear%400 == 0 {
		gregorianMonths[6] = 29
	}

	// calculate number of days up to that date
	until := ((month - 1) * 30) + date
	if until <= 37 && year <= 1575 { //mysterious rule
		until += 28
		gregorianMonths[0] = 31
	} else {
		until += newYearDay - 1
	}

	// if ethiopian year is leap year, paguemain has six days
	if (year-1)%4 == 3 {
		until++
	}

	//calculate month and date incrementally
	m := 0
	for i := range gregorianMonths {
		if until <= gregorianMonths[i] {
			m = i
			gregorianDate = until
			break
		} else {
			m = i
			until -= gregorianMonths[i]
		}
	}
	//  if m > 4, we're already on next Gregorian year
	if m > 4 {
		gregorianYear++
	}
	//  Gregorian months ordered according to Ethiopian
	order := []int{8, 9, 10, 11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	gregorianMonth := order[m]

	dateResult := fmt.Sprintf("%04d-%02d-%02d", gregorianYear, gregorianMonth, gregorianDate)
	res, err := time.Parse("2006-01-02", dateResult)
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to parse converted date %q: %w", dateResult, err)
	}
	return res, nil
}
