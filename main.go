package main

import (
	"myapp/db"
	"myapp/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1309"))
}
