package controller

import (
	"goSiteProject/service"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func MessageWall(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path := filepath.Join("template", "index.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}

	err = tmpl.ExecuteTemplate(rw, "message", service.GetMesaages())
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}
}
