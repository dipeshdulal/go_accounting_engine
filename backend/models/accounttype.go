package models

import (
	"github.com/jinzhu/gorm"
)

// AccountType model
type AccountType struct {
	gorm.Model
	Name string `json:"name"`
	Code string `json:"code"`
}

// Save accountType
func (act *AccountType) Save(db *gorm.DB) *AccountType {
	db.Save(act)
	return act
}

// Delete account type
func (act *AccountType) Delete(db *gorm.DB) {
	db.Delete(act)
}
