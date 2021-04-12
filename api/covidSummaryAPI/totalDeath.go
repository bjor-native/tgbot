package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// TotalDeath отдает кол-во смертей во всем мире
func TotalDeath() string {
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
	return strconv.Itoa(covid.Global.TotalDeaths)
}
