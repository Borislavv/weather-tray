package entity

type Location struct {
	Title     string  `schema:"title"`
	Latitude  float64 `schema:"latitude"`
	Longitude float64 `schema:"longitude"`
}
