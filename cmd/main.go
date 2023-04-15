package main

import (
	app "github.com/io-m/clean/internal"
)

func main() {
	port := 9000
	app := app.AppSetup(port)
	app.GogoBaby()
}
