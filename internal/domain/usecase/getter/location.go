package getterUsecase

import (
	"github.com/Borislavv/weather-tray/internal/data/repository"
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	"github.com/Borislavv/weather-tray/internal/domain/builder"
	getterDto "github.com/Borislavv/weather-tray/internal/domain/dto/getter"
)

type LocationUsecase struct {
	locationRepository   repository.Location
	locationQueryBuilder builder.LocationQuery
}

func NewLocationUsecase(
	locationRepository repository.Location,
	locationQueryBuilder builder.LocationQuery,
) *LocationUsecase {
	return &LocationUsecase{
		locationRepository:   locationRepository,
		locationQueryBuilder: locationQueryBuilder,
	}
}

func (l *LocationUsecase) Find(request getterDto.LocationRequest) (agg.Location, error) {
	location, err := l.locationRepository.Find(
		l.locationQueryBuilder.Build(request),
	)
	if err != nil {
		return agg.Location{}, err
	}
	return location, nil
}
