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

type Context struct {
	context  echo.Context
	page     int
	quantity int
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
	wp := templates.NewWebPage("Журнал", es, lp, cp, c)
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
	if ps == "" && name == "page" {
		p = 1
	} else if ps == "" && name == "quantity" {
		p = 50
	} else {
		p, err = strconv.Atoi(ps)
	}
	return
}

func NewContext(c echo.Context) (nc Context) {
	nc.context = c
	return
}

// func (c *Context) Page (err error)

// TODO: Reimplement with regexps?
func (c *Context) QueryParams() (err error) {
	// validPage := regexp.MustCompile(`%d`)
	m := map[string]struct{}{
		"page" : {"d"}
	}
	ps := c.context.QueryParam("page")
	if ps == "" {
		c.page = 1
	}
	n, err := strconv.Atoi(ps)
	if err != nil {
		return
	}
	c.page = c.AbsInt(n)
	ps = c.context.QueryParam("quantity")
	if ps == "" {
		c.quantity = 50
	}
	n, err = strconv.Atoi(ps)
	c.quantity = c.AbsInt(n)
	return
}

func (c *Context) AbsInt(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}
