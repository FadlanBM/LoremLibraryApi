package models

import "gorm.io/gorm"

type Employees struct {
	gorm.Model
	Google_ID    string `gorm:"type:varchar(255);not null" json:"google_id" form:"google_id"`
	Email        string `gorm:"type:varchar(255);not null" json:"email" form:"email"`
	Name         string `gorm:"type:varchar(50)" json:"name" form:"name"`
	Phone_Number string `gorm:"type:varchar(35)" json:"phone_number" form:"phone_number"`
	Address      string `gorm:"type:varchar(50)" json:"address" form:"address"`
	Role         string `gorm:"type:varchar(255);default:employee" json:"role" form:"role"`
	Avatar       string `gorm:"type:varchar(255)" json:"avatar" form:"avatar"`
	Active       string `gorm:"type:varchar(10);default:false" json:"active" form:"active"`
}
