package query

type Weather struct {
	Id          string  `json:"id" schema:"id"`
	LocationId  string  `json:"locationId" schema:"locationId"`
	Temperature float64 `json:"temperature" schema:"temperature"`
	WindSpeed   float64 `json:"windSpeed" schema:"windSpeed"`
}
