package models

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	BorrowersID uint `json:"borrowers_id"`
	BooksID     uint `json:"book_id"`
}
