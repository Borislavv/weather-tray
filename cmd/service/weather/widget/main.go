package widget

import (
	"fmt"
	"log"
)

func main() {
	widget := NewWidget()
	if err := widget.Run(); err != nil {
		log.Fatalln(
			fmt.Sprintf("error occurred while %s service started", widget.GetName()),
		)
	}
}
