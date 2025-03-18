package domain

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Title       string    `gorm:"type:varchar(255);not null"`
	Author      string    `gorm:"type:varchar(255);not null"`
	ISBN        string    `gorm:"type:varchar(20);unique;not null"`
	PublishedAt time.Time `gorm:"not null"`
	Copies      int       `gorm:"not null;default:1"`
}
