package controllers

import (
	"database/sql"
	"fmt"
	"santa-roulette/models"

	"github.com/gin-gonic/gin"
)

func GetUserByName(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := fmt.Sprintf("call GetUserByPseudo(\"%v\")", c.Param("pseudo"))
		row := db.QueryRow(query)
		var id int
		var pseudo string
		var famille int
		var photo string
		row.Scan(&id, &pseudo, &famille, &photo)
		participant := models.Participant{
			Id:      id,
			Pseudo:  pseudo,
			Famille: famille,
			Photo:   photo,
		}
		c.JSON(200, gin.H{"data": participant})
	}
}
