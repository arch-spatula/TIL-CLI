/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// draftCmd represents the draft command
var draftCmd = &cobra.Command{
	Use:              "draft [OPTIONS]",
	TraverseChildren: true,
	Short:            "draft (today/tomorrow/retro) 3가지 파일 중 1개 생성",
	Long: `

- today:    오늘 TIL 템플릿을 생성합니다. 2311/TIL231112.md
- tomorrow: 내일 TIL 템플릿을 생성합니다. 2311/TIL231113.md
- retro:    회고 TIL 템플릿을 생성합니다. 2311/TIL231112Retro.md`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.Open("setting.json")
		if err != nil {
			panic(`setting.json 파일이 없습니다.

./TIL-CLI init 명령을 먼저 해주세요.`)
		}

		defer data.Close()

		byteValue, _ := io.ReadAll(data)

		var info map[string]interface{}
		json.Unmarshal([]byte(byteValue), &info)

		draft := info["draft"]

		today := draft.(map[string]interface{})["today"]
		tomorrow := draft.(map[string]interface{})["tomorrow"]
		retro := draft.(map[string]interface{})["retro"]
		none := draft.(map[string]interface{})["none"]
		empty := draft.(map[string]interface{})["empty"]

		fmt.Println("draft called", args, "today", today, "tomorrow", tomorrow, "retro", retro, "none", none, "empty", empty)
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
