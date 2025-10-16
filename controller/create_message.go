package controller

import (
	"goSiteProject/model"
	"goSiteProject/service"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func CreateMessage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	message := r.FormValue("message")
	name := r.FormValue("name")

	t := time.Now()

	service.AddMesaage(model.Message{Name: name, Message: message, Date: t.Format("01-02-2006 15:04:05")})
	http.Redirect(rw, r, "/", http.StatusMovedPermanently)
}
