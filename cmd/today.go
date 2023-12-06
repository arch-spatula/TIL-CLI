/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"

	"github.com/arch-spatula/TIL-CLI/utilFn/markdownReadAndWriter"
	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "오늘을 기준으로 TIL 문서를 자동생성합니다.",
	Long: `오늘을 기준으로 TIL 문서를 자동생성합니다.

이미 파일이 생성되어 있으면 실행하지 않습니다.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// 이번달 폴더 오늘 TIL 마크다운 파일이름 만들기

		markdownReadAndWriter.WriteMarkdown(time.Now())

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
