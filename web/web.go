package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/paul-stern/admission-registry-web/model"
	"github.com/paul-stern/admission-registry-web/templates"
)

type Params map[string]string

func ShowJournal(c echo.Context) error {
	l := c.Logger()
	l.Print("Logger started")
	wp := templates.NewWebPage("Журнал", model.GenEntries(100), c)
	c.Echo().Renderer = &wp.Template
	return c.Render(http.StatusOK, wp.Name(), wp)
}
