package main

import (
	"xsis/test/config"
	"xsis/test/factory"
	"xsis/test/middlewares"
	"xsis/test/routes"
)

func main() {
	dbConn := config.ConnectionDB()

	presenter := factory.InitFactory(dbConn)
	e := routes.New(presenter)
	middlewares.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":80"))
}
