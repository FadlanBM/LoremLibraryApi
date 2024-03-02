package models

import (
	"gorm.io/gorm"
)

type Reviews struct {
	gorm.Model
	BorrowersID uint `json:"borrowers_id"`
	BooksID     uint `json:"buku_id"`
	Rating      float32
	Messages    string `gorm:"type:text" json:"messages"`
}
