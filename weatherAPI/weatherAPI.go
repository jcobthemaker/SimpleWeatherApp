package main

import (
	 "weatherAPI/urlBuilder"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	
	params := map[string]string{
		"q":  "Lodz",
		"api": "no",
	}

	apiUrl, _ := urlBuilder.BuildAPIUrl(params)

	res, err := http.Get(apiUrl)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != 200 {
		panic("Could not find API")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Response
	err = json.Unmarshal(body, &weather)
	
	if err != nil {
		panic(err)
	}

	location, current := weather.Location, weather.Current
	fmt.Printf(
		"%s, %s: %.0fC, %s \n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)

	defer res.Body.Close()
}


type Response struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
		TimeZone string `json:"tz_id"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		WindKPH float64 `json:"wind_kph"`
		WindDirection string `json:"wind_dir"`
		Humidity float64`json:"humidity"`
	} `json:"current"`


}


