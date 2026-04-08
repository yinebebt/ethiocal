package cmd

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

var gtoeCmd = &cobra.Command{
	Use:   "gtoe [year] [month] [day]",
	Short: "Convert Gregorian date to Ethiopian date",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		year, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: invalid year:", args[0])
			return
		}
		month, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: invalid month:", args[1])
			return
		}
		day, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error: invalid day:", args[2])
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
		year, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: invalid year:", args[0])
			return
		}
		month, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: invalid month:", args[1])
			return
		}
		day, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error: invalid day:", args[2])
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
