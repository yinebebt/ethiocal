package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/yinebebt/ethiocal/dateconverter"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert dates between Ethiopian and Gregorian calendars",
}

// parseDate parses a YYYY-MM-DD date string.
func parseDate(s string) (year, month, day int, err error) {
	parts := strings.Split(s, "-")
	if len(parts) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid date %q, expected YYYY-MM-DD", s)
	}
	fields := []struct {
		name string
		out  *int
	}{
		{"year", &year},
		{"month", &month},
		{"day", &day},
	}
	for i, f := range fields {
		if *f.out, err = strconv.Atoi(parts[i]); err != nil {
			return 0, 0, 0, fmt.Errorf("invalid %s in %q", f.name, s)
		}
	}
	return year, month, day, nil
}

var gtoeCmd = &cobra.Command{
	Use:     "gtoe [date]",
	Short:   "Convert Gregorian date to Ethiopian date",
	Long:    "Convert a Gregorian date (YYYY-MM-DD) to Ethiopian.",
	Args:    cobra.ExactArgs(1),
	Example: "ethiocal-cli convert gtoe 2025-02-1",
	Run: func(cmd *cobra.Command, args []string) {
		year, month, day, err := parseDate(args[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		etDate, err := dateconverter.Ethiopian(year, month, day)
		if err != nil {
			fmt.Println("Error converting date:", err)
			return
		}

		fmt.Printf("\nGregorian Date: %s\n", gregorianLong(year, month, day))
		fmt.Printf("  %04d-%02d-%02d\n", year, month, day)
		if long, err := dateconverter.FormatEthiopian(etDate.Year(), int(etDate.Month()), etDate.Day()); err == nil {
			fmt.Println("Converted Ethiopian Date:", long)
		}
		fmt.Println(" ", etDate.Format("2006-01-02"))
	},
}

var etogCmd = &cobra.Command{
	Use:     "etog [date]",
	Short:   "Convert Ethiopian date to Gregorian date",
	Long:    "Convert an Ethiopian date (YYYY-MM-DD) to Gregorian.",
	Args:    cobra.ExactArgs(1),
	Example: "ethiocal-cli convert etog 2017-5-25",
	Run: func(cmd *cobra.Command, args []string) {
		year, month, day, err := parseDate(args[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		gregDate, err := dateconverter.Gregorian(year, month, day)
		if err != nil {
			fmt.Println("Error converting date:", err)
			return
		}
		if long, err := dateconverter.FormatEthiopian(year, month, day); err == nil {
			fmt.Println("\nEthiopian Date:", long)
		}
		fmt.Printf("  %04d-%02d-%02d\n", year, month, day)
		fmt.Println("Converted Gregorian Date:", dateconverter.FormatGregorian(gregDate))
		fmt.Println(" ", gregDate.Format("2006-01-02"))
	},
}

func init() {
	convertCmd.AddCommand(gtoeCmd)
	convertCmd.AddCommand(etogCmd)
}

// gregorianLong formats a date given in the calendar that Ethiopian reads:
// Julian before 15 Oct 1582, Gregorian after. time.Date uses Gregorian rules
// only, so it would turn Julian 1500-02-29 into March 1 and give Julian dates
// the wrong weekday.
func gregorianLong(year, month, day int) string {
	if year > 1582 || (year == 1582 && (month > 10 || (month == 10 && day >= 15))) {
		return dateconverter.FormatGregorian(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
	}
	// Julian day number of a Julian calendar date; JDN + 1 mod 7 is the weekday.
	a := (14 - month) / 12
	y := year + 4800 - a
	m := month + 12*a - 3
	jdn := day + (153*m+2)/5 + 365*y + y/4 - 32083
	return fmt.Sprintf("%s, %s %d, %d", time.Weekday((jdn+1)%7), time.Month(month), day, year)
}
