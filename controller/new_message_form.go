package controller

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func NewMessageWall(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path := filepath.Join("public", "html", "new_message_form.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}

	err = tmpl.Execute(rw, nil)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}
}
