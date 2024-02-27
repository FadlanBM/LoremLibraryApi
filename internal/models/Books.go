package models

import (
	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Title         string `gorm:"type:varchar(100);not null" json:"title"`
	Author        string `gorm:"type:varchar(50)" json:"author"`
	No_Inventaris string `gorm:"type:varchar(50)" json:"no_inventaris"`
	Publisher     string `gorm:"type:varchar(200)" json:"publisher"`
	Code          string `gorm:"type:varchar(200)" json:"code"`
	Description   string `gorm:"type:text" json:"description"`
	Img           string `gorm:"type:varchar(255)" json:"img"`
	YearPublished uint   `json:"year_published"`
	Book_Category []Book_Category
	Collection    []Collection
	ListLending   []ListLending
	Reviews       []Reviews
}
