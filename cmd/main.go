package main

import "github.com/eXoterr/StrimaNET_Backend/internal/app"

func main() {
	app, e := app.New()
	if e != nil {
		panic(e)
	}

	e = app.Run()
	if e != nil {
		panic(e)
	}
}
