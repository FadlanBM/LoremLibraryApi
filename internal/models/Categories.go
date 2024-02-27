package models

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	Name          string `gorm:"type:varchar(50);not null" json:"name"`
	Book_Category []Book_Category
}
