package controllers

import (
	"net/http"
	"path"
	"html/template"
)


func Index(w http.ResponseWriter, r *http.Request)  {
	t := template.New("index.tpl")
	t, _ = template.ParseFiles(path.Join("templates", "index.tpl"))
	t.ExecuteTemplate(w, "index.tpl", nil)
}
