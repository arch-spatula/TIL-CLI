/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/arch-spatula/TIL-CLI/utilFn/weather"
	"github.com/spf13/cobra"
)

// experimentalCmd represents the experimental command
var experimentalCmd = &cobra.Command{
	Use:   "experimental",
	Short: "정식 기능이 아니고 실험적인 명령입니다.",
	Long: `정식 기능이 아니고 실험적인 명령입니다.

./TIL-CLI experimental
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("experimental called")
		weather.ReadWeather()
		weatherText := "날씨: "
		for _, HdayFcast := range weather.ReadWeather() {
			// 날짜를 입력
			date := time.Now().Format("20060102")
			// 조건부로 출력하기
			if date == HdayFcast.AplYmd {
				weatherText += HdayFcast.AmWetrTxt
				weatherText += " / "
				weatherText += HdayFcast.PmWetrTxt
				weatherText += "\n\n"
			}
		}
		fmt.Println(weatherText)
	},
}

func init() {
	rootCmd.AddCommand(experimentalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// experimentalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// experimentalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
