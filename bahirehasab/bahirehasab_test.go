package bahirehasab

import (
	"testing"
)

func TestGetYear(t *testing.T) {
	sampleYear := Year{
		Year:          2016,
		EvangelistNum: 0,
		Evangelist:    "ዮሐንስ(John)",
		DayOfTheWeek:  "ማክሰኞ(Tuesday)",
	}
	year := getYear(sampleYear.Year)
	if year != sampleYear {
		t.Errorf("Test failed, expected %v got %v", sampleYear, year)
	}
}

func TestBasic(t *testing.T) {
	// Basic for the year 2016
	expected := Basic{
		Medeb:       11,
		Wenber:      10,
		Abektie:     20,
		Metiq:       10,
		BealeMetiq:  2,
		MebajaHamer: 18,
		Nenewie: Date{
			DateOfTheMonth: 18,
			MonthOfTheYear: 6,
		},
	}
	got := getBasic(2016)
	if got != expected {
		t.Errorf("Test failed: expected %v got %v", expected, got)
	}
}

func TestFasting(t *testing.T) {
	// year 2016 sample fasting dates
	expected := Fasting{
		Abiy: Date{
			DateOfTheMonth: 2,
			MonthOfTheYear: 7,
		},
		DebreZeit: Date{
			DateOfTheMonth: 29,
			MonthOfTheYear: 7,
		},
		Hosanna: Date{
			DateOfTheMonth: 20,
			MonthOfTheYear: 8,
		},
		Siklet: Date{
			DateOfTheMonth: 25,
			MonthOfTheYear: 8,
		},
		Tinsaye: Date{
			DateOfTheMonth: 27,
			MonthOfTheYear: 8,
		},
		RkbeKahnat: Date{
			DateOfTheMonth: 21,
			MonthOfTheYear: 9,
		},
		Erget: Date{
			DateOfTheMonth: 6,
			MonthOfTheYear: 10,
		},
		Peraklitos: Date{
			DateOfTheMonth: 16,
			MonthOfTheYear: 10,
		},
		Hawariyat: Date{
			DateOfTheMonth: 17,
			MonthOfTheYear: 10,
		},
		Dihnet: Date{
			DateOfTheMonth: 19,
			MonthOfTheYear: 10,
		},
	}

	got := getFasting(2016)
	if expected != got {
		t.Errorf("Test failed expected %v got %v", expected, got)
	}
}

func TestBahireHasab(t *testing.T) {
	res, err := BahireHasab(2016)
	if err != nil {
		t.Errorf("Test failed, expected %v but got %v ", nil, err)
	}
	if res.Year.Year != 2016 {
		t.Errorf("Test failed, expected %v but got %v ", 2016, res.Year.Year)
	}
}
