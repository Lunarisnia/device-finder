package main

import (
	"log"

	"github.com/Lunarisnia/device-finder/internal/finder"
	"github.com/Lunarisnia/device-finder/internal/tinycli"
)

func main() {
	app := tinycli.New()
	app.SetProgram(finder.Run)
	err := app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
