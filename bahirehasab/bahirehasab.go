package bahirehasab

import (
	"fmt"
)

// የበዓላትና የአጽዋማት ተውሳክ
var tewsak = map[string]int{
	"abiy":       14,
	"debrezeit":  41,
	"hosahna":    62,
	"siklet":     67,
	"tinsaye":    69,
	"rkbekahnat": 93,
	"erget":      108,
	"peraklitos": 118,
	"hawariyat":  119,
	"dihnet":     121,
}

// BahireHasab computes the Ethiopian religious calendar (fasting and festival
// dates) for the given Ethiopian year.
func BahireHasab(etYear int) (Festival, error) {
	if etYear < 0 {
		return Festival{}, fmt.Errorf("invalid Ethiopian Year %v", etYear)
	}
	year := getYear(etYear)
	basic := getBasic(etYear)
	fasting := getFasting(etYear)
	return Festival{
		Year:    year,
		Basic:   basic,
		Fasting: fasting,
	}, nil
}

func getYear(year int) Year {
	var wngName string
	ameteAlem := year + 5500
	wngNum := ameteAlem % 4
	switch wngNum {
	case 0:
		wngName = "ዮሐንስ(John)"
	case 1:
		wngName = "ማቲዎስ(Matthew)"
	case 2:
		wngName = "ማርቆስ(Mark)"
	case 3:
		wngName = "ሉቃስ(Luke)"
	}
	meteneRabit := (ameteAlem - wngNum) / 4
	tinteQemer := (ameteAlem + meteneRabit) % 7

	return Year{
		Year:          year,
		EvangelistNum: wngNum,
		Evangelist:    wngName,
		DayOfTheWeek:  getDayOfTheWeek(tinteQemer),
	}
}

func getDayOfTheWeek(day int) string {
	switch day {
	case 0:
		return "ሰኞ(Monday)"
	case 1:
		return "ማክሰኞ(Tuesday)"
	case 2:
		return "እሮብ(Wednesday)"
	case 3:
		return "ሃሙስ(Thursday)"
	case 4:
		return "አርብ(Friday)"
	case 5:
		return "ቅዳሜ(Saturday)"
	case 6:
		return "እሁድ(Sunday)"
	default:
		return "invalid day"
	}
}

func getBasic(year int) Basic {
	wenber := 0
	amtAlem := year + 5500
	medeb := amtAlem % 19
	if medeb == 0 {
		wenber = 18
	} else {
		wenber = medeb - 1
	}

	abektie := (wenber * TinteAbektie) % 30
	metiq := (wenber * TinteMetiq) % 30
	wngNum := amtAlem % 4
	meteneRabit := (amtAlem - wngNum) / 4
	tinteQemer := (amtAlem + meteneRabit) % 7

	var dayOfWeekForMtqNum int
	bealeMetq := 0
	if metiq > 14 {
		bealeMetq = 1
		dayOfWeekForMtqNum = (metiq-1)%7 + tinteQemer
	} else if metiq < 14 {
		bealeMetq = 2
		dayOfWeekForMtqNum = (metiq+29)%7 + tinteQemer
	}

	twsakOfDay := 0
	switch dayOfWeekForMtqNum {
	case 0:
		twsakOfDay = monday
	case 1:
		twsakOfDay = tuesday
	case 2:
		twsakOfDay = wednesday
	case 3:
		twsakOfDay = thursday
	case 4:
		twsakOfDay = friday
	case 5:
		twsakOfDay = saturday
	case 6:
		twsakOfDay = sunday
	}

	mebajaHamer := metiq + twsakOfDay // የነነዌ ፆም የሚውልበት ቀን
	nenewie := Date{}
	if bealeMetq == 1 && mebajaHamer <= 30 {
		nenewie = Date{
			DateOfTheMonth: mebajaHamer,
			MonthOfTheYear: 5,
		}
	} else if bealeMetq == 2 || (bealeMetq == 1 && mebajaHamer > 30) {
		nenewie = Date{
			DateOfTheMonth: mebajaHamer % 30,
			MonthOfTheYear: 6,
		}
	}

	return Basic{
		Medeb:       medeb,
		Wenber:      wenber,
		Abektie:     abektie,
		Metiq:       metiq,
		BealeMetiq:  bealeMetq,
		MebajaHamer: mebajaHamer,
		Nenewie:     nenewie,
	}
}

func getFasting(etYear int) Fasting {
	basic := getBasic(etYear)
	fastingMap := map[string]Date{}
	for name, val := range tewsak {
		daySum, offset := 0, 0
		daySum = basic.Nenewie.DateOfTheMonth + val
		offset = (daySum - (daySum % 30)) / 30
		date := Date{
			DateOfTheMonth: daySum % 30,
			MonthOfTheYear: basic.Nenewie.MonthOfTheYear + offset,
		}
		fastingMap[name] = date
	}

	fasting := Fasting{}
	fasting.Abiy = fastingMap["abiy"]
	fasting.DebreZeit = fastingMap["debrezeit"]
	fasting.Hosanna = fastingMap["hosahna"]
	fasting.Siklet = fastingMap["siklet"]
	fasting.Tinsaye = fastingMap["tinsaye"]
	fasting.RkbeKahnat = fastingMap["rkbekahnat"]
	fasting.Erget = fastingMap["erget"]
	fasting.Peraklitos = fastingMap["peraklitos"]
	fasting.Hawariyat = fastingMap["hawariyat"]
	fasting.Dihnet = fastingMap["dihnet"]

	return fasting
}
