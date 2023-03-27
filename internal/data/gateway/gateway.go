package gateway

import "github.com/Borislavv/weather-tray/internal/domain/agg"

type Weather interface {
	Get(location agg.Location)
}
