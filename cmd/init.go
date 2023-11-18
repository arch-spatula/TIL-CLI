/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

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

		// utilfn.ReadJson()
		jsonReader.ReadJson()
		jsonReader.WriteJson()

		settingJsonFileName := "til-config.json"

		if _, err := os.Stat(settingJsonFileName); os.IsNotExist(err) {
			settingJsonFile, err := os.Create(settingJsonFileName)
			if err != nil {
				fmt.Printf("Unable to write file: %v\n", err)
			}
			defer settingJsonFile.Close()

			s := time.Now().Format("2006-01-02")
			today := fmt.Sprintf(`{
	"current-project": "진행 중인 프로젝트를 입력해주세요. 지금은 {current-project-start-day}일차입니다.\n\n",
	"current-project-start-day": "%s",
	"show-current-project": true,
	"days-without-accident-day": "%s",
	"days-without-accident": true,
	"gratification-format": "## 감사일기\n\n1. ???\n\n",
	"gratification-diary": true,
	"draft": {
		"today": "",
		"tomorrow": "",
		"retro": ""
	}
}`, s, s)

			fmt.Fprintln(settingJsonFile, string(today))

			fmt.Println(settingJsonFileName, "을 만들어두겠습니다.")
		} else {
			fmt.Println(settingJsonFileName, "이 이미 만들어졌습니다.")
		}

		gitignoreFileName := ".gitignore"
		if _, err := os.Stat(gitignoreFileName); os.IsNotExist(err) {
			settingJsonFile, err := os.Create(gitignoreFileName)
			if err != nil {
				fmt.Printf("Unable to write file: %v\n", err)
			}
			defer settingJsonFile.Close()

			fmt.Fprintln(settingJsonFile, string(`# Ignore all
*

# Unignore all with extensions
!*.*

# Unignore all dirs
!*/

### Above combination will ignore all files without extension ###

# Ignore files with extension .class & .sm
*.class
*.sm

# Ignore bin dir
bin/
# or
*/bin/*

# Unignore all .jar in bin dir
!*/bin/*.jar

# Ignore all library.jar in bin dir
*/bin/library.jar

# Ignore a file with extension
relative/path/to/dir/filename.extension

# Ignore a file without extension
template.md
setting.json`))

			fmt.Println(gitignoreFileName, "을 만들어두겠습니다.")
		} else {
			fmt.Println(gitignoreFileName, "이 이미 만들어졌습니다.")
		}
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
