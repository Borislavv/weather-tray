package file

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/Borislavv/weather-tray/internal/domain/agg"
	"github.com/Borislavv/weather-tray/internal/domain/agg/query"
	"github.com/Borislavv/weather-tray/internal/domain/entity"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

type LocationRepository struct {
	filepath string
	data     []agg.Location
}

func NewLocationRepository(filepath string) *LocationRepository {
	return &LocationRepository{filepath: filepath}
}

func (r *LocationRepository) Find(query query.LocationQuery) (agg.Location, error) {
	storage, err := r.buildDataSlice()
	if err != nil {
		return agg.Location{}, err
	}

	fmt.Println(storage)

	for _, location := range storage {
		if query.Title != "" {
			if query.Title == location.Title {
				return location, nil
			}
		} else if query.Latitude != 0 && query.Longitude != 0 {
			if query.Latitude == location.Latitude && query.Longitude == location.Longitude {
				return location, nil
			}
		}
	}

	return agg.Location{}, errors.New(
		fmt.Sprintf("location is not found by given query: %+v", query),
	)
}

func (r *LocationRepository) buildDataSlice() ([]agg.Location, error) {
	var storage []agg.Location

	path, err := r.getFullFilepath()
	if err != nil {
		return storage, err
	}

	f, err := os.Open(path)
	if err != nil {
		return storage, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return storage, err
	}

	for l, record := range records {
		if l == 0 {
			// skip csv header
			continue
		}

		location := agg.Location{
			Location: entity.Location(
				struct {
					Title     string
					Latitude  float64
					Longitude float64
				}{Title: record[0], Latitude: 0, Longitude: 0},
			),
		}

		if latitude, err := strconv.ParseFloat(record[1], 64); err == nil {
			location.Latitude = latitude
		} else {
			return storage, err
		}

		if longitude, err := strconv.ParseFloat(record[2], 64); err == nil {
			location.Longitude = longitude
		} else {
			return storage, err
		}

		storage = append(storage, location)
	}

	return storage, nil
}

func (r *LocationRepository) getFullFilepath() (string, error) {
	dir, err := r.getDir()
	if err != nil {
		return "", err
	}
	return dir + "/" + r.filepath, nil
}

func (r *LocationRepository) getDir() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filepath.Dir(filename), nil
}
