package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yinebebt/ethiocal/dateconverter"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert dates between Ethiopian and Gregorian calendars",
}

// parseYMD parses the [year month day] positional args, naming the first
// non-integer field in the error.
func parseYMD(args []string) (year, month, day int, err error) {
	fields := []struct {
		name string
		out  *int
	}{
		{"year", &year},
		{"month", &month},
		{"day", &day},
	}
	for i, f := range fields {
		if *f.out, err = strconv.Atoi(args[i]); err != nil {
			return 0, 0, 0, fmt.Errorf("invalid %s: %s", f.name, args[i])
		}
	}
	return year, month, day, nil
}

var gtoeCmd = &cobra.Command{
	Use:   "gtoe [year] [month] [day]",
	Short: "Convert Gregorian date to Ethiopian date",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		year, month, day, err := parseYMD(args)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		etDate, err := dateconverter.Ethiopian(year, month, day)
		if err != nil {
			fmt.Println("Error converting date:", err)
			return
		}

		fmt.Printf("\nGregorian Date: %04d-%02d-%02d\n", year, month, day)
		fmt.Println("Converted Ethiopian Date:", etDate.Format("2006-01-02"))
	},
}

var etogCmd = &cobra.Command{
	Use:   "etog [year] [month] [day]",
	Short: "Convert Ethiopian date to Gregorian date",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		year, month, day, err := parseYMD(args)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		gregDate, err := dateconverter.Gregorian(year, month, day)
		if err != nil {
			fmt.Println("Error converting date:", err)
			return
		}
		fmt.Printf("\nEthiopian Date: %04d-%02d-%02d\n", year, month, day)
		fmt.Println("Converted Gregorian Date:", gregDate.Format("2006-01-02"))
	},
}

func init() {
	convertCmd.AddCommand(gtoeCmd)
	convertCmd.AddCommand(etogCmd)
}
