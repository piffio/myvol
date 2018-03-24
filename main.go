package main

import (
	"log"

	"github.com/piffio/myvol/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
