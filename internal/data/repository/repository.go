package repository

import (
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	"github.com/Borislavv/weather-tray/internal/domain/agg/query"
)

/**
 * Repository must accept an aggregate and return it if it's possible.
 */

type Weather interface {
	// change to use locationQuery ass parameter
	Get(location agg.Location) agg.Weather
	Create(weather agg.Weather) (string, error)
}

type Location interface {
	Find(query query.LocationQuery) (agg.Location, error)
}
