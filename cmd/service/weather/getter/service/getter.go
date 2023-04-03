package service

import (
	"fmt"
	gateway "github.com/Borislavv/weather-tray/internal/data/gateway/netw"
	"github.com/Borislavv/weather-tray/internal/data/repository/file"
	aggBuilder "github.com/Borislavv/weather-tray/internal/domain/builder/agg"
	queryBuilder "github.com/Borislavv/weather-tray/internal/domain/builder/query"
	getterDto "github.com/Borislavv/weather-tray/internal/domain/dto/getter"
	getterUsecase "github.com/Borislavv/weather-tray/internal/domain/usecase/getter"
	"time"
)

type Getter struct {
	name string
}

func NewGetter() *Getter {
	return &Getter{
		name: "Weather-Getter",
	}
}

func (g *Getter) Run() error {
	weatherGateway := gateway.NewWeatherGateway(
		"https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current_weather=true&hourly=temperature_2m",
		time.Duration(time.Second*10),
	)

	weatherAggBuilder := aggBuilder.NewWeatherBuilder()

	locationQueryBuilder := queryBuilder.NewLocationQueryBuilder()

	locationUsecase := getterUsecase.NewLocationUsecase(
		fileRepository.NewLocationRepository("locations.csv"),
		locationQueryBuilder,
	)

	weatherUsecase := getterUsecase.NewWeatherGetter(
		weatherGateway,
		weatherAggBuilder,
		locationUsecase,
	)

	weather, err := weatherUsecase.Get(
		getterDto.LocationRequest{Title: "Miass"},
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("%+v\n", weather)

	return nil
}

func (g *Getter) GetName() string {
	return g.name
}
