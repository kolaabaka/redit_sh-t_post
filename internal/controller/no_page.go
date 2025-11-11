package controller

import (
	"goSiteProject/internal/monitoring"

	"github.com/gin-gonic/gin"
)

func NoPage(c *gin.Context) {

	c.HTML(404, "no_page.html", nil)
	monitoring.IncrementTotalhttpCounter()
}
