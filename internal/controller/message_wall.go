package controller

import (
	"fmt"
	"goSiteProject/internal/model"
	"goSiteProject/internal/monitoring"
	"goSiteProject/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

func MessageWall(c *gin.Context) {
	topic := c.DefaultQuery("topic", "message_main_table")

	messageList, err := service.GetMesaages(topic)

	//Topic is not exists
	if err != nil {
		c.String(404, fmt.Sprintf("No topic with name %s", topic))
		return
	}

	c.HTML(200, "index.html", gin.H{
		"Header":      topic,
		"MessageList": messageList,
	})
	monitoring.IncrementEndpointHttpCounter("/home")
	monitoring.IncrementTotalhttpCounter()
}

func NewMessageWall(c *gin.Context) {
	c.HTML(200, "new_message_form.html", nil)

	monitoring.IncrementEndpointHttpCounter("/new")
	monitoring.IncrementTotalhttpCounter()
}

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
