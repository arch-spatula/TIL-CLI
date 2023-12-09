package markdownReadAndWriter

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/arch-spatula/TIL-CLI/utilFn/common"
	"github.com/arch-spatula/TIL-CLI/utilFn/jsonReader"
	"github.com/arch-spatula/TIL-CLI/utilFn/weather"
)

func WriteMarkdown(createTime time.Time) {
	folder := createTime.Format("0601")
	markdown := createTime.Format("060102")

	if err := os.Mkdir(folder, 0755); !os.IsExist(err) {
		fmt.Println("이번달 폴더를 만들어두겠습니다.")
	}

	fileName := folder + "/TIL" + markdown + ".md"

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		todayFile, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("Unable to write file: %v\n", err)
		}
		defer todayFile.Close()

		// # vue로 만드는 디자인 시스템 23일차
		text := jsonReader.ReadJson().CurrentProject
		key := "current-project-start-day"
		diffDays := strconv.Itoa(common.DiffDays(jsonReader.ReadJson().CurrentProjectStartDay, createTime))
		title := jsonReader.ParseToKey(text, key, diffDays)

		// 1일1커밋 무사고: 358일차
		daysWithoutAccident := jsonReader.ParseToKey(jsonReader.ReadJson().DaysWithoutAccidentFormat, "days-without-accident-day", strconv.Itoa(common.DiffDays(jsonReader.ReadJson().DaysWithoutAccidentDay, createTime)))

		// 날씨: 맑음 / 맑음
		foo := "날씨: "
		for _, HdayFcast := range weather.ReadWeather().HdayFcastList {
			// 날짜를 입력
			date := createTime.Format("20060102")
			// 조건부로 출력하기
			if date == HdayFcast.AplYmd {
				foo += HdayFcast.AmWetrTxt
				foo += " / "
				foo += HdayFcast.PmWetrTxt
				foo += "\n\n"
			}
		}

		// 감사일기
		gratificationDiary := jsonReader.ReadJson().GratificationFormat

		// todo
		todo := jsonReader.ReadJson().TodoFormat

		template := "# " + title + daysWithoutAccident + foo + gratificationDiary + todo

		// 오늘 TIL에 쓰기
		fmt.Fprintln(todayFile, string(template))

		fmt.Println(fileName, "을 만들어두겠습니다.")
	} else {
		fmt.Println(fileName, "이 이미 만들어졌습니다.")
	}
}

func WriteRetro(createTime time.Time, retroKind string) {
	folder := createTime.Format("0601")
	markdown := createTime.Format("060102")

	if err := os.Mkdir(folder, 0755); !os.IsExist(err) {
		fmt.Println("이번달 폴더를 만들어두겠습니다.")
	}

	fileName := folder + "/TIL" + markdown + "Retro" + retroKind + ".md"

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		todayFile, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("Unable to write file: %v\n", err)
		}
		defer todayFile.Close()
		title := jsonReader.ParseToKey(jsonReader.ReadJson().CurrentProject, "current-project-start-day", strconv.Itoa(common.DiffDays(jsonReader.ReadJson().CurrentProjectStartDay, createTime)))

		daysWithoutAccident := jsonReader.ParseToKey(jsonReader.ReadJson().DaysWithoutAccidentFormat, "days-without-accident-day", strconv.Itoa(common.DiffDays(jsonReader.ReadJson().DaysWithoutAccidentDay, createTime)))

		// 날씨: 맑음 / 맑음
		foo := "날씨: "
		for _, HdayFcast := range weather.ReadWeather().HdayFcastList {
			// 날짜를 입력
			date := createTime.Format("20060102")
			// 조건부로 출력하기
			if date == HdayFcast.AplYmd {
				foo += HdayFcast.AmWetrTxt
				foo += " / "
				foo += HdayFcast.PmWetrTxt
				foo += "\n\n"
			}
		}

		gratificationDiary := jsonReader.ReadJson().GratificationFormat

		var todo = ""

		switch retroKind {
		case "W":
			todo = "- [ ] 주간회고\n- [ ] ???\n\n---\n\n"
		case "M":
			todo = "- [ ] 월간 & 주간회고\n- [ ] ???\n\n---\n\n"
		case "Q":
			todo = "- [ ] 분기별 & 주간회고\n- [ ] ???\n\n---\n\n"
		default:
			panic("없는 플래그입니다.")
		}

		retro := jsonReader.ReadJson().RetroFormat

		template := "# " + title + daysWithoutAccident + foo + gratificationDiary + todo + retro

		fmt.Fprintln(todayFile, string(template))

		fmt.Println(fileName, "을 만들어두겠습니다.")
	} else {
		fmt.Println(fileName, "이 이미 만들어졌습니다.")
	}
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
