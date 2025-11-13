package main

import (
	"context"
	"goSiteProject/internal/controller"
	"goSiteProject/internal/monitoring"
	"goSiteProject/internal/repository"
	"goSiteProject/internal/service"
	"log/slog"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	colored_logger "github.com/kolaabaka/coloured_logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// TODO: switch httprouter to GIN
func main() {
	//server withoud default configuration
	r := gin.Default()
	routes(r)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go server.ListenAndServe()

	logger := slog.New(colored_logger.NewSimpleHandler(os.Stdout, slog.LevelDebug))

	monitoring.MustInitPrometheusStat()

	service.MustInitService(logger)
	//Check SQLite connection
	repository.MustCheckConnection(logger)

	context := context.TODO() //funny find =)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP) //SIGTERM - kill <pid>, SIGHUP - close terminal
	<-done
	logger.Warn("OS was interrupted")
	server.Shutdown(context)
}

func routes(r *gin.Engine) {
	r.Static("/static", "./public")
	r.LoadHTMLGlob("template/*")

	r.GET("/", controller.MessageWall)

	r.GET("/new", controller.NewMessageWall)
	r.POST("/create_message", controller.CreateMessage)

	r.GET("/login", controller.LoginPage)
	r.POST("/login_form", controller.LoginFromPage)

	r.GET("/registration", controller.RegistrationPage)
	r.POST("/registration_form", controller.RegistrationFormPage)

	auth := r.Group("auth")
	{
		auth.GET("/profile", nil)
	}

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	//Handler for default profiler "pprof", using in Prometheus
	r.GET("/debug/pprof/*item", gin.WrapH(http.HandlerFunc(pprof.Index)))

	r.NoRoute(controller.NoPage)
}
