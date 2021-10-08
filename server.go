package main

import (
	"code-echo/db"
	"code-echo/route"
)

func main() {
	db.Init()

	e := route.Init()

	e.Logger.Fatal(e.Start(":1324"))

}
