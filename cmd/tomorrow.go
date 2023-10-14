/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// tomorrowCmd represents the tomorrow command
var tomorrowCmd = &cobra.Command{
	Use:   "tomorrow",
	Short: "A brief description of your command",
	Long: `내일을 기준으로 TIL 문서를 자동생성합니다.

	이미 파일이 생성되어 있으면 실행하지 않습니다.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// 이번달 폴더 내일 TIL 마크다운 파일이름 만들기
		now := time.Now()
		folder := now.Format("0601")
		markdown := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC).Format("060102")

		if err := os.Mkdir(folder, 0755); !os.IsExist(err) {
			fmt.Println("이번달 폴더를 만들어두겠습니다.")
		}

		tomorrowMarkdownFile := folder + "/TIL" + markdown + ".md"

		if _, err := os.Stat(tomorrowMarkdownFile); os.IsNotExist(err) {
			todayFile, err := os.Create(tomorrowMarkdownFile)
			if err != nil {
				fmt.Printf("Unable to write file: %v\n", err)
			}
			defer todayFile.Close()

			// template.md 읽기
			template, err := os.ReadFile("template.md")
			if err != nil {
				fmt.Printf("Unable to read file: %v\n", err)
			}

			// 오늘 TIL에 쓰기
			fmt.Fprintln(todayFile, string(template))

			fmt.Println(tomorrowMarkdownFile, "을 만들어두겠습니다.")
		} else {
			fmt.Println(tomorrowMarkdownFile, "이 이미 만들어졌습니다.")
		}

	},
}

func init() {
	rootCmd.AddCommand(tomorrowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tomorrowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tomorrowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
