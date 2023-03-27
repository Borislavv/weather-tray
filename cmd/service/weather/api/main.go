package api

import (
	"fmt"
	"log"
)

func main() {
	api := NewApi()
	if err := api.Run(); err != nil {
		log.Fatalln(
			fmt.Sprintf("error occurred while %s service started", api.GetName()),
		)
	}
}
