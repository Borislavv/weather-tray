package getter

import (
	"fmt"
	"log"
)

func main() {
	getter := NewGetter()
	if err := getter.Run(); err != nil {
		log.Fatalln(
			fmt.Sprintf("error occurred while %s service started", getter.GetName()),
		)
	}
}
