package repository

import "github.com/Borislavv/weather-tray/internal/domain/agg"

/**
 * Repository must accept an aggregate and return it if it's possible.
 */

type Weather interface {
	Get(location agg.Location) agg.Weather
	Create(weather agg.Weather) (string, error)
}

type Location interface {
	Get() agg.Location
}
