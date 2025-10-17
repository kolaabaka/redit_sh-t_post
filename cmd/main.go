package main

import (
	"goSiteProject/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	routes(r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

func routes(r *httprouter.Router) {

	r.ServeFiles("/css/*filepath", http.Dir("./public/css"))
	r.ServeFiles("/js/*filepath", http.Dir("./public/js"))
	r.ServeFiles("/img/*filepath", http.Dir("./public/img"))

	r.GET("/", controller.MessageWall)
	r.GET("/new", controller.NewMessageWall)
	r.POST("/create_message", controller.CreateMessage)
}
