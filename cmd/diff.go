/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "무사고일자를 알아냅니다.",
	Long:  `무사고일자를 알아냅니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.Open("setting.json")
		if err != nil {
			fmt.Println(err)
		}

		defer data.Close()

		byteValue, _ := io.ReadAll(data)

		var info map[string]interface{}
		json.Unmarshal([]byte(byteValue), &info)

		daysWithoutAccident := strings.Split(fmt.Sprint(info["days-without-accident"]), "-")

		yyyy := 1000
		mm := 1
		dd := 1
		for i, v := range daysWithoutAccident {
			if i == 0 {
				yyyy, err = strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
			}
			if i == 1 {
				mm, err = strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
			}
			if i == 2 {
				dd, err = strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
			}
		}

		t1 := Date(yyyy, mm, dd)
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

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diffCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
