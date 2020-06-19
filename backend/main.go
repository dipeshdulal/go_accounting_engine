package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"wesionary.team/dipeshdulal/accountengine/models"
)

func main() {
	log.Print("Initializing GIN")
	route := gin.Default()

	db := models.InitModels()

	// set variable
	route.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	route.Run(":5000")
}
