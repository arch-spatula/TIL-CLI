package weather

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type ThisWeekWeather []struct {
	NaverRgnCd string  `json:"naverRgnCd"`
	LareaNm    string  `json:"lareaNm"`
	MareaNm    string  `json:"mareaNm"`
	SareaNm    string  `json:"sareaNm"`
	AplYmd     string  `json:"aplYmd"`
	AmWetrCd   string  `json:"amWetrCd"`
	AmWetrTxt  string  `json:"amWetrTxt"`
	AmRainProb int     `json:"amRainProb"`
	PmWetrCd   string  `json:"pmWetrCd"`
	PmWetrTxt  string  `json:"pmWetrTxt"`
	PmRainProb int     `json:"pmRainProb"`
	MinTmpr    float64 `json:"minTmpr"`
	MaxTmpr    float64 `json:"maxTmpr"`
	FcastYmdt  string  `json:"fcastYmdt"`
	DayString  string  `json:"dayString"`
}

func ReadWeather() ThisWeekWeather {
	res, err := http.Get("https://weather.naver.com/")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("서비스 혹은 엔드포인트 이용 불가")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	doc := string(data)
	startDoc := strings.Split(doc, `domesticWeeklyFcastList":`)[1]
	endDoc := strings.Split(startDoc, "},\"sliderImagePlay~~2")[0]

	var thisWeekWeather ThisWeekWeather
	err = json.Unmarshal([]byte(endDoc), &thisWeekWeather)
	if err != nil {
		panic(err)
	}

	return thisWeekWeather
}
