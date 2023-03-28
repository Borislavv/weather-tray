package gateway

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	"github.com/Borislavv/weather-tray/internal/domain/dto/getter"
	"io/ioutil"
	"net/http"
	"time"
)

type WeatherGateway struct {
	endpoint string // pattern (latitude=%s&longitude=%s)
	timeout  time.Duration
}

func NewWeather(
	endpoint string,
	timeout time.Duration,
) *WeatherGateway {
	return &WeatherGateway{
		endpoint: endpoint,
		timeout:  timeout,
	}
}

// Get is a method which return a temperature and windSpeed by given Location.
func (w *WeatherGateway) Get(location agg.Location) (*getter.WeatherResponse, error) {
	client := &http.Client{Timeout: w.timeout}
	url := fmt.Sprintf(w.endpoint, location.Latitude, location.Longitude)

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("received status is not OK: %d", res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &getter.WeatherResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
