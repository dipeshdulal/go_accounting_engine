package transactions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"wesionary.team/dipeshdulal/accountengine/models"
)

// GetAllTransactions get all trx's
func GetAllTransactions(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var trxs []models.Transactions
	if err := db.Find(&trxs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    trxs,
	})
}

// GetOneTransaction gets one trx
func GetOneTransaction(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var trx models.Transactions
	if err := db.Where("id = ?", id).First(&trx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	// change this later after
	// good way of eager loading is found

	var coa models.ChartOfAccounts
	var acctype models.AccountType

	db.Model(&trx).Related(&coa)
	db.Model(&coa).Related(&acctype)
	coa.AccountType = acctype
	trx.ChartOfAccounts = coa

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": gin.H{
			"trx": trx,
		},
	})

}

// UpdateTransaction update trx
func UpdateTransaction(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var trx models.Transactions
	var input trxInput

	if err := db.Where("id = ?", id).First(&trx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	copier.Copy(&trx, &input)

	if _, err := trx.Save(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Saved successfully",
		"trx":     trx,
	})
}

type trxInput struct {
	Name              string  `json:"name" binding:"required"`
	ChartOfAccountsID uint    `json:"chart_of_accounts_id" binding:"required"`
	Amount            float64 `json:"amount" binding:"required"`
	IsDebit           bool    `json:"isDebit" binding:"required"`
}

// SaveTransaction saves trx
func SaveTransaction(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var trx models.Transactions
	var input trxInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	copier.Copy(&trx, &input)

	if _, err := trx.Save(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Saved successfully",
		"trx":     trx,
	})
}

// DeleteTransaction deletes trx
func DeleteTransaction(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var trx models.Transactions
	if err := db.Where("id = ?", id).First(&trx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	if _, err := trx.Delete(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted successfully",
		"trx":     trx,
	})

}
