package dto

type WeatherResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		WindSpeed   float64 `json:"windspeed"`
	} `json:"current_weather"`
}
