/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/arch-spatula/TIL-CLI/utilFn/jsonReader"
	"github.com/spf13/cobra"
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "무사고일자를 알아냅니다.",
	Long:  `무사고일자를 알아냅니다.`,
	Run: func(cmd *cobra.Command, args []string) {

		daysWithoutAccident := strings.Split(fmt.Sprint(jsonReader.ReadJson().DaysWithoutAccidentDay), "-")

		dateBuffer := [3]int{1000, 1, 1}
		for i, v := range daysWithoutAccident {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			dateBuffer[i] = num

		}

		t1 := Date(dateBuffer[0], dateBuffer[1], dateBuffer[2])
		t2 := Date(time.Now().Year(), int(time.Now().Month()), time.Now().Day())
		days := t2.Sub(t1).Hours() / 24
		fmt.Println(days)
	},
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func init() {
	rootCmd.AddCommand(diffCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diffCmd.PersistentFlags().String("foo", "", "A help for foo")

	diffCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diffCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
