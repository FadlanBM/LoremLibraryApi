package models

import "gorm.io/gorm"

type ListLending struct {
	gorm.Model
	BooksID   uint `json:"books_id"`
	LendingID uint `json:"lending_id"`
}
