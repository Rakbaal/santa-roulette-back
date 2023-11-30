package utils

import "github.com/gin-gonic/gin"

func Headers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
}
