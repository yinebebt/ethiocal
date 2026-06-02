package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yinebebt/ethiocal/bahirehasab"
)

var bahirCmd = &cobra.Command{
	Use:   "bahir [year]",
	Short: "Get Ethiopian fasting and religious festival dates for a given year",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		year, err := strconv.Atoi(args[0])
		if err != nil || year < 0 {
			fmt.Println("Please provide a valid Ethiopian year.")
			return
		}

		festival, err := bahirehasab.BahireHasab(year)
		if err != nil {
			fmt.Println("Error fetching bahire-hasab:", err)
			return
		}

		printFestivalInfo(festival)
	},
}

func printFestivalInfo(festival bahirehasab.Festival) {
	fmt.Printf("\nBahire-hasab Calendar for year %d\n", festival.Year.Year)
	fmt.Println("Year Information:")
	fmt.Printf("  Evangelist: %s (Number: %d)\n", festival.Year.Evangelist, festival.Year.EvangelistNum)
	fmt.Printf("  New Year falls on: %s\n", festival.Year.DayOfTheWeek)

	fmt.Println("\nBasic Information:")
	fmt.Printf("  Medeb: %d\n", festival.Basic.Medeb)
	fmt.Printf("  Wenber: %d\n", festival.Basic.Wenber)
	fmt.Printf("  Abektie: %d\n", festival.Basic.Abektie)
	fmt.Printf("  Metiq: %d\n", festival.Basic.Metiq)
	fmt.Printf("  Beale Metiq: %d\n", festival.Basic.BealeMetiq)
	fmt.Printf("  Mebaja Hamer: %d\n", festival.Basic.MebajaHamer)
	fmt.Printf("  Nenewie: %02d-%02d\n", festival.Basic.Nenewie.MonthOfTheYear, festival.Basic.Nenewie.DateOfTheMonth)

	fmt.Println("\nFasting Dates:")
	fmt.Printf("  Abiy Tsome: %02d-%02d\n", festival.Fasting.Abiy.MonthOfTheYear, festival.Fasting.Abiy.DateOfTheMonth)
	fmt.Printf("  Debre Zeit: %02d-%02d\n", festival.Fasting.DebreZeit.MonthOfTheYear, festival.Fasting.DebreZeit.DateOfTheMonth)
	fmt.Printf("  Hosanna: %02d-%02d\n", festival.Fasting.Hosanna.MonthOfTheYear, festival.Fasting.Hosanna.DateOfTheMonth)
	fmt.Printf("  Siklet: %02d-%02d\n", festival.Fasting.Siklet.MonthOfTheYear, festival.Fasting.Siklet.DateOfTheMonth)
	fmt.Printf("  Tinsaye: %02d-%02d\n", festival.Fasting.Tinsaye.MonthOfTheYear, festival.Fasting.Tinsaye.DateOfTheMonth)
	fmt.Printf("  Rkbe Kahnat: %02d-%02d\n", festival.Fasting.RkbeKahnat.MonthOfTheYear, festival.Fasting.RkbeKahnat.DateOfTheMonth)
	fmt.Printf("  Dihnet: %02d-%02d\n", festival.Fasting.Dihnet.MonthOfTheYear, festival.Fasting.Dihnet.DateOfTheMonth)
	fmt.Printf("  Hawariyat: %02d-%02d\n", festival.Fasting.Hawariyat.MonthOfTheYear, festival.Fasting.Hawariyat.DateOfTheMonth)
	fmt.Printf("  Erget: %02d-%02d\n", festival.Fasting.Erget.MonthOfTheYear, festival.Fasting.Erget.DateOfTheMonth)
	fmt.Printf("  Peraklitos: %02d-%02d\n", festival.Fasting.Peraklitos.MonthOfTheYear, festival.Fasting.Peraklitos.DateOfTheMonth)
}
