package main

import (
	"goSiteProject/internal/controller"
	"goSiteProject/internal/monitoring"
	"goSiteProject/internal/repository"
	"goSiteProject/internal/service"
	"log/slog"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/gin-gonic/gin"
	colored_logger "github.com/kolaabaka/coloured_logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// TODO: switch httprouter to GIN
func main() {
	logger := slog.New(colored_logger.NewSimpleHandler(os.Stdout, slog.LevelDebug))

	monitoring.MustInitPrometheusStat()

	service.MustInitService(logger)
	//Check SQLite connection
	repository.MustCheckConnection(logger)

	r := gin.Default()
	routes(r)

	//Handler for default profiler "pprof", using in Prometheus
	r.GET("/debug/pprof/*item", gin.WrapH(http.HandlerFunc(pprof.Index)))

	err := r.Run(":8080")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func routes(r *gin.Engine) {
	r.Static("/static", "./public")
	r.LoadHTMLGlob("template/*")

	r.GET("/", controller.MessageWall)

	r.GET("/new", controller.NewMessageWall)
	r.POST("/create_message", controller.CreateMessage)

	r.GET("/login", controller.LoginPage)
	r.POST("/login_form", controller.LoginFromPage)

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.NoRoute(controller.NoPage)
}
