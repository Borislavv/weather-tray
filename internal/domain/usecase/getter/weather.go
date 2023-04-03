package getterUsecase

import (
	"fmt"
	"github.com/Borislavv/weather-tray/internal/data/gateway"
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	"github.com/Borislavv/weather-tray/internal/domain/builder"
	"github.com/Borislavv/weather-tray/internal/domain/dto/getter"
	"github.com/Borislavv/weather-tray/internal/domain/usecase"
)

type WeatherGetter struct {
	weatherGateway    gateway.Weather
	weatherAggBuilder builder.WeatherAgg
	locationUsecase   usecase.Location
}

func NewWeatherGetter(
	weatherGateway gateway.Weather,
	weatherAggBuilder builder.WeatherAgg,
	locationUsecase usecase.Location,
) *WeatherGetter {
	return &WeatherGetter{
		weatherGateway:    weatherGateway,
		weatherAggBuilder: weatherAggBuilder,
		locationUsecase:   locationUsecase,
	}
}

func (w *WeatherGetter) Get(request getterDto.LocationRequest) (agg.Weather, error) {
	location, err := w.locationUsecase.Find(request)
	if err != nil {
		return agg.Weather{}, err
	}

	resp, err := w.weatherGateway.Get(location)
	if err != nil {
		return agg.Weather{}, err
	}

	fmt.Printf("ApiResponse: %+v\n", resp)

	weather := w.weatherAggBuilder.BuildFromResponseDto(resp, location)

	return weather, nil
}
