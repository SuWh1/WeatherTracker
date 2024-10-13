package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type weatherCondition struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Temperature float64 `json:"temp"`
		Humidity    int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Weather []weatherCondition `json:"weather"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return apiConfigData{}, err
	}

	var cfg apiConfigData
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return apiConfigData{}, err
	}

	return cfg, nil
}

func renderTemplate(w http.ResponseWriter, city string, weather weatherData, unit string, errMessage string) {
	tmpl, err := template.ParseFiles("templates/weather.html")
	if err != nil {
		http.Error(w, "Could not load template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, struct {
		Weather weatherData
		Unit    string
		Error   string
	}{
		Weather: weather,
		Unit:    unit,
		Error:   errMessage,
	})
	if err != nil {
		http.Error(w, "Could not execute template: "+err.Error(), http.StatusInternalServerError)
	}
}

func query(city string, unit string) (weatherData, error) {
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}

	city = url.QueryEscape(city)
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?appid=%s&q=%s&units=%s", apiConfig.OpenWeatherMapApiKey, city, unit)

	resp, err := http.Get(url)
	if err != nil {
		return weatherData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return weatherData{}, fmt.Errorf("City not found!")
	}

	var data weatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return weatherData{}, err
	}

	return data, nil
}

func main() {
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		city := r.URL.Query().Get("city")
		unit := r.URL.Query().Get("unit")
		var data weatherData
		var errMessage string

		if city != "" && (unit == "metric" || unit == "imperial") {
			var err error
			data, err = query(city, unit)
			if err != nil {
				errMessage = err.Error()
			}
		}

		renderTemplate(w, city, data, unit, errMessage)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "", weatherData{}, "", "")
	})

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
