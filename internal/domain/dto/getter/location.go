package getterDto

// TODO add opts to queryRequest
type LocationRequest struct {
	Id        string  `json:"id"`
	Title     string  `json:"title"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
