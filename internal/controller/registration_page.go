package controller

import "github.com/gin-gonic/gin"

func RegistrationPage(c *gin.Context) {
	c.HTML(200, "registration.html", nil)
}
