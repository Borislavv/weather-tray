package builder

import (
	"github.com/Borislavv/weather-tray/internal/data/repository"
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	apiRequestDto "github.com/Borislavv/weather-tray/internal/domain/dto/api"
	"github.com/Borislavv/weather-tray/internal/domain/entity"
)

type LocationBuilder struct {
	locationRepository *repository.Location
}

func NewLocationBuilder() *LocationBuilder {
	return &LocationBuilder{}
}

func (b *LocationBuilder) BuildFromRaw(title string, latitude float64, longitude float64) agg.Location {
	return agg.Location{
		Location: entity.Location{
			Title:     title,
			Latitude:  latitude,
			Longitude: longitude,
		},
	}
}

func (b *LocationBuilder) BuildFromDto(request apiRequestDto.LocationRequest) agg.Location {
	return agg.Location{
		Location: entity.Location{
			Title:     request.Title,
			Latitude:  request.Latitude,
			Longitude: request.Longitude,
		},
	}
}
