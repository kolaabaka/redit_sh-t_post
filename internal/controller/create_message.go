package controller

import (
	"goSiteProject/internal/model"
	"goSiteProject/internal/monitoring"
	"goSiteProject/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateMessage(c *gin.Context) {
	message := c.PostForm("comment")
	name := c.PostForm("name")
	topic := c.PostForm("topic")

	t := time.Now()

	if topic != "" {
		service.AddMesaage(topic, model.Message{Name: name, Message: message, Date: t.Format("01-02-2006 15:04:05")})
	}

	c.Redirect(301, "/")

	monitoring.IncrementEndpointHttpCounter("/create_message")
	monitoring.IncrementTotalhttpCounter()
}
