package controller

import (
	"goSiteProject/internal/monitoring"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

// Cache
var tmplMessageForm *template.Template

func InitTemplateMessageForm() {
	path := filepath.Join("template", "new_message_form.html")
	tmplMessageForm = template.Must(template.ParseFiles(path))
}

func NewMessageWall(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := tmplMessageForm.Execute(rw, nil)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}

	monitoring.IncrementEndpointHttpCounter("/new")
	monitoring.IncrementTotalhttpCounter()
}
