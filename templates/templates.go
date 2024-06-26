package templates

import (
	"bytes"
	"embed"
	"html/template"
	"io"
	"io/fs"

	"github.com/labstack/echo/v4"
	"github.com/paul-stern/admission-registry-web/model"
)

type Template struct {
	templates *template.Template
}

type WebPage struct {
	Title   string
	Content template.HTML
	Template
}

type Table struct {
	model.Entries
	Pages       int
	CurrentPage int
	Quantities  []int
	Template
}

const (
	templatesDir = "front"
)

var (
	//go:embed front/*
	files     embed.FS
	Templates map[string]*Template

	funcMap = map[string]template.FuncMap{
		"table.html": {
			"PageRange": Table.PageRange,
			"Count":     Table.Count,
		},
	}
)

func init() {
	LoadTemplates()
}

func LoadTemplates() error {
	if Templates == nil {
		Templates = make(map[string]*Template)
	}
	tmpFiles, err := fs.ReadDir(files, templatesDir)
	if err != nil {
		return err
	}

	for _, tmpl := range tmpFiles {
		if tmpl.IsDir() {
			continue
		}

		t := template.New(tmpl.Name())

		t.Funcs(funcMap[t.Name()])

		pt, err := t.ParseFS(files, templatesDir+"/"+tmpl.Name())
		if err != nil {
			return err
		}

		Templates[tmpl.Name()] = &Template{templates: pt}
	}
	return nil
}

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func GetTemplate(name string, c echo.Context) *Template {
	l := c.Logger()
	t, ok := Templates[name]
	if !ok {
		l.Errorf("Template %s not found", name)
	}
	return t
}

func (t Template) Name() string {
	return t.templates.Name()
}

func NewTable(e model.Entries, p int, cp int, c echo.Context) Table {
	return Table{
		Entries:     e,
		Pages:       p,
		CurrentPage: cp,
		Quantities:  []int{20, 50, 100},
		Template:    *GetTemplate("table.html", c),
	}
}

func (t Table) HTML(c echo.Context) template.HTML {
	b := new(bytes.Buffer)
	t.Render(b, t.Template.Name(), t, c)
	return template.HTML(b.String())
}

func (t Table) PageRange() []int {
	// Limit sets the amount pf page link shown on page
	limit := 18
	p := make([]int, limit+2)
	// for n := range t.Pages {

	// }
	p[0] = 1
	p[limit+1] = t.Pages
	mid := limit / 2
	switch {
	case t.Pages < limit || t.CurrentPage < mid:
		for n := 0; n < limit; n++ {
			p[n] = n + 1
		}
	case t.CurrentPage+mid < t.Pages:
		p[mid] = t.CurrentPage
		for n := mid; n >= 0; n-- {
			p[mid-n+1] = t.CurrentPage - n + 1
			p[mid+n-1] = t.CurrentPage + n - 1
		}
	case t.CurrentPage >= (t.Pages - limit):
		for n := 0; n < limit; n++ {
			p[n+1] = t.Pages - limit + n
		}
	}
	return p
}

func NewWebPage(title string, e model.Entries, p int, cp int, c echo.Context) WebPage {
	t := NewTable(e, p, cp, c)
	wp := WebPage{
		Title:    title,
		Template: *GetTemplate("base.html", c),
		Content:  t.HTML(c),
	}
	return wp
}

func (t Table) Count() int {
	return len(t.Entries)
}
