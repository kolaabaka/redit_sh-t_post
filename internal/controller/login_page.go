package controller

import (
	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
