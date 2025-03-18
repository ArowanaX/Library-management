package domain

import (
	"gorm.io/gorm"
	"time"
)

type Loan struct {
	gorm.Model
	BookID   uint      `gorm:"not null"`
	UserID   uint      `gorm:"not null"`
	DueDate  time.Time `gorm:"not null"`
	Returned bool      `gorm:"not null;default:false"`
}
