package query

type LocationQuery struct {
	Id        string  `json:"id" schema:"id"`
	Title     string  `json:"title" schema:"title"`
	Latitude  float64 `json:"latitude" schema:"latitude"`
	Longitude float64 `json:"longitude" schema:"longitude"`
}
