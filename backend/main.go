package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"wesionary.team/dipeshdulal/accountengine/controllers/accounttype"
	"wesionary.team/dipeshdulal/accountengine/controllers/chartofaccount"
	"wesionary.team/dipeshdulal/accountengine/controllers/transactions"
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

	coa := route.Group("/coa")
	{
		coa.GET("/", chartofaccount.GetAllChartOfAccount)
		coa.GET("/:id", chartofaccount.GetOneChartOfAccount)
		coa.POST("/", chartofaccount.SaveChartOfAccount)
		coa.PATCH("/:id", chartofaccount.UpdateChartOfAccount)
		coa.DELETE("/:id", chartofaccount.DeleteChartOfAccount)
	}

	trx := route.Group("/trx")
	{
		trx.GET("/", transactions.GetAllTransactions)
		trx.GET("/:id", transactions.GetOneTransaction)
		trx.POST("/", transactions.SaveTransaction)
		trx.PATCH("/:id", transactions.UpdateTransaction)
		trx.DELETE("/:id", transactions.DeleteTransaction)
	}

	route.Run(":5000")
}
