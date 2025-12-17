package main

import (
	"log"

	"github.com/Lunarisnia/device-finder/internal/bot"
	"github.com/Lunarisnia/device-finder/internal/tinycli"
)

func main() {
	app := tinycli.New()
	app.SetProgram(bot.Run)
	err := app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
