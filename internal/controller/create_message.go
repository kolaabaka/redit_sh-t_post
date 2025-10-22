package controller

import (
	"goSiteProject/internal/model"
	"goSiteProject/internal/service"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func CreateMessage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	message := r.FormValue("comment")
	name := r.FormValue("name")
	topic := r.FormValue("topic")

	t := time.Now()

	if topic != "" {
		service.AddMesaage(topic, model.Message{Name: name, Message: message, Date: t.Format("01-02-2006 15:04:05")})
	}

	http.Redirect(rw, r, "/", http.StatusMovedPermanently)
}
