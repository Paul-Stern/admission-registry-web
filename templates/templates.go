package templates

import (
	"embed"
	"html/template"
	"io"
	"io/fs"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

type Page struct {
	Title   string
	Content template.HTML
}

const (
	templatesDir = "front"
)

var (
	//go:embed front/*
	files embed.FS
	// Templates map[string]*template.Template
	Templates map[string]*Template
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

		pt, err := t.ParseFS(files, templatesDir+"/"+tmpl.Name())
		if err != nil {
			return err
		}

		// Templates[tmpl.Name()] = pt
		Templates[tmpl.Name()] = &Template{templates: pt}
	}
	return nil
}

/*
	func RenderTemplate(w http.ResponseWriter, tmpl string, data any) {
		t, ok := Templates[tmpl+".html"]
		if !ok {
			log.Printf("template %s not found", tmpl+".html")
			return
		}

		if err := t.Execute(w, data); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	func RenderTable(w http.ResponseWriter, es model.Entries) {
		t, ok := Templates["table.html"]
		if !ok {
			log.Printf("template %s not found", "table.html")
			return
		}
		b := new(bytes.Buffer)
		// ents := model.GenEntries(100)
		if err := t.Execute(b, es); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		p := Page{
			Title:   "Журнал",
			Content: template.HTML(b.String()),
		}

		RenderTemplate(w, "base", p)
	}
*/
func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
	// t, ok := Templates[tmpl+".html"]
	// if !ok {
	// 	log.Printf("template %s not found", tmpl+".html")
	// 	return
	// }

	// if err := t.Execute(w, data); err != nil {
	// 	log.Println(err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

}

func GetTemplates() map[string]*Template {
	return Templates
}
