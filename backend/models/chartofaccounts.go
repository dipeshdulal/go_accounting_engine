// chart of account model

package models

import (
	"github.com/jinzhu/gorm"
)

// ChartOfAccounts model
type ChartOfAccounts struct {
	gorm.Model
	Code          string `json:"code"`
	Description   string `json:"description"`
	AccountTypeID uint
	AccountType   AccountType
}

// Save chart of account
func (coa *ChartOfAccounts) Save(db *gorm.DB) (*ChartOfAccounts, error) {

	var err error
	err = db.Save(coa).Error
	if err != nil {
		return &ChartOfAccounts{}, err
	}

	return coa, nil
}

// Delete chart of account
func (coa *ChartOfAccounts) Delete(db *gorm.DB) (*ChartOfAccounts, error) {
	return coa, db.Delete(coa).Error
}
