package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Joke struct
type Joke struct {
	ID   uint32 `json:"id"`
	Joke string `json:"joke"`
}

//JokeResponse struct for API joke
type JokeResponse struct {
	Type  string `json:"type"`
	Value Joke   `json:"value"`
}

// GetJoke отдает шутку про Чака
func GetJoke() string {
	c := http.Client{}
	resp, err := c.Get("http://api.icndb.com/jokes/random?limitTo=[nerdy]")
	if err != nil {
		return "jokes API not responding"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	joke := JokeResponse{}

	err = json.Unmarshal(body, &joke)
	if err != nil {
		return "Joke error"
	}
	return joke.Value.Joke
}
