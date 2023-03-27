package gateway

import (
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	dto "github.com/Borislavv/weather-tray/internal/domain/dto/getter/response"
)

/**
 * Gateway must accept an aggregate and return a DTO if it's possible.
 */

type Weather interface {
	Get(location agg.Location) (dto.WeatherResponse, error)
}
