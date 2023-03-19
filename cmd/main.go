package main

import (
	"encoding/json"
	"fmt"
	"github.com/getlantern/systray"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	Miass     = "Миасс"
	Sochi     = "Сочи"
	Krasnodar = "Краснодар"
)

type State struct {
	Temperature        float64
	WindSpeed          float64
	SelectedLocation   string
	LocationToUrl      map[string]string
	LocationToMenuItem map[string]*systray.MenuItem
	Cron               *cron.Cron
}

type WeatherResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		WindSpeed   float64 `json:"windspeed"`
	} `json:"current_weather"`
}

func main() {
	state := &State{
		SelectedLocation: Sochi,
		LocationToUrl: map[string]string{
			Miass:     "https://api.open-meteo.com/v1/forecast?latitude=55.0832254&longitude=60.0960609&current_weather=true&hourly=temperature_2m",
			Sochi:     "https://api.open-meteo.com/v1/forecast?latitude=43.6358199&longitude=39.7117688&current_weather=true&hourly=temperature_2m",
			Krasnodar: "https://api.open-meteo.com/v1/forecast?latitude=45.0250369&longitude=39.0414887&current_weather=true&hourly=temperature_2m",
		},
		LocationToMenuItem: map[string]*systray.MenuItem{},
	}

	systray.Run(state.onReady, state.onExit)
}

func (state *State) onReady() {
	state.updateWeather()

	state.Cron = cron.New()
	if err := state.Cron.AddFunc("@every 30s", state.updateWeather); err != nil {
		log.Fatalln(err)
	}
	state.Cron.Start()

	for location, _ := range state.LocationToUrl {
		state.LocationToMenuItem[location] = systray.AddMenuItem(location, "")
	}

	for {
		select {
		case <-state.LocationToMenuItem[Miass].ClickedCh:
			state.SelectedLocation = Miass
			state.updateWeather()
		case <-state.LocationToMenuItem[Sochi].ClickedCh:
			state.SelectedLocation = Sochi
			state.updateWeather()
		case <-state.LocationToMenuItem[Krasnodar].ClickedCh:
			state.SelectedLocation = Krasnodar
			state.updateWeather()
		}
	}
}

func (state *State) onExit() {

}

func (state *State) updateWeather() {
	client := &http.Client{Timeout: 10 * time.Second}
	url := state.LocationToUrl[state.SelectedLocation]
	res, err := client.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalln(fmt.Sprintf("received status is not OK: %d", res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	weatherResponse := WeatherResponse{}
	if err = json.Unmarshal(body, &weatherResponse); err != nil {
		log.Fatalln(err)
	}

	systray.SetTitle(
		fmt.Sprintf(
			"%s: %v°C, %vM/s",
			state.SelectedLocation,
			weatherResponse.CurrentWeather.Temperature,
			weatherResponse.CurrentWeather.WindSpeed,
		),
	)
}
