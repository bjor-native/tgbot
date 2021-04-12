package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// TotalRecovered отдает общее кол-во людей в мире которые вылечились от ковид
func TotalRecovered() string {
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
	return strconv.Itoa(covid.Global.TotalRecovered)
}
