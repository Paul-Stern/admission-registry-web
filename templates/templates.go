package templates

import (
	"bytes"
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Page struct {
	Title   string
	Content template.HTML
}

const (
	templatesDir = "front"
)

var (
	//go:embed front/*
	files     embed.FS
	Templates map[string]*template.Template
)

func init() {
	LoadTemplates()
}

func LoadTemplates() error {
	if Templates == nil {
		Templates = make(map[string]*template.Template)
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

		Templates[tmpl.Name()] = pt
	}
	return nil
}

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

func RenderTable(w http.ResponseWriter) {
	t, ok := Templates["table.tmpl"]
	if !ok {
		log.Printf("template %s not found", "table.tmpl")
		return
	}
	b := new(bytes.Buffer)
	if err := t.Execute(b, nil); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	p := Page{
		Title:   "Журнал",
		Content: template.HTML(b.String()),
	}

	RenderTemplate(w, "base", p)
}

func GetTmplts() map[string]*template.Template {
	return Templates
}
