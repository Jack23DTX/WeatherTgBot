package entity

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
