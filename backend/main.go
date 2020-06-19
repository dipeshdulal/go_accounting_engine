package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"wesionary.team/dipeshdulal/accountengine/controllers/accounttype"
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

	accType := route.Group("/acctype")
	{
		accType.GET("/", accounttype.GetAllAccount)
		accType.POST("/", accounttype.SaveAccount)
		accType.GET("/:id", accounttype.GetOneAccount)
		accType.PATCH("/:id", accounttype.EditAccount)
		accType.DELETE("/:id", accounttype.DeleteAccount)
	}

	route.Run(":5000")
}
