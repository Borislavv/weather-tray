package netw

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	"github.com/Borislavv/weather-tray/internal/domain/dto"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Weather struct {
	endpoint string
	timeout  time.Duration
}

func NewWeather(
	endpoint string,
	timeout time.Duration,
) *Weather {
	return &Weather{
		endpoint: endpoint,
		timeout:  timeout,
	}
}

func (w *Weather) Get(location agg.Location) (*dto.WeatherResponse, error) {
	client := &http.Client{Timeout: w.timeout}
	url := fmt.Sprintf(w.endpoint, location.Latitude, location.Longitude)

	res, err := client.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("received status is not OK: %d", res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &dto.WeatherResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
