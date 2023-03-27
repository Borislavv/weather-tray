package agg

import (
	"github.com/Borislavv/weather-tray/internal/domain/entity"
	"github.com/Borislavv/weather-tray/internal/domain/vo"
)

type Weather struct {
	// Entities
	entity.Location
	entity.Weather
	// ValueObjects
	vo.Timestamp
}
