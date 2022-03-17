package main

import (
	"weather-forecast/routes"
)

func main() {
	e := routes.Init()
	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Serve Static File
	e.File("", "views/build/index.html")
	e.Static("/static", "views/build/static")
	e.File("/favicon.ico", "views/build/favicon.ico")

	e.Logger.Fatal(e.Start(":8000"))
}
