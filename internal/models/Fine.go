package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Fine struct {
	gorm.Model
	LendingID uint            `json:"lending_id"`
	Nominal   decimal.Decimal `gorm:"type:decimal(10,2);" json:"nominal"`
	Pay       decimal.Decimal `gorm:"type:decimal(10,2);" json:"pay"`
	Status    string          `gorm:"type:varchar(10);default:false;" json:"status" form:"status"`
}
