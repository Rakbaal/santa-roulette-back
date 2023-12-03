package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Ping(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var version string
		db.QueryRow("SELECT VERSION()").Scan(&version)
		c.JSON(200, gin.H{"data": "Connected to: " + version})
	}
}
