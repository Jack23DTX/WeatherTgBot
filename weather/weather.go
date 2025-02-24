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
	Temp_c      float64
	Feelslike_c float64
	Condition   Condition
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
		w.Location.Name, w.Location.Country, w.Location.Localtime, w.Current.Temp_c, w.Current.Feelslike_c, w.Current.Condition.Text)

	return resp
}
