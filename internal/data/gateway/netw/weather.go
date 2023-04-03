package gateway

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	"github.com/Borislavv/weather-tray/internal/domain/dto/getter"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type WeatherNetwGateway struct {
	endpoint string // pattern (latitude=%s&longitude=%s)
	timeout  time.Duration
}

func NewWeatherGateway(
	endpoint string,
	timeout time.Duration,
) *WeatherNetwGateway {
	return &WeatherNetwGateway{
		endpoint: endpoint,
		timeout:  timeout,
	}
}

// Get is a method which return a temperature and windSpeed by given Location.
func (w *WeatherNetwGateway) Get(location agg.Location) (getterDto.WeatherResponse, error) {
	client := &http.Client{Timeout: w.timeout}
	url := fmt.Sprintf(
		w.endpoint,
		strconv.FormatFloat(location.Latitude, 'f', -1, 64),
		strconv.FormatFloat(location.Longitude, 'f', -1, 64),
	)

	fmt.Printf("%+v\n", url)

	res, err := client.Get(url)
	if err != nil {
		return getterDto.WeatherResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return getterDto.WeatherResponse{}, errors.New(fmt.Sprintf("received status is not OK: %d", res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return getterDto.WeatherResponse{}, err
	}

	response := &getterDto.WeatherResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return getterDto.WeatherResponse{}, err
	}

	return *response, nil
}
