package queryBuilder

import (
	"github.com/Borislavv/weather-tray/internal/domain/agg/query"
	getterDto "github.com/Borislavv/weather-tray/internal/domain/dto/getter"
)

type LocationQueryBuilder struct {
}

func NewLocationQueryBuilder() *LocationQueryBuilder {
	return &LocationQueryBuilder{}
}

func (b *LocationQueryBuilder) Build(request getterDto.LocationRequest) query.LocationQuery {
	q := query.LocationQuery{}

	if request.Id != "" {
		q.Id = request.Id
	}

	if request.Title != "" {
		q.Title = request.Title
	}

	if request.Latitude != 0 && request.Longitude != 0 {
		q.Latitude = request.Latitude
		q.Longitude = request.Longitude
	}

	return q
}
