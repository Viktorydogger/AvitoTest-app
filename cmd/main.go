package main

import (
	avitotestapp "forgo/AvitoTest-app"
	"log"
)

func main() {
	srv := new(avitotestapp.Server)
	if err := srv.Run("0000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
