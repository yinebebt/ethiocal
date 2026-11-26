package dateconverter

import (
	"testing"
	"time"
)

func TestConvert(t *testing.T) {
	t.Parallel()

	pairs := []struct{ gy, gm, gd, ey, em, ed int }{
		{2022, 9, 28, 2015, 1, 18},
		{2022, 9, 11, 2015, 1, 1},
		{2025, 2, 2, 2017, 5, 25},
		{1582, 10, 4, 1575, 2, 7},
		{1582, 10, 15, 1575, 2, 8},
	}
	for _, tt := range pairs {
		g := time.Date(tt.gy, time.Month(tt.gm), tt.gd, 0, 0, 0, 0, time.UTC)
		e := time.Date(tt.ey, time.Month(tt.em), tt.ed, 0, 0, 0, 0, time.UTC)
		t.Run(g.Format("2006-01-02"), func(t *testing.T) {
			got, err := Ethiopian(tt.gy, tt.gm, tt.gd)
			if err != nil {
				t.Fatal(err)
			}
			if !got.Equal(e) {
				t.Errorf("Ethiopian: got %v, want %v", got, e)
			}

			got, err = Gregorian(tt.ey, tt.em, tt.ed)
			if err != nil {
				t.Fatal(err)
			}
			if !got.Equal(g) {
				t.Errorf("Gregorian: got %v, want %v", got, g)
			}
		})
	}
}

func TestGregorian_pagume(t *testing.T) {
	t.Parallel()

	got, err := Gregorian(2015, 13, 6)
	if err != nil {
		t.Fatal(err)
	}
	if !got.Equal(time.Date(2023, 9, 11, 0, 0, 0, 0, time.UTC)) {
		t.Errorf("leap pagume 6: got %v", got)
	}

	got, err = Gregorian(2016, 13, 5)
	if err != nil {
		t.Fatal(err)
	}
	if !got.Equal(time.Date(2024, 9, 10, 0, 0, 0, 0, time.UTC)) {
		t.Errorf("non-leap pagume 5: got %v", got)
	}
}

func TestEthiopian_leapDay(t *testing.T) {
	t.Parallel()
	got, err := Ethiopian(2024, 2, 29)
	if err != nil {
		t.Fatal(err)
	}
	if !got.Equal(time.Date(2016, 6, 21, 0, 0, 0, 0, time.UTC)) {
		t.Errorf("got %v, want 2016-06-21", got)
	}
}

func TestInvalid(t *testing.T) {
	t.Parallel()
	if _, err := Ethiopian(0, 1, 1); err == nil {
		t.Error("year 0")
	}
	if _, err := Ethiopian(1582, 10, 10); err == nil {
		t.Error("reform gap")
	}
	if _, err := Gregorian(2015, 13, 7); err == nil {
		t.Error("pagume 7")
	}
	if _, err := Gregorian(2016, 13, 6); err == nil {
		t.Error("pagume 6 non-leap")
	}
}

