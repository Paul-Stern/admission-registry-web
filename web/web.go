package web

import (
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/paul-stern/admission-registry-web/config"
	"github.com/paul-stern/admission-registry-web/model"
	"github.com/paul-stern/admission-registry-web/templates"
)

var entries model.Entries

type Body struct {
	io.Reader
}

func init() {
	entries = model.GenEntries(1000)
}

func SignUp(c echo.Context) error {
	l := config.Node("hello")
	// http.Get(config.Conf)
	r, err := http.Get(l)
	if err != nil {
		return err
	}
	b := Body{r.Body}
	// return c.JSON(http.StatusOK, r.Body)
	return c.String(http.StatusOK, b.String())
}

func ShowJournal(c echo.Context) error {
	// l := c.Logger()
	// l.Print("Logger started")
	lp, err := LastPage(c, entries)
	if err != nil {
		return err
	}
	es, err := Page(c, entries)
	if err != nil {
		return err
	}
	cp, err := ParseParam(c, "page")
	if err != nil {
		return err
	}
	tbl := templates.NewTable(es, lp, cp, c)
	wp := templates.NewWebPage("Журнал", tbl.HTML(c), c)
	c.Echo().Renderer = &wp.Template
	return c.Render(http.StatusOK, wp.Name(), wp)
}

// Page returns Entries for one page of journal and an error
// It takes Context, Entries and n number of entries per page
func Page(c echo.Context, es model.Entries) (model.Entries, error) {
	var err error
	var p int
	p, err = ParseParam(c, "page")
	if err != nil {
		return model.Entries{}, err
	}
	n, err := ParseParam(c, "quantity")
	if err != nil {
		return model.Entries{}, err
	}

	// Define start and end of the Entries slice
	s := (p - 1) * n
	e := s + n
	return es[s:e], err
}

func LastPage(c echo.Context, es model.Entries) (int, error) {
	q, err := ParseParam(c, "quantity")
	return len(es) / q, err
}

// ParseParam parses integer named parameter
func ParseParam(c echo.Context, name string) (p int, err error) {
	ps := c.QueryParam(name)
	if (ps == "" || ps == "0") && name == "page" {
		p = 1
	} else if ps == "" && name == "quantity" {
		p = 50
	} else {
		p, err = strconv.Atoi(ps)
	}
	return
}

func (body Body) String() string {
	b, err := io.ReadAll(body)
	if err != nil {
		return ""
	}
	return string(b)
}
