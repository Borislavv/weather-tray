package entity

type Weather struct {
	LocationId  string  `schema:"locationId"`
	Temperature float64 `schema:"temperature"`
	WindSpeed   float64 `schema:"windSpeed"`
}
