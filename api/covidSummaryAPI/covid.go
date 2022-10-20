package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// CovidSummary JSON struct
type CovidSummary struct {
	Global struct {
		NewConfirmed   int64 `json:"NewConfirmed"`
		TotalConfirmed int64 `json:"TotalConfirmed"`
		NewDeaths      int64 `json:"NewDeaths"`
		TotalDeaths    int64 `json:"TotalDeaths"`
		NewRecovered   int64 `json:"NewRecovered"`
		TotalRecovered int64 `json:"TotalRecovered"`
	} `json:"Global"`
	Countries []struct {
		Country        string    `json:"Country"`
		CountryCode    string    `json:"CountryCode"`
		Slug           string    `json:"Slug"`
		NewConfirmed   int64     `json:"NewConfirmed"`
		TotalConfirmed int64     `json:"TotalConfirmed"`
		NewDeaths      int64     `json:"NewDeaths"`
		TotalDeaths    int64     `json:"TotalDeaths"`
		NewRecovered   int64     `json:"NewRecovered"`
		TotalRecovered int64     `json:"TotalRecovered"`
		Date           time.Time `json:"Date"`
	} `json:"Countries"`
	Date time.Time `json:"Date"`
}

func GetCovidStatistic(value string) string {
	c := http.Client{}
	resp, err := c.Get("https://api.covid19api.com/summary")
	if err != nil {
		return "covid-19 API not responding"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	covid := CovidSummary{}

	err = json.Unmarshal(body, &covid)
	if err != nil {
		return "error :("
	}

	

	switch value {
	case "newRecovered":
		return strconv.FormatInt(covid.Global.NewRecovered, 10)
	case "totalRecovered":
		return strconv.FormatInt(covid.Global.TotalRecovered, 10)
	case "newConfirmed":
		return strconv.FormatInt(covid.Global.NewConfirmed, 10)
	case "totalConfirmed":
		return strconv.FormatInt(covid.Global.TotalConfirmed, 10)
	case "newDeaths":
		return strconv.FormatInt(covid.Global.NewDeaths, 10)
	case "totalDeaths":
		return strconv.FormatInt(covid.Global.TotalDeaths, 10)
	}
	return "OOPS!"
}
