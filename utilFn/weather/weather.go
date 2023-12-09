package weather

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type ThisWeekWeather struct {
	NowFcast struct {
		NaverRgnCd     string      `json:"naverRgnCd"`
		LareaNm        string      `json:"lareaNm"`
		MareaNm        string      `json:"mareaNm"`
		SareaNm        string      `json:"sareaNm"`
		CpName         string      `json:"cpName"`
		AplYmd         string      `json:"aplYmd"`
		AplTm          string      `json:"aplTm"`
		AplYmdt        string      `json:"aplYmdt"`
		WetrCd         string      `json:"wetrCd"`
		WetrTxt        string      `json:"wetrTxt"`
		Tmpr           float64     `json:"tmpr"`
		Ytmpr          float64     `json:"ytmpr"`
		Stmpr          float64     `json:"stmpr"`
		RainProb       interface{} `json:"rainProb"`
		OneHourRainAmt string      `json:"oneHourRainAmt"`
		WindDrctn      string      `json:"windDrctn"`
		WindSpd        float64     `json:"windSpd"`
		Humd           int         `json:"humd"`
		FcastYmdt      string      `json:"fcastYmdt"`
		RainAmt        interface{} `json:"rainAmt"`
		SnowAmt        interface{} `json:"snowAmt"`
		LastUdtYmdt    string      `json:"lastUdtYmdt"`
		WindDrctnName  string      `json:"windDrctnName"`
		FullAreaName   string      `json:"fullAreaName"`
	} `json:"nowFcast"`
	Uv struct {
		NaverRegionCode  string `json:"naverRegionCode"`
		LifeIndexCode    string `json:"lifeIndexCode"`
		LifeIndexName    string `json:"lifeIndexName"`
		FcastYmdt        string `json:"fcastYmdt"`
		AplYmdt          string `json:"aplYmdt"`
		Grade            int    `json:"grade"`
		Figure           string `json:"figure"`
		LabelText        string `json:"labelText"`
		GuideTextSummary string `json:"guideTextSummary"`
		GuideTextDetail  string `json:"guideTextDetail"`
		LegendDataList   []struct {
			LifeIndexCode string      `json:"lifeIndexCode"`
			NowGrade      int         `json:"nowGrade"`
			MaxGrade      int         `json:"maxGrade"`
			LabelText     string      `json:"labelText"`
			MinFigure     string      `json:"minFigure"`
			MaxFigure     interface{} `json:"maxFigure"`
			LevelClass    string      `json:"levelClass"`
		} `json:"legendDataList"`
		ClickCode        string  `json:"clickCode"`
		DonutDegreeValue float64 `json:"donutDegreeValue"`
		LevelClass       string  `json:"levelClass"`
	} `json:"uv"`
	SunRiseTm         string `json:"sunRiseTm"`
	AirDailyTrendList []struct {
		NaverRgnCd            string  `json:"naverRgnCd"`
		AplYmd                string  `json:"aplYmd"`
		AplYmdt               string  `json:"aplYmdt"`
		Pm10Avg               float64 `json:"pm10Avg"`
		Pm10AvgLegend1        string  `json:"pm10AvgLegend1"`
		Pm25Avg               float64 `json:"pm25Avg"`
		Pm25AvgLegend1        string  `json:"pm25AvgLegend1"`
		IsForecastPM          bool    `json:"isForecastPM"`
		Pm10AvgGrade          string  `json:"pm10AvgGrade"`
		Pm10AvgLevelClass     string  `json:"pm10AvgLevelClass"`
		Pm25AvgGrade          string  `json:"pm25AvgGrade"`
		Pm25AvgLevelClass     string  `json:"pm25AvgLevelClass"`
		Pm10AvgConvertPercent float64 `json:"pm10AvgConvertPercent"`
		Pm25AvgConvertPercent float64 `json:"pm25AvgConvertPercent"`
		DayString             string  `json:"dayString"`
		FullAreaName          string  `json:"fullAreaName"`
	} `json:"airDailyTrendList"`
	HdayFcastList []struct {
		NaverRgnCd   string      `json:"naverRgnCd"`
		LareaNm      string      `json:"lareaNm"`
		MareaNm      string      `json:"mareaNm"`
		SareaNm      string      `json:"sareaNm"`
		CpName       interface{} `json:"cpName"`
		AplYmd       string      `json:"aplYmd"`
		AplTm        interface{} `json:"aplTm"`
		AplYmdt      string      `json:"aplYmdt"`
		FcastYmdt    string      `json:"fcastYmdt"`
		WetrCd       interface{} `json:"wetrCd"`
		AmWetrCd     string      `json:"amWetrCd"`
		PmWetrCd     string      `json:"pmWetrCd"`
		WetrTxt      interface{} `json:"wetrTxt"`
		AmWetrTxt    string      `json:"amWetrTxt"`
		PmWetrTxt    string      `json:"pmWetrTxt"`
		MinTmpr      float64     `json:"minTmpr"`
		MaxTmpr      float64     `json:"maxTmpr"`
		AmRainProb   int         `json:"amRainProb"`
		PmRainProb   int         `json:"pmRainProb"`
		DayString    string      `json:"dayString"`
		TmprRange    float64     `json:"tmprRange"`
		FullAreaName string      `json:"fullAreaName"`
	} `json:"hdayFcastList"`
	SunSetTm string `json:"sunSetTm"`
	AirFcast struct {
		NaverRgnCd                string      `json:"naverRgnCd"`
		LareaNm                   string      `json:"lareaNm"`
		MareaNm                   string      `json:"mareaNm"`
		SareaNm                   string      `json:"sareaNm"`
		CpName                    interface{} `json:"cpName"`
		AplYmd                    string      `json:"aplYmd"`
		AplTm                     string      `json:"aplTm"`
		AplYmdt                   string      `json:"aplYmdt"`
		StationID                 string      `json:"stationId"`
		StationName               string      `json:"stationName"`
		StationAddress            interface{} `json:"stationAddress"`
		StationO3                 float64     `json:"stationO3"`
		StationO3Legend1          string      `json:"stationO3Legend1"`
		StationPM10               int         `json:"stationPM10"`
		StationPM10Aqi            int         `json:"stationPM10Aqi"`
		StationPM10Legend1        string      `json:"stationPM10Legend1"`
		StationPM25               int         `json:"stationPM25"`
		StationPM25Aqi            int         `json:"stationPM25Aqi"`
		StationPM25Legend1        string      `json:"stationPM25Legend1"`
		StationKhai               int         `json:"stationKhai"`
		StationKhaiLegend1        string      `json:"stationKhaiLegend1"`
		StationSo2                float64     `json:"stationSo2"`
		StationSo2Legend1         string      `json:"stationSo2Legend1"`
		StationCo                 float64     `json:"stationCo"`
		StationCoLegend1          string      `json:"stationCoLegend1"`
		StationNo2                float64     `json:"stationNo2"`
		StationNo2Legend1         string      `json:"stationNo2Legend1"`
		ObsYmdt                   string      `json:"obsYmdt"`
		Pm25ExistYn               string      `json:"pm25ExistYn"`
		StationLatitude           float64     `json:"stationLatitude"`
		StationLongitude          float64     `json:"stationLongitude"`
		StationGubun              string      `json:"stationGubun"`
		StationNaverRgnCd         string      `json:"stationNaverRgnCd"`
		StationNaverRgnName       string      `json:"stationNaverRgnName"`
		ZoomLevel                 int         `json:"zoomLevel"`
		FcastYmdt                 string      `json:"fcastYmdt"`
		StationPM10LevelClass     string      `json:"stationPM10LevelClass"`
		StationPM25LevelClass     string      `json:"stationPM25LevelClass"`
		StationPM10LegendNum      string      `json:"stationPM10LegendNum"`
		StationPM10ConvertPercent float64     `json:"stationPM10ConvertPercent"`
		StationPM25LegendNum      string      `json:"stationPM25LegendNum"`
		StationPM25ConvertPercent float64     `json:"stationPM25ConvertPercent"`
		FullAreaName              string      `json:"fullAreaName"`
	} `json:"airFcast"`
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
	startDoc := strings.Split(doc, "weatherSummary = ")[1]
	endDoc := strings.Split(startDoc, ";")[0]

	var thisWeekWeather ThisWeekWeather
	err = json.Unmarshal([]byte(endDoc), &thisWeekWeather)
	if err != nil {
		panic(err)
	}

	return thisWeekWeather
}
