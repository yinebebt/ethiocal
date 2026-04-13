package dateconverter

import (
	"errors"
	"fmt"
	"time"
)

// Ethiopian converts a Gregorian date to an Ethiopian date, returned as a time.Time.
func Ethiopian(year, month, date int) (time.Time, error) {
	var december int
	var ethiopianDate int

	if !isValidGregorian(year, month, date) {
		return time.Time{}, errors.New("not a valid date")
	}

	// Dates between 5 and 14 of October 1582 are invalid (Gregorian reform gap).
	if month == 10 && date >= 5 && date <= 14 && year == 1582 {
		return time.Time{}, errors.New("invalid date: October 5-14, 1582 do not exist in the Gregorian calendar")
	}

	// Number of days in gregorian months starting with January (index 1)
	//	Index 0 is reserved for leap years switches.
	gregorianMonths := []int{0, 31, 28, 31, 30, 31, 30,
		31, 31, 30, 31, 30, 31}

	ethiopianMonths := []int{0, 30, 30, 30, 30, 30, 30, 30,
		30, 30, 5, 30, 30, 30, 30}

	//  if gregorian leap year, February has 29 days.
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		gregorianMonths[2] = 29
	}
	//  September sees 8y difference
	ethiopianYear := year - 8

	//  if ethiopian leap year pagumain has 6 days
	if ethiopianYear%4 == 3 {
		ethiopianMonths[10] = 6
	} else {
		ethiopianMonths[10] = 5
	}
	//  Ethiopian new year in Gregorian calendar
	newYearDay := startDayOfEthiopian(year - 8)

	// calculate number of days up to that date
	until := 0
	for i := 1; i < month; i++ {
		until += gregorianMonths[i]
	}
	until += date

	// # update december (december) to match january 1st
	if ethiopianYear%4 == 0 {
		december = 26
	} else {
		december = 25
	}

	// take into account the 1582 change
	if year < 1582 {
		ethiopianMonths[1] = 0
		ethiopianMonths[2] = december
	} else if until <= 277 && year == 1582 {
		ethiopianMonths[1] = 0
		ethiopianMonths[2] = december
	} else {
		december = newYearDay - 3
		ethiopianMonths[1] = december
	}
	// calculate month and date incrementally
	m := 0
	for m = range ethiopianMonths {
		if until <= ethiopianMonths[m] {
			if m == 1 || ethiopianMonths[m] == 0 {
				ethiopianDate = until + (30 - december)
			} else {
				ethiopianDate = until
			}
			break
		} else {
			until -= ethiopianMonths[m]
		}
	}
	//  if m > 4, we're already on next Ethiopian year
	if m > 10 {
		ethiopianYear++
	}
	// Ethiopian months ordered according to Gregorian
	order := []int{0, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 1, 2, 3, 4}
	ethiopianMonth := order[m]

	dateResult := fmt.Sprintf("%04d-%02d-%02d", ethiopianYear, ethiopianMonth, ethiopianDate)
	res, err := time.Parse("2006-01-02", dateResult)
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to parse converted date %q: %w", dateResult, err)
	}
	return res, nil
}
