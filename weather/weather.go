package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Location struct {
	Name      string
	Country   string
	Localtime string
}

type Condition struct {
	Text string
}

type Current struct {
	TempCelsius      float64 `json:"temp_c"`
	FeelsLikeCelsius float64 `json:"feelslike_c"`
	Condition        Condition
}

type Weather struct {
	Location  Location
	Current   Current
	Condition Condition
}

func Get(weatherApi, loc string) Weather {
	req, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", weatherApi, loc))
	if err != nil {
		log.Println("Error getting weather data")
	}
	defer req.Body.Close()
	var w Weather
	err = json.NewDecoder(req.Body).Decode(&w)
	if err != nil {
		log.Println("Error decoding weather data")
	}

	return w
}
