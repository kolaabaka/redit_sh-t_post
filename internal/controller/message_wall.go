package controller

import (
	"fmt"
	"goSiteProject/internal/monitoring"
	"goSiteProject/internal/service"

	"github.com/gin-gonic/gin"
)

func MessageWall(c *gin.Context) {
	topic := c.DefaultQuery("topic", "main_table")

	messageList, err := service.GetMesaages(topic)

	//Topic is not exists
	if err != nil {
		c.String(404, fmt.Sprintf("No topic with name %s", topic))
	}

	c.HTML(200, "index.html", gin.H{
		"Header":      topic,
		"MessageList": messageList,
	})
	monitoring.IncrementEndpointHttpCounter("/home")
	monitoring.IncrementTotalhttpCounter()
}
