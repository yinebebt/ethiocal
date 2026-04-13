package bahirehasab

const (
	TinteMetiq   = 19
	TinteAbektie = 11
)

// የዕለታት ተውሳክ
const (
	saturday  = 8
	sunday    = 7
	monday    = 6
	tuesday   = 5
	wednesday = 4
	thursday  = 3
	friday    = 2
)

// Festival contains all relevant information on ባሐረ-ሃሳብ calendar such as Year, Basic events and Fasting dates.
type Festival struct {
	Year    Year    `json:"year"`
	Basic   Basic   `json:"basic"`
	Fasting Fasting `json:"fasting"`
}

// Basic contains bahirehasab related values such as መጥቅ including ነነዌ ፆም which is used to get all others.
type Basic struct {
	//Medeb (መደብ)
	Medeb int `json:"medeb"`
	//Wenber (ወንበር)
	Wenber int `json:"wenber"`
	// Abekti (አበቅቴ)
	Abektie int `json:"abektie"`
	// Metiq (መጥቅ)
	Metiq int `json:"metiq"`
	// BealeMetiq (በዓለ-መጥቅ), 1 =መስከረም፣ 2 = ጥቅምት
	BealeMetiq int `json:"beale_metiq"`
	// MebajaHamer (መባጃ ሃመር)
	MebajaHamer int `json:"mebaja_hamer"`
	// Nenewie የነነዌ ፆም የሚውልበት ቀን
	Nenewie Date `json:"nenewie"`
}

// Year contains New year data such as Evangelist.
type Year struct {
	// Year in Ethiopian calendar
	Year int `json:"year"`
	// EvangelistNum (Evangelist) ወንጌላዊው in number
	EvangelistNum int `json:"evangelist_num"`
	// Evangelist ወንጌላዊው in name
	Evangelist string `json:"evangelist"`
	// DayOfTheWeek the day in which new year fall
	DayOfTheWeek string `json:"day_of_the_week"`
}

// Date contains date of the month and month of the year.
// MonthOfTheYear is 1-based: መስከረም = 1 .. ጳጉሜን = 13. DateOfTheMonth goes 1 to 30.
type Date struct {
	DateOfTheMonth int `json:"date_of_the_month"`
	MonthOfTheYear int `json:"month_of_the_year"`
}

// Fasting አጽዋማትና በዓላት
type Fasting struct {
	// Abiy tsome (አብይ ፆም)
	Abiy Date `json:"abiy"`
	// DebreZeit ደብረ ዘይት
	DebreZeit Date `json:"debrezeit"`
	// Hosanna ሆሳህና
	Hosanna Date `json:"hosanna"`
	// Siklet (ስቅለት), Good Friday
	Siklet Date `json:"siklet"`
	// Tinsaye (ፋሲካ), resurrection
	Tinsaye Date `json:"tinsaye"`
	// RkbeKahnat (ርክበ ካህናት)
	RkbeKahnat Date `json:"rkbekahnat"`
	// Dihnet (ፆመ ድህነት), Wednesday and Friday Weekly fast
	Dihnet Date `json:"dihnet"`
	// Hawariyat ( ፆመ ሐዋሪያት), fast of the Holy Apostles
	Hawariyat Date `json:"hawariyat"`
	// Erget (እርገት)፣ Ascension
	Erget Date `json:"erget"`
	// Peraklitos (ጰራቅሊጦስ)፣ Paraclete
	Peraklitos Date `json:"peraklitos"`
}
