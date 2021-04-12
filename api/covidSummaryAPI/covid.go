package api

import "time"

// CovidSummary JSON struct
type CovidSummary struct {
	Global struct {
		NewConfirmed   int `json:"NewConfirmed"`
		TotalConfirmed int `json:"TotalConfirmed"`
		NewDeaths      int `json:"NewDeaths"`
		TotalDeaths    int `json:"TotalDeaths"`
		NewRecovered   int `json:"NewRecovered"`
		TotalRecovered int `json:"TotalRecovered"`
	} `json:"Global"`
	Countries []struct {
		Country        string    `json:"Country"`
		CountryCode    string    `json:"CountryCode"`
		Slug           string    `json:"Slug"`
		NewConfirmed   int       `json:"NewConfirmed"`
		TotalConfirmed int       `json:"TotalConfirmed"`
		NewDeaths      int       `json:"NewDeaths"`
		TotalDeaths    int       `json:"TotalDeaths"`
		NewRecovered   int       `json:"NewRecovered"`
		TotalRecovered int       `json:"TotalRecovered"`
		Date           time.Time `json:"Date"`
	} `json:"Countries"`
	Date time.Time `json:"Date"`
}
