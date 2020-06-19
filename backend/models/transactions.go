package models

import "github.com/jinzhu/gorm"

// Transactions model
type Transactions struct {
	gorm.Model
	Name              string `json:"name"`
	ChartOfAccountsID uint   `json:"chart_of_accounts_id"`
	ChartOfAccounts   ChartOfAccounts
	Amount            float64 `json:"amount"`
	IsDebit           bool    `json:"isDebit"`
}

// Save transaction
func (trx *Transactions) Save(db *gorm.DB) (*Transactions, error) {

	var coa ChartOfAccounts
	if err := db.Where("id = ?", trx.ChartOfAccountsID).First(&coa).Error; err != nil {
		return &Transactions{}, err
	}

	if err := db.Save(trx).Error; err != nil {
		return &Transactions{}, err
	}
	trx.ChartOfAccounts = coa
	return trx, nil
}

// Delete transaction
func (trx *Transactions) Delete(db *gorm.DB) (*Transactions, error) {
	if err := db.Delete(trx).Error; err != nil {
		return &Transactions{}, err
	}
	return trx, nil
}
