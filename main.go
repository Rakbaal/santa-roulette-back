package main

import (
	"database/sql"
	"santa-roulette/controllers"
	"santa-roulette/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	db, err := sql.Open("mysql", "root:Ragnarok0@tcp(127.0.0.1:40000)/santaroulette")
	utils.Check(err)
	defer db.Close()

	router.Use(utils.Headers())

	router.GET("/ping", controllers.Ping)

	router.GET("/own/:ownerid/:famille", controllers.Own(db))

	router.GET("/user/:pseudo", controllers.GetUserByName(db))

	router.GET("/owned/:id", controllers.GetOwned(db))

	router.GET("/images", controllers.GetImages(db))

	router.Run(":8081")
}
