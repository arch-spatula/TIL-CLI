package common

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 2023-12-06 명령일 2023-12-8 2반환
func DiffDays(startDate string) int {
	splitDate := strings.Split(fmt.Sprint(startDate), "-")
	dateBuffer := [3]int{1000, 1, 1}
	for i, v := range splitDate {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		dateBuffer[i] = num

	}
	t1 := Date(dateBuffer[0], dateBuffer[1], dateBuffer[2])
	t2 := Date(time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	days := t2.Sub(t1).Hours() / 24
	return int(days)
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
