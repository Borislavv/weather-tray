package query

type LocationQuery struct {
	// Filter by `ID`
	//
	// required: false
	// example: `62565ef869ece65cdec4a43c`
	Id string `json:"id" schema:"id"`
	// Filter by `Title`
	//
	// required: false
	// example: `Sochi`
	Title string `json:"title" schema:"title"`
	// Filter by `Latitude`
	//
	// required: false
	// example: `56.864568`
	Latitude float64 `json:"latitude" schema:"latitude"`
	// Filter by `Longitude`
	//
	// required: false
	// example: `44.895674`
	Longitude float64 `json:"longitude" schema:"longitude"`
	// Query options
	//
	// required: false
	// example:
	//	- SortBy: "title"
	//  - OrderBy: "title"
	//  - Offset: 100
	//  - Limit: 1000
	Options OptsQuery `json:"options" schema:"options"`
}

func (q LocationQuery) GetOpts() OptsQuery {
	return q.Options
}
