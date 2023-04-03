package aggBuilder

import (
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	getterResponseDto "github.com/Borislavv/weather-tray/internal/domain/dto/getter"
	"github.com/Borislavv/weather-tray/internal/domain/entity"
	"github.com/Borislavv/weather-tray/internal/domain/vo"
	"time"
)

type WeatherAgg struct {
}

func NewWeatherBuilder() *WeatherAgg {
	return &WeatherAgg{}
}

func (b *WeatherAgg) BuildFromResponseDto(
	response getterResponseDto.WeatherResponse,
	location agg.Location,
) agg.Weather {
	return agg.Weather{
		Location: location.Location,
		Weather: entity.Weather{
			Temperature: response.CurrentWeather.Temperature,
			WindSpeed:   response.CurrentWeather.WindSpeed,
		},
		Timestamp: vo.Timestamp{CreatedAt: time.Now()},
	}
}
