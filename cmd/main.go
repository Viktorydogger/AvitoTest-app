package main

import (
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(avitotestapp.Server)
	if err := srv.Run("0000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
