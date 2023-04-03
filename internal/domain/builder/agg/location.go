package aggBuilder

import (
	"github.com/Borislavv/weather-tray/internal/data/repository"
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	apiDto "github.com/Borislavv/weather-tray/internal/domain/dto/api"
	"github.com/Borislavv/weather-tray/internal/domain/entity"
)

type LocationAgg struct {
	locationRepository *repository.Location
}

func NewLocationBuilder() *LocationAgg {
	return &LocationAgg{}
}

func (b *LocationAgg) BuildFromRaw(title string, latitude float64, longitude float64) agg.Location {
	return agg.Location{
		Location: entity.Location{
			Title:     title,
			Latitude:  latitude,
			Longitude: longitude,
		},
	}
}

func (b *LocationAgg) BuildFromDto(request apiDto.LocationRequest) agg.Location {
	return agg.Location{
		Location: entity.Location{
			Title:     request.Title,
			Latitude:  request.Latitude,
			Longitude: request.Longitude,
		},
	}
}
