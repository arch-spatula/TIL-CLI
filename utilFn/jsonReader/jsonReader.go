package jsonReader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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

func WriteJson() {
	fmt.Println("hello writer json")
}

func ReadJson() TILConfig {
	data, err := os.Open("til-config.json")
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
