package web

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/paul-stern/admission-registry-web/model"
	"github.com/paul-stern/admission-registry-web/templates"
)

var entries model.Entries

func init() {
	entries = model.GenEntries(1000)
}

func ShowJournal(c echo.Context) error {
	// l := c.Logger()
	// l.Print("Logger started")
	es, err := Page(c, entries, 100)
	if err != nil {
		return err
	}
	wp := templates.NewWebPage("Журнал", es, c)
	c.Echo().Renderer = &wp.Template
	return c.Render(http.StatusOK, wp.Name(), wp)
}

func Page(c echo.Context, es model.Entries, n int) (model.Entries, error) {
	var err error
	var p int
	ps := c.QueryParam("p")
	if ps == "" {
		p = 1
	} else {
		p, err = strconv.Atoi(ps)
		if err != nil {
			return model.Entries{}, err
		}

	}
	// Define start and end of the Entries slice
	s := (p - 1) * n
	e := s + 100
	return es[s:e], err
}

func LastPage(c echo.Context, es model.Entries, n int) int {
	return len(es) % n
}
