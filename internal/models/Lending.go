package models

import (
	"gorm.io/gorm"
	"time"
)

type Lending struct {
	gorm.Model
	BorrowersID uint      `json:"borrowers_id"`
	EmployeesID uint      `json:"employees_id"`
	DateLast    time.Time `json:"lastdate" form:"lastdate"`
	ReturnDate  time.Time `json:"returnDate" form:"returnDate"`
	BorrowDate  time.Time `json:"borrowdate" form:"borrowdate"`
	Status      string    `gorm:"type:varchar(10);default:false;" json:"status" form:"status"`
	Fine        []Fine
	ListLending []ListLending
}
