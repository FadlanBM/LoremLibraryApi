package models

import (
	"gorm.io/gorm"
	"time"
)

type Lending struct {
	gorm.Model
	BorrowersID uint       `json:"borrowers_id"`
	EmployeesID uint       `gorm:"nullable;" json:"employees_id"`
	DateLast    *time.Time `json:"lastdate" form:"lastdate"`
	ReturnDate  time.Time  `json:"returnDate" form:"returnDate"`
	BorrowDate  time.Time  `json:"borrowdate" form:"borrowdate"`
	Code        string     `gorm:"type:varchar(200)" json:"code"`
	Status      string     `gorm:"type:varchar(10);default:false;" json:"status" form:"status"`
	Fine        []Fine
	ListLending []ListLending
}
