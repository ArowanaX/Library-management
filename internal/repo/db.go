package repo

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"libraryManagment/config"
	"libraryManagment/internal/domain"
	"log"
)

func InitDB(cfg *config.Config) {
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// migrate
	if err := db.AutoMigrate(&domain.Book{}, &domain.User{}, &domain.Loan{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	return db
}
