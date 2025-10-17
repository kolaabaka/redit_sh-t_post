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
	r.ServeFiles("/public/*filepath", http.Dir("./public"))
	r.GET("/", controller.MessageWall)
	r.GET("/new", controller.NewMessageWall)
	r.POST("/create_message", controller.CreateMessage)
}
