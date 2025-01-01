package httpPackage

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"weatherAPI/urlBuilder"
)

func getApiResponse(params map[string]string) *Response {
	apiUrl, err := urlBuilder.BuildAPIUrl(params)

	if err != nil {
		log.SetPrefix("BuildAPIUrl(): ")
		log.Print(err)
		return nil
	}

	res, err := http.Get(apiUrl)
	if err != nil {
		log.SetPrefix("getApiResponse(): ")
		log.Print(err)
		return nil
	}

	if res.StatusCode != 200 {
		log.SetPrefix("getApiResponse(): ")
		log.Print("Could not find API")
		return nil
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.SetPrefix("getApiResponse(): ")
		log.Print(err)
		return nil
	}

	var weather Response
	err = json.Unmarshal(body, &weather)

	if err != nil {
		log.SetPrefix("getApiResponse(): ")
		log.Print(err)
		return nil
	}

	return &weather
}

func RunServer() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/api/process", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			body, err := io.ReadAll(r.Body)

			if err != nil {
				http.Error(w, "Unable to read request body", http.StatusBadRequest)
				log.Print(err)
				return
			}

			var requestData RequestData
			err = json.Unmarshal(body, &requestData)

			if err != nil {
				http.Error(w, "Invalid JSON format", http.StatusBadRequest)
				log.Print(err)
				return
			}

			params := map[string]string{
				"q":   requestData.City,
				"api": "no",
			}

			responseData := getApiResponse(params)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responseData)
		}
	})

	http.ListenAndServe(":8080", nil)
}

type RequestData struct {
	City string `json:"city"`
}

type Response struct {
	Location struct {
		Name     string `json:"name"`
		Country  string `json:"country"`
		TimeZone string `json:"tz_id"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		WindKPH       float64 `json:"wind_kph"`
		WindDirection string  `json:"wind_dir"`
		Humidity      float64 `json:"humidity"`
	} `json:"current"`
}
