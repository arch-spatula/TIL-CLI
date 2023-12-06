package jsonReader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type TILConfig struct {
	CurrentProject         string `json:"current-project"`
	CurrentProjectStartDay string `json:"current-project-start-day"`
	ShowCurrentProject     bool   `json:"show-current-project"`
	DaysWithoutAccidentDay string `json:"days-without-accident-day"`
	DaysWithoutAccident    bool   `json:"days-without-accident"`
	GratificationFormat    string `json:"gratification-format"`
	GratificationDiary     bool   `json:"gratification-diary"`
	Draft                  struct {
		Today    string `json:"today"`
		Tomorrow string `json:"tomorrow"`
		Retro    string `json:"retro"`
	} `json:"draft"`
}

var settingJsonFileName = "til-config.json"

func WriteJson() {
	if _, err := os.Stat(settingJsonFileName); os.IsNotExist(err) {
		settingJsonFile, err := os.Create(settingJsonFileName)
		if err != nil {
			fmt.Printf("Unable to write file: %v\n", err)
		}
		defer settingJsonFile.Close()
		formattedNow := time.Now().Format("2006-01-02")

		var tilConfig TILConfig

		tilConfig.CurrentProject = "진행 중인 프로젝트를 입력해주세요. 지금은 {{current-project-start-day}}일차입니다.\n\n"
		tilConfig.CurrentProjectStartDay = formattedNow
		tilConfig.ShowCurrentProject = true
		tilConfig.GratificationFormat = "## 감사일기\n\n1. ???\n\n"
		tilConfig.GratificationDiary = true
		tilConfig.DaysWithoutAccidentDay = formattedNow
		tilConfig.DaysWithoutAccident = true
		tilConfig.Draft.Retro = ""
		tilConfig.Draft.Today = ""
		tilConfig.Draft.Tomorrow = ""

		doc, err := json.Marshal(tilConfig)
		if err != nil {
			panic(err)
		}
		os.WriteFile(settingJsonFileName, doc, os.FileMode(0644))

		fmt.Println(settingJsonFileName, "을 만들어두겠습니다.")
	} else {
		fmt.Println(settingJsonFileName, "이 이미 만들어졌습니다.")
	}
}

func ReadJson() TILConfig {
	data, err := os.Open(settingJsonFileName)
	if err != nil {
		panic(`til-config 파일이 없습니다.

./TIL-CLI init 명령을 먼저 해주세요.`)
	}
	defer data.Close()

	byteValue, _ := io.ReadAll(data)

	var tilConfig TILConfig
	err = json.Unmarshal(byteValue, &tilConfig)
	if err != nil {
		panic(err)
	}

	return tilConfig
}

func ParseToKey(text, key, val string) string {

	return strings.ReplaceAll(text, "{{"+key+"}}", val)
}
