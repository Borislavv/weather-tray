package getterUsecase

import (
	"github.com/Borislavv/weather-tray/internal/data/gateway"
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	"github.com/Borislavv/weather-tray/internal/domain/builder"
	getterDto "github.com/Borislavv/weather-tray/internal/domain/dto/getter"
)

type WeatherGetter struct {
	weatherGateway *gateway.Weather
	weatherBuilder *builder.Weather
}

func NewWeatherGetter(gateway *gateway.Weather, builder *builder.Weather) *WeatherGetter {
	return &WeatherGetter{
		weatherGateway: gateway,
		weatherBuilder: builder,
	}
}

func (w *WeatherGetter) Get(location getterDto.LocationRequest) (agg.Weather, error) {

}
