package models

import "gorm.io/gorm"

type Book_Category struct {
	gorm.Model
	CategoriesID uint `json:"categories_id"`
	BooksID      uint `json:"book_id"`
}
