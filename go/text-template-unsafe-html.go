package main

import (
	"bytes"
	"embed"
	"fmt"
	"net/http"
	"text/template"
)

var embedded embed.FS

// Top-level const with raw HTML
const (
	notificationTemplate = `<h1>New CI failures for {{.Project}}</h1>`
)

func directXSS(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	// ruleid: text-template-unsafe-html
	tmpl, _ := template.New("greet").Parse("<h1>Hello " + name + "!</h1>")
	tmpl.Execute(w, nil)
}

func htmlAttrXSS(w http.ResponseWriter, r *http.Request) {
	color := r.FormValue("color")
	// ruleid: text-template-unsafe-html
	tmpl, _ := template.New("color").Parse(`<div style="color:` + color + `">Colored text</div>`)
	tmpl.Execute(w, nil)
}

func fmtSprintfXSS(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("msg")
	// ruleid: text-template-unsafe-html
	tmpl := template.Must(template.New("msg").Parse(fmt.Sprintf("<p>%s</p>", msg)))
	tmpl.Execute(w, nil)
}

func constOutsideXSS(w http.ResponseWriter, r *http.Request) {
	project := r.URL.Query().Get("project")
	// ruleid: text-template-unsafe-html
	t := template.Must(template.New("notify").Parse(notificationTemplate))
	t.Execute(w, map[string]string{"Project": project})
}

func fileIncludeXSS(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	// ruleid: text-template-unsafe-html
	template.ParseFS(embedded, page+".html")
}

func nestedXSS(w http.ResponseWriter, r *http.Request) {
	comment := r.FormValue("comment")
	// ruleid: text-template-unsafe-html
	template.New("c").Funcs(template.FuncMap{}).
		Parse("<span>" + comment + "</span>")
}

func safeFileExt(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	// ok: text-template-unsafe-html
	template.ParseFS(embedded, page+".tmpl")
}

func main() {
	http.HandleFunc("/direct", directXSS)
	http.HandleFunc("/attr", htmlAttrXSS)
	http.HandleFunc("/sprintf", fmtSprintfXSS)
	http.HandleFunc("/const", constOutsideXSS)
	http.HandleFunc("/file", fileIncludeXSS)
	http.HandleFunc("/nested", nestedXSS)
	http.HandleFunc("/safe-ext", safeFileExt)
	http.ListenAndServe(":8080", nil)
}
