package main

import (
	"encoding/json"
	"fmt"
	"github.com/getlantern/systray"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	state := &State{}

	systray.Run(state.onReady, state.onExit)
}

func (state *State) onReady() {
	state.updateWeather()

	state.Cron = cron.New()
	if err := state.Cron.AddFunc("@every 30s", state.updateWeather); err != nil {
		log.Fatalln(err)
	}
	state.Cron.Start()
}

func (state *State) onExit() {

}

type State struct {
	Temperature float64 `json:"temperature"`
	Windspeed   float64 `json:"windspeed"`
	Cron        *cron.Cron
}

type WeatherResponse struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	CurrentWeather       struct {
		Temperature   float64 `json:"temperature"`
		Windspeed     float64 `json:"windspeed"`
		Winddirection float64 `json:"winddirection"`
		Weathercode   int     `json:"weathercode"`
		Time          string  `json:"time"`
	} `json:"current_weather"`
	HourlyUnits struct {
		Time          string `json:"time"`
		Temperature2M string `json:"temperature_2m"`
	} `json:"hourly_units"`
	Hourly struct {
		Time          []string  `json:"time"`
		Temperature2M []float64 `json:"temperature_2m"`
	} `json:"hourly"`
}

func (state *State) updateWeather() {
	res, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=43.6358199&longitude=39.7117688&current_weather=true&hourly=temperature_2m")
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
			"Weather: %v°C, Wind: %vM/s",
			weatherResponse.CurrentWeather.Temperature,
			weatherResponse.CurrentWeather.Windspeed,
		),
	)
}
