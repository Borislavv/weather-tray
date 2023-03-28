package entity

type Weather struct {
	Temperature float64 `schema:"temperature"`
	WindSpeed   float64 `schema:"windSpeed"`
}
