package builder

import (
	"github.com/Borislavv/weather-tray/internal/domain/agg/query"
	getterDto "github.com/Borislavv/weather-tray/internal/domain/dto/getter"
)

type LocationQuery interface {
	Build(request getterDto.LocationRequest) query.LocationQuery
}
