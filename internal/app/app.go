package app

import (
	"context"
	"goSiteProject/internal/controller"
	"goSiteProject/internal/middleware"
	"goSiteProject/internal/monitoring"
	"goSiteProject/internal/repository"
	"goSiteProject/internal/service"
	"log/slog"
	"net/http"
	"net/http/pprof"
	"os"
	"path"

	"github.com/gin-gonic/gin"

	colored_logger "github.com/kolaabaka/coloured_logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Run(done chan os.Signal) {
	//server withoud default configuration
	r := gin.Default()
	SetupRoutes(r)

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

	<-done
	logger.Warn("SERVER was interrupted")
	server.Shutdown(context)
}

func SetupRoutes(r *gin.Engine) {
	templatePath := os.Getenv("TEMPLATE_PATH")
	if templatePath == "" {
		templatePath = "."
	}

	r.Static("/static", "./public")
	r.LoadHTMLGlob(path.Join(templatePath, "template/*"))

	r.GET("/", controller.MessageWall)

	r.GET("/new", controller.NewMessageWall)
	r.POST("/create_message", controller.CreateMessage)

	r.GET("/login", controller.LoginPage)
	r.POST("/login_form", controller.LoginFromPage)
	r.GET("/log_out", controller.LogOut)

	r.GET("/registration", controller.RegistrationPage)
	r.POST("/registration_form", controller.RegistrationFormPage)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddl())
	{
		auth.GET("/profile", controller.AuthProfile)
	}

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	//Handler for default profiler "pprof", using in Prometheus
	r.GET("/debug/pprof/*item", gin.WrapH(http.HandlerFunc(pprof.Index)))

	r.NoRoute(controller.NoPage)
}
