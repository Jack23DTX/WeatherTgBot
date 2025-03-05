package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"ProjectBot1/entity"
)

type Client struct {
	apiKey string
}

func NewClient(apiKey string) Client {
	return Client{
		apiKey: apiKey,
	}
}

func (c Client) Get(loc string) entity.Weather {
	req, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", c.apiKey, loc))
	if err != nil {
		log.Println("Error getting weather data")
	}
	defer req.Body.Close()
	var w entity.Weather
	err = json.NewDecoder(req.Body).Decode(&w)
	if err != nil {
		log.Println("Error decoding weather data")
	}

	return w
}
