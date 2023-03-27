package getter

type Getter struct {
	name string
}

func NewGetter() *Getter {
	return &Getter{
		name: "Weather-Getter",
	}
}

func (g *Getter) Run() error {
	return nil
}

func (g *Getter) GetName() string {
	return g.name
}
