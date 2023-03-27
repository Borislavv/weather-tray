package entity

type Location struct {
	Name      string  `schema:"name"`
	Latitude  float64 `schema:"latitude"`
	Longitude float64 `schema:"longitude"`
}
