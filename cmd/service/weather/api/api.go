package api

type Api struct {
	name string
}

func NewApi() *Api {
	return &Api{
		name: "Weather-Api",
	}
}

func (api *Api) Run() error {
	return nil
}

func (api *Api) GetName() string {
	return api.name
}
