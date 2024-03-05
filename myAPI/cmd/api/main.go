package main

import (
	"log"
	"myAPI/pkg/api"
)

func main() {
	app := api.Default()
	err := app.Start()
	if err != nil {
		log.Print(err)
		panic(err)
	}
}
