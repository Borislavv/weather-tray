package main

import (
	"fmt"
	"github.com/Borislavv/weather-tray/cmd/service/weather/getter/service"
	"log"
)

func main() {
	getter := service.NewGetter()
	if err := getter.Run(); err != nil {
		log.Fatalln(
			fmt.Sprintf("error occurred while %s service started", getter.GetName()),
		)
	}
}
