package main

import "github.com/merge-hotel-data/app"

func main() {
	app := app.NewApp()
	app.Init()
	app.Run()
}
