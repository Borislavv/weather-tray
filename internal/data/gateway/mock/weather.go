package gateway

import (
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	"github.com/Borislavv/weather-tray/internal/domain/dto/getter"
)

type WeatherMockGateway struct{}

func NewWeather() *WeatherMockGateway {
	return &WeatherMockGateway{}
}

func (w *WeatherMockGateway) Get(location agg.Location) (*getterDto.WeatherResponse, error) {
	return &getterDto.WeatherResponse{
		CurrentWeather: struct {
			Temperature float64 `json:"temperature"`
			WindSpeed   float64 `json:"windspeed"`
		}(struct {
			Temperature float64
			WindSpeed   float64
		}{Temperature: 7.08, WindSpeed: 2.56}),
	}, nil
}
