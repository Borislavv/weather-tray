package widget

const (
	Sochi     = "Сочи"
	Miass     = "Миасс"
	Krasnodar = "Краснодар"
)

type Widget struct {
	name string
}

func NewWidget() *Widget {
	return &Widget{
		name: "Weather-Widget",
	}
}

func (w *Widget) Run() error {
	return nil
}

func (w *Widget) GetName() string {
	return w.name
}
