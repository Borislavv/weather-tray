package builder

import (
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	"github.com/Borislavv/weather-tray/internal/domain/dto/api"
	getterDto "github.com/Borislavv/weather-tray/internal/domain/dto/getter"
)

type WeatherAgg interface {
	BuildFromResponseDto(response getterDto.WeatherResponse, location agg.Location) agg.Weather
}

type LocationAgg interface {
	BuildFromRaw(title string, latitude float64, longitude float64) agg.Location
	BuildFromDto(request apiDto.LocationRequest) agg.Location
}
