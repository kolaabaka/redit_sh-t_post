package controller

import (
	"goSiteProject/internal/model"
	"goSiteProject/internal/monitoring"
	"goSiteProject/internal/service"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

// Cache
var tmplMessageWall *template.Template

func InitTemplateMessageWall() {
	path := filepath.Join("template", "index.html")
	tmplMessageWall = template.Must(template.ParseFiles(path))
}

func MessageWall(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	topic := r.URL.Query().Get("topic")

	if topic == "" {
		topic = "main_table"
	}

	messageList, err := service.GetMesaages(topic)

	//Topic is not exists
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = tmplMessageWall.ExecuteTemplate(rw, "message", model.MessageWallTempalte{Header: topic, MessageList: messageList})
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}
	monitoring.IncrementEndpointHttpCounter("/")
	monitoring.IncrementTotalhttpCounter()
}
