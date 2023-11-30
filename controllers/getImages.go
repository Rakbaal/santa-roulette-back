package controllers

import (
	"database/sql"
	"santa-roulette/utils"

	"github.com/gin-gonic/gin"
)

func GetImages(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := "call GetImages()"
		rows, err := db.Query(query)
		utils.Check(err)

		var images []string
		for rows.Next() {
			var image string
			err := rows.Scan(&image)
			utils.Check(err)
			images = append(images, image)
		}
		c.JSON(200, gin.H{"data": images})
	}
}
