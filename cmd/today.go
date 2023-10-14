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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "A brief description of your command",
	Long: `오늘을 기준으로 TIL 문서를 자동생성합니다.

이미 파일이 생성되어 있으면 실행하지 않습니다.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// 이번달 폴더 오늘 TIL 마크다운 파일이름 만들기
		folder := time.Now().Format("0601")
		markdown := time.Now().Format("20060102")

		err := os.Mkdir(folder, 0755)
		if err != nil {
			fmt.Printf("Unable to write file: %v\n", err)
		}

		todayMarkdownFile := folder + "/" + markdown + ".md"

		todayFile, err := os.Create(todayMarkdownFile)
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
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
