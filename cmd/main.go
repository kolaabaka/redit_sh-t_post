package main

import (
	"goSiteProject/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP) //SIGTERM - kill <pid>, SIGHUP - close terminal

	app.Run(done)
}
