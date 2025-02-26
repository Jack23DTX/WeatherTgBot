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

func Get(weatherApi, loc string) string {
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
	resp := fmt.Sprintf("%s, %s\nДата и время: %s\nТемпература: %.1f°C\nОщущается как: %.1f°C,\nПогодные условия: %s\n",
		w.Location.Name, w.Location.Country, w.Location.Localtime, w.Current.TempCelsius, w.Current.FeelsLikeCelsius, w.Current.Condition.Text)

	return resp
}
