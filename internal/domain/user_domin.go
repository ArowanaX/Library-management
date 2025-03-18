package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Email    string `gorm:"type:varchar(50);unique;not null"`
	Password string `gorm:"type:varchar(15);not null"`
}
