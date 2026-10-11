package mailactions

import (
	_ "embed"
	"html/template"
	"net/http"
)

//go:embed templates/page.html.tmpl
var pageTmpHTML string

type page struct {
	Title string
	Body  string
}

var tmpl = template.Must(template.New("page.html.tmpll").Parse(pageTmpHTML))

func renderPage(w http.ResponseWriter, status int, p page) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.WriteHeader(status)
	_ = tmpl.Execute(w, p)
}
