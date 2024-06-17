package main

import (
	"github.com/labstack/echo/v4"
	"github.com/paul-stern/admission-registry-web/web"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.GET("/", web.ShowJournal)
	e.Logger.Fatal(e.Start(":8083"))
}
