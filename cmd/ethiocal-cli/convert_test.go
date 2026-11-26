package main

import "testing"

func TestParseDate(t *testing.T) {
	tests := []struct {
		in                  string
		wantY, wantM, wantD int
		wantErr             bool
	}{
		{"2025-2-2", 2025, 2, 2, false},
		{"2025-02-02", 2025, 2, 2, false},
		{"2017-5-25", 2017, 5, 25, false},
		{"2025/2/2", 0, 0, 0, true},
		{"2025-2", 0, 0, 0, true},
		{"2025-a-2", 0, 0, 0, true},
		{"", 0, 0, 0, true},
	}
	for _, tt := range tests {
		y, m, d, err := parseDate(tt.in)
		if tt.wantErr {
			if err == nil {
				t.Errorf("parseDate(%q): want error", tt.in)
			}
			continue
		}
		if err != nil {
			t.Errorf("parseDate(%q): %v", tt.in, err)
			continue
		}
		if y != tt.wantY || m != tt.wantM || d != tt.wantD {
			t.Errorf("parseDate(%q) = %d-%d-%d, want %d-%d-%d", tt.in, y, m, d, tt.wantY, tt.wantM, tt.wantD)
		}
	}
}

func TestGregorianLong(t *testing.T) {
	tests := []struct {
		y, m, d int
		want    string
	}{
		// Julian dates: century years are leap years.
		{1500, 2, 29, "Saturday, February 29, 1500"},
		{1500, 3, 1, "Sunday, March 1, 1500"},
		{1581, 1, 1, "Sunday, January 1, 1581"},
		// The last Julian day and the first Gregorian day.
		{1582, 10, 4, "Thursday, October 4, 1582"},
		{1582, 10, 15, "Friday, October 15, 1582"},
		{2025, 2, 1, "Saturday, February 1, 2025"},
	}
	for _, tt := range tests {
		if got := gregorianLong(tt.y, tt.m, tt.d); got != tt.want {
			t.Errorf("gregorianLong(%d, %d, %d) = %q, want %q", tt.y, tt.m, tt.d, got, tt.want)
		}
	}
}