func TestGregorian_yearsDivisibleByFour(t *testing.T) {
	t.Parallel()

	pairs := []struct{ ey, em, ed, gy, gm, gd int }{
		{2016, 1, 1, 2023, 9, 12},
		{2016, 9, 14, 2024, 5, 22},
		{2020, 1, 1, 2027, 9, 12},
		{1992, 1, 1, 1999, 9, 12},
		// century years: the Julian-Gregorian gap grows after February 1900 and 2100
		{1892, 1, 1, 1899, 9, 11},
		{1893, 1, 1, 1900, 9, 11},
		{2093, 1, 1, 2100, 9, 12},
	}
	for _, tt := range pairs {
		got, err := Gregorian(tt.ey, tt.em, tt.ed)
		if err != nil {
			t.Fatal(err)
		}
		if want := time.Date(tt.gy, time.Month(tt.gm), tt.gd, 0, 0, 0, 0, time.UTC); !got.Equal(want) {
			t.Errorf("Gregorian(%d, %d, %d): got %v, want %v", tt.ey, tt.em, tt.ed, got.Format("2006-01-02"), want.Format("2006-01-02"))
		}
	}

	got, err := Ethiopian(1901, 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	if want := time.Date(1893, 4, 23, 0, 0, 0, 0, time.UTC); !got.Equal(want) {
		t.Errorf("Ethiopian(1901, 1, 1): got %v, want %v", got.Format("2006-01-02"), want.Format("2006-01-02"))
	}
}

func TestGregorian_consecutiveDays(t *testing.T) {
	t.Parallel()

	// The last day of each Ethiopian year must be followed by Meskerem 1 of the next.
	for year := 1600; year <= 2400; year++ {
		pagume := 5
		if year%4 == 3 {
			pagume = 6
		}
		last, err := Gregorian(year, 13, pagume)
		if err != nil {
			t.Fatal(err)
		}
		first, err := Gregorian(year+1, 1, 1)
		if err != nil {
			t.Fatal(err)
		}
		if !first.Equal(last.AddDate(0, 0, 1)) {
			t.Errorf("Pagume %d, %d = %v but Meskerem 1, %d = %v", pagume, year, last.Format("2006-01-02"), year+1, first.Format("2006-01-02"))
		}
	}
}

func TestGregorian_julianYearsDivisibleByFour(t *testing.T) {
	t.Parallel()

	// Before 1582 the result is a Julian date. After an Ethiopian leap year,
	// Meskerem 1 falls on August 30 instead of August 29.
	pairs := []struct{ ey, em, ed, gy, gm, gd int }{
		{4, 1, 1, 11, 8, 30},
		{1500, 1, 1, 1507, 8, 30},
		{1572, 1, 1, 1579, 8, 30},
		{1572, 2, 7, 1579, 10, 5},
		{1571, 1, 1, 1578, 8, 29},
		{1573, 1, 1, 1580, 8, 29},
	}
	for _, tt := range pairs {
		got, err := Gregorian(tt.ey, tt.em, tt.ed)
		if err != nil {
			t.Fatal(err)
		}
		if want := time.Date(tt.gy, time.Month(tt.gm), tt.gd, 0, 0, 0, 0, time.UTC); !got.Equal(want) {
			t.Errorf("Gregorian(%d, %d, %d): got %v, want %v", tt.ey, tt.em, tt.ed, got.Format("2006-01-02"), want.Format("2006-01-02"))
		}
	}
}

func TestEthiopian_julian(t *testing.T) {
	t.Parallel()

	// Before 15 Oct 1582 the input is a Julian date. Genna (Tahsas 29) is
	// December 25, so January 1 is Tir 6, or Tir 5 after an Ethiopian leap year.
	pairs := []struct{ gy, gm, gd, ey, em, ed int }{
		{1581, 1, 1, 1573, 5, 6},
		{1580, 1, 1, 1572, 5, 5},
		{1581, 1, 25, 1573, 5, 30},
		{1581, 1, 26, 1573, 6, 1},
		{1581, 12, 25, 1574, 4, 29},
		{1581, 12, 27, 1574, 5, 1},
		{1581, 12, 31, 1574, 5, 5},
		// Julian century years are leap years.
		{1500, 1, 1, 1492, 5, 5},
		{1500, 2, 29, 1492, 7, 4},
		{1500, 3, 1, 1492, 7, 5},
		{1500, 12, 31, 1493, 5, 5},
	}
	for _, tt := range pairs {
		got, err := Ethiopian(tt.gy, tt.gm, tt.gd)
		if err != nil {
			t.Errorf("Ethiopian(%d, %d, %d): %v", tt.gy, tt.gm, tt.gd, err)
			continue
		}
		if want := time.Date(tt.ey, time.Month(tt.em), tt.ed, 0, 0, 0, 0, time.UTC); !got.Equal(want) {
			t.Errorf("Ethiopian(%d, %d, %d): got %v, want %v", tt.gy, tt.gm, tt.gd, got.Format("2006-01-02"), want.Format("2006-01-02"))
		}
	}
}
