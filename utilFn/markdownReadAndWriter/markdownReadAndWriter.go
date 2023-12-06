package markdownReadAndWriter

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/arch-spatula/TIL-CLI/utilFn/common"
	"github.com/arch-spatula/TIL-CLI/utilFn/jsonReader"
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
		title := jsonReader.ReadJson().CurrentProject

		// 1일1커밋 무사고: 358일차
		days := common.DiffDays(jsonReader.ReadJson().DaysWithoutAccidentDay)
		daysWithoutAccident := "1일1커밋 무사고: " + strconv.Itoa(days) + "일차\n\n"

		gratificationDiary := jsonReader.ReadJson().GratificationFormat

		template := "# " + title + daysWithoutAccident + gratificationDiary

		// 오늘 TIL에 쓰기
		fmt.Fprintln(todayFile, string(template))

		fmt.Println(fileName, "을 만들어두겠습니다.")
	} else {
		fmt.Println(fileName, "이 이미 만들어졌습니다.")
	}
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
