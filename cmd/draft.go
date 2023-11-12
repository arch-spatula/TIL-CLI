/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

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

		// 하위 플래그를 지정하지 않으면 차단
		if len(args) != 1 {
			panic(`현재 draft 뒤에 입력한 flag가 없습니다.

draft 뒤에 today, tomorrow, retro 중 하나를 입력해주세요.

./TIL-CLI draft today`)
		}

		// setting.json에 없는 키워드 접근하면 차단
		key := args[0]
		if settingText, ok := draft.(map[string]interface{})[key]; ok {
			// 이번달 폴더 오늘 TIL 마크다운 파일이름 만들기
			folder := time.Now().Format("0601")
			markdown := time.Now().Format("060102")

			if err := os.Mkdir(folder, 0755); !os.IsExist(err) {
				fmt.Println("이번달 폴더를 만들어두겠습니다.")
			}

			markdownFileName := folder + "/TIL" + markdown + ".md"

			if _, err := os.Stat(markdownFileName); os.IsNotExist(err) {
				markdownFile, err := os.Create(markdownFileName)
				if err != nil {
					fmt.Printf("Unable to write file: %v\n", err)
				}
				defer markdownFile.Close()

				// 오늘 TIL에 쓰기
				fmt.Fprintln(markdownFile, string(fmt.Sprint(settingText)))

				fmt.Println(markdownFileName, "을 만들어두겠습니다.")
			} else {
				fmt.Println(markdownFileName, "이 이미 만들어졌습니다.")
			}

		} else {
			panic(`draft 뒤에 today, tomorrow, retro 중 하나를 입력해주세요

./TIL-CLI draft today`)
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
