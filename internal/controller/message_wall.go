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

func MessageWall(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path := filepath.Join("template", "index.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}

	topic := r.URL.Query().Get("topic")

	if topic == "" {
		topic = "main_table"
	}

	messageList, err := service.GetMesaages(topic)

	//Topic is not exists
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = tmpl.ExecuteTemplate(rw, "message", model.MessageWallTempalte{Header: topic, MessageList: messageList})
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}
	monitoring.IncrementEndpointHttpCounter("/")
	monitoring.IncrementTotalhttpCounter()
}
