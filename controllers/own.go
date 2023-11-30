package controllers

import (
	"database/sql"
	"fmt"
	"math/rand"
	"santa-roulette/utils"

	"github.com/gin-gonic/gin"
)

func Own(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		notOwnedParticipants := utils.GetNotOwned(db, c.Param("famille"))
		owned := notOwnedParticipants[rand.Intn(len(notOwnedParticipants))]

		query := fmt.Sprintf("call SetOwnership(%v, %v)", c.Param("ownerid"), owned.Id)
		db.Exec(query)
		c.JSON(200, gin.H{"data": "Success"})
	}
}
