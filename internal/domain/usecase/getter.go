package usecase

import (
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	getterDto "github.com/Borislavv/weather-tray/internal/domain/dto/getter"
)

type Location interface {
	Find(request getterDto.LocationRequest) (agg.Location, error)
}
