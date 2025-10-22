package main

import (
	"goSiteProject/internal/controller"
	"goSiteProject/internal/service"
	"log/slog"
	"net/http"
	"net/http/pprof"
	"os"

	colored_logger "github.com/kolaabaka/coloured_logger"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := slog.New(colored_logger.NewSimpleHandler(os.Stdout, slog.LevelDebug))

	r := httprouter.New()
	routes(r)

	//Check SQLite connection
	service.MustCheckConnection(*logger)

	//Handler for default profiler "pprof"
	r.Handler("GET", "/debug/pprof/*item", http.HandlerFunc(pprof.Index))

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
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
