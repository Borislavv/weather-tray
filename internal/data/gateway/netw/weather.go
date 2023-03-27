package netw

import (
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	"github.com/Borislavv/weather-tray/internal/domain/entity"
	"time"
)

type Weather struct {
	endpoint string
	timeout  time.Duration
}

func NewWeather(
	endpoint string,
	timeout time.Duration,
) *Weather {
	return &Weather{
		endpoint: endpoint,
		timeout:  timeout,
	}
}

// todo Подумать, возвращать аггрегат или же dto
func (w *Weather) Get(location agg.Location) agg.Weather {
	return agg.Weather{
		Weather: entity.Weather{},
	}
}
