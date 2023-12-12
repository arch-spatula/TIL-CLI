/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/arch-spatula/TIL-CLI/utilFn/gitignoreWriter"
	"github.com/arch-spatula/TIL-CLI/utilFn/jsonReader"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "프로젝트 초기 설정 명령입니다.",
	Long: `.gitignore과 setting.json을 만들어 둡니다.

	- .gitignore가 이미 있으면 안 만듭니다.
	- setting.json가 이미 있으면 안 만듭니다.

setting.json를 읽고 다른 커맨드가 활용할 기준 파일을 즉 설정에 관한 파일을 만듭니다.`,
	Run: func(cmd *cobra.Command, args []string) {

		jsonReader.WriteJson()
		gitignoreWriter.WriteGitIgnore()

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
