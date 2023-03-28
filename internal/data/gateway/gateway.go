package gateway

import (
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	"github.com/Borislavv/weather-tray/internal/domain/dto/getter"
)

/**
 * Gateway must accept an aggregate and return a DTO if it's possible.
 */

type Weather interface {
	Get(location agg.Location) (*getter.WeatherResponse, error)
}
