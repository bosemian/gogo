package view

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
)

var (
	tpIndex      = parseTemplate("root.tmpl", "index.tmpl")
	tpAdminLogin = parseTemplate("root.tmpl", "admin/login.tmpl")
	tpAdminList  = parseTemplate("root.tmpl", "admin/list.tmpl")
)

var m = minify.New()

const templateDir = "template"

func init() {
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/javascript", js.Minify)
}

func joinTemplate(files ...string) []string {
	r := make([]string, len(files))
	for i, f := range files {
		r[i] = filepath.Join(templateDir, f)
	}
	return r
}

func parseTemplate(files ...string) *template.Template {
	t := template.New("")
	t.Funcs(template.FuncMap{})
	_, err := t.ParseFiles(joinTemplate(files...)...)
	if err != nil {
		panic(err)
	}
	t = t.Lookup("root")
	return t
}

// Index render index view
func Index(w http.ResponseWriter, data interface{}) {
	render(tpIndex, w, data)
}

// AdminLogin render admin view
func AdminLogin(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}

// AdminList render adminlist view
func AdminList(w http.ResponseWriter, data interface{}) {
	render(tpAdminList, w, data)
}

func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf := bytes.Buffer{}
	err := t.Execute(&buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m.Minify("text/html", w, &buf)
}