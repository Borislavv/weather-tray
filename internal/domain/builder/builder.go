package builder

import (
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	"github.com/Borislavv/weather-tray/internal/domain/dto/api"
	apiRequestDto "github.com/Borislavv/weather-tray/internal/domain/dto/getter"
)

type Weather interface {
	BuildFromResponseDto(response apiRequestDto.WeatherResponse, location agg.Location) agg.Weather
}

type Location interface {
	BuildFromRaw(title string, latitude float64, longitude float64) agg.Location
	BuildFromDto(request apiDto.LocationRequest) agg.Location
}
