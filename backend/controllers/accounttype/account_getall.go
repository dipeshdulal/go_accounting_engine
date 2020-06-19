package accounttype

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"wesionary.team/dipeshdulal/accountengine/models"
)

// GetAllAccount gets all the account
func GetAllAccount(c *gin.Context) {
	_db, _ := c.Get("db")
	if _db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "cannot get db value",
		})
		return
	}
	db, _ := _db.(*gorm.DB)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "cannot get db value",
		})
	}
	var accounts []models.AccountType
	db.Find(&accounts)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    accounts,
	})
}

type createAccountInput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// SaveAccount function to save the account
func SaveAccount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input createAccountInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	var accounttype models.AccountType
	copier.Copy(&accounttype, &input)
	accounttype.Save(db)

	c.JSON(http.StatusInternalServerError, gin.H{
		"message":     "account saved.",
		"accounttype": accounttype,
	})
}

// EditAccount to edit accounttype
func EditAccount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input createAccountInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
	}

	var account models.AccountType
	id := c.Param("id")
	err := db.Where("id = ?", id).First(&account).Error

	if err == nil {
		copier.Copy(&account, &input)
		account.Save(db)
		c.JSON(http.StatusOK, gin.H{
			"message": "save successfully",
			"data":    account,
		})
		return
	}

	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "account type not found",
			"error":   true,
		})
	}

}

// GetOneAccount function gets one account
func GetOneAccount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var account models.AccountType
	err := db.Where("id = ?", id).First(&account).Error

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"account": account,
			},
		})
		return
	}

	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "account type not found",
			"error":   true,
		})
	}
}

// DeleteAccount delete the accounttype resource
func DeleteAccount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var account models.AccountType
	err := db.Where("id = ?", id).First(&account).Error

	if err == nil {
		account.Delete(db)
		c.JSON(http.StatusOK, gin.H{
			"message": "account deleted successfully",
			"account": account,
		})
		return
	}

	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "account type not found",
			"error":   true,
		})
	}
}
