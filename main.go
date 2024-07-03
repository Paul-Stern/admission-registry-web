package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/paul-stern/admission-registry-web/config"
	"github.com/paul-stern/admission-registry-web/web"
)

func main() {
	e := echo.New()
	e.Debug = true
	c, err := config.Read()
	if err != nil {
		e.Logger.Fatal("Config error: ", err)
	}
	e.GET("/", web.ShowJournal)
	e.GET("/signup", web.SignUp)
	e.Logger.Fatal(e.StartTLS(
		fmt.Sprintf("%s:%s", c.Server.Addr, c.Server.Port),
		c.Server.Cert,
		c.Server.Key,
	))
}
