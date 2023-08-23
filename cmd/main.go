package main

import (
	"log"

	"github.com/eXoterr/StrimaNET_Backend/internal/app"
)

func main() {
	app, e := app.New()
	if e != nil {
		log.Fatalln(e)
	}

	e = app.Run()
	if e != nil {
		log.Fatalln(e)
	}
}
