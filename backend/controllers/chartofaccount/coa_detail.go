package chartofaccount

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"wesionary.team/dipeshdulal/accountengine/models"
)

// GetAllChartOfAccount gets all the chart of accounts
func GetAllChartOfAccount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var coas []models.ChartOfAccounts

	err := db.Find(&coas).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    coas,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error":   true,
		"message": err.Error(),
	})
}

// GetOneChartOfAccount get one chart of account
func GetOneChartOfAccount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var coa models.ChartOfAccounts
	var acctype models.AccountType
	id := c.Param("id")
	err := db.Where("id = ?", id).First(&coa).Error
	err = db.Model(&coa).Related(&acctype).Error
	coa.AccountType = acctype
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"chartofaccount": coa,
			},
		})
		return
	}

	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "coa not found",
			"error":   true,
		})
	}

}

type coaInput struct {
	Code          string `json:"code" binding:"required"`
	Description   string `json:"description" binding:"required"`
	AccountTypeID uint   `json:"accounttype_id" binding:"required"`
}

// SaveChartOfAccount to save chart of account
func SaveChartOfAccount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var coa models.ChartOfAccounts
	var input coaInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	copier.Copy(&coa, &input)
	_, err := coa.Save(db)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Coa saved",
		"coa":     coa,
	})

}

// UpdateChartOfAccount to update chart of a/c
func UpdateChartOfAccount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var coa models.ChartOfAccounts

	if err := db.Where("id = ?", id).First(&coa).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	var input coaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	copier.Copy(&coa, &input)

	if _, err := coa.Save(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "updated coa",
		"coa":     coa,
	})
}

// DeleteChartOfAccount to delete chart of account
func DeleteChartOfAccount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var coa models.ChartOfAccounts
	if err := db.Where("id = ?", id).First(&coa).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	if _, err := coa.Delete(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted coa",
		"coa":     coa,
	})

}
