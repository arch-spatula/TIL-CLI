/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"

	"github.com/arch-spatula/TIL-CLI/utilFn/markdownReadAndWriter"
	"github.com/spf13/cobra"
)

// draftCmd represents the draft command
var draftCmd = &cobra.Command{
	Use:              "draft [OPTIONS]",
	TraverseChildren: true,
	Short:            "draft (today/tomorrow/retro) 3가지 파일 중 1개 생성",
	Long: `예를 들어 오늘이 2023년 11월 12일이면 다음처첨 생성합니다.

- today:    오늘 TIL 템플릿을 생성합니다. 2311/TIL231112.md
- tomorrow: 내일 TIL 템플릿을 생성합니다. 2311/TIL231113.md
- retro:    회고 TIL 템플릿을 생성합니다. 2311/TIL231112Retro.md`,
	Run: func(cmd *cobra.Command, args []string) {

		// 하위 플래그를 지정하지 않으면 차단
		if len(args) < 1 {
			panic(`현재 draft 뒤에 입력한 flag가 없습니다.

draft 뒤에 today, tomorrow, retro 중 하나를 입력해주세요.

./TIL-CLI draft today`)
		}

		// 실행 차단
		// return
		// til-config.json에 없는 키워드 접근하면 차단
		key := args[0]
		if key == "today" {
			markdownReadAndWriter.WriteMarkdown(time.Now())
		}
		if key == "tomorrow" {
			markdownReadAndWriter.WriteMarkdown(time.Now().AddDate(0, 0, 1))
		}
		if key == "retro" {
			// @todo: RetroW 뒤에 붙이기 추가
			markdownReadAndWriter.WriteMarkdown(time.Now())
		}
	},
}

func init() {
	rootCmd.AddCommand(draftCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// draftCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// draftCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
