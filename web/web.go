package web

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/paul-stern/admission-registry-web/model"
	"github.com/paul-stern/admission-registry-web/templates"
)

type Params map[string]string

type Journal struct {
	Page    int
	Entries model.Entries
}

var storage model.Entries

func init() {
	storage = model.GenEntries(1000)
}

func ShowJournal(c echo.Context) error {
	l := c.Logger()
	tt := templates.GetTemplates()
	l.Print(tt)
	t := tt["base.html"]
	c.Echo().Renderer = t
	l.Print("Logger started")
	j := Journal{}
	// Get p query param and convert to it to int
	p, err := strconv.Atoi(c.QueryParam("p"))
	if err != nil {
		// log.Print()
		l.Printf("strconv err: %s", err)
	}
	j.Page = p

	// return c.String(http.StatusOK, p)
	d := templates.Page{
		Title:   "Test",
		Content: "Hello!",
	}
	return c.Render(http.StatusOK, "base.html", d)
}
