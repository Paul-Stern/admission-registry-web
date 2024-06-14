package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/paul-stern/admission-registry-web/model"
)

type Params map[string]string

var storage model.Entries

func init() {
	storage = model.GenEntries(1000)
}

func Journal(w http.ResponseWriter, r *http.Request) {
	q := r.URL.RawQuery
	// templates.RenderTable(w, storage[:100])
	// fmt.Fprintf(w, "query: %+v", ParseQuery(q))
	pars := ParseQuery(q)
	p := pars["p"]
	fmt.Fprintf(w, "page: %s", p)

}

func ParseQuery(q string) (pars Params) {
	pars = make(Params)
	qslice := strings.Split(q, "&")
	for _, p := range qslice {
		s := strings.Split(p, "=")
		pars[s[0]] = s[1]
	}
	return
}
