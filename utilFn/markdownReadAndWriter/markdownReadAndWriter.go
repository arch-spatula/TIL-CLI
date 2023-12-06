package markdownReadAndWriter

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/arch-spatula/TIL-CLI/utilFn/jsonReader"
)

func WriteMarkdown(fileName string) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		todayFile, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("Unable to write file: %v\n", err)
		}
		defer todayFile.Close()

		// # vue로 만드는 디자인 시스템 23일차
		title := jsonReader.ReadJson().CurrentProject

		// 1일1커밋 무사고: 358일차
		daysWithoutAccidentStrFromJson := strings.Split(fmt.Sprint(jsonReader.ReadJson().DaysWithoutAccidentDay), "-")

		dateBuffer := [3]int{1000, 1, 1}
		for i, v := range daysWithoutAccidentStrFromJson {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			dateBuffer[i] = num

		}

		t1 := Date(dateBuffer[0], dateBuffer[1], dateBuffer[2])
		t2 := Date(time.Now().Year(), int(time.Now().Month()), time.Now().Day())
		days := t2.Sub(t1).Hours() / 24
		daysWithoutAccident := "1일1커밋 무사고: " + strconv.Itoa(int(days)) + "일차\n\n"

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
