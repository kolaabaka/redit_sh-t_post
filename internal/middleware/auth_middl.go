package middleware

import "github.com/gin-gonic/gin"

func AuthMiddl() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("session_token")
		if err != nil {
			c.HTML(401, "profile.html", gin.H{
				"Message": "YOU ARE NOT USER",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
