package main

import (
	di "github.com/io-m/clean/internal"
)

func main() {
	port := 9000
	app := di.AppSetup(port)
	app.GogoBaby()
}
