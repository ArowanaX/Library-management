package main

import (
	"fmt"
	"libraryManagment/config"
	"libraryManagment/internal/repo"
)

func main() {
	cfg := config.LoadConfig()
	db := repo.InitDB(cfg)
	fmt.Println("Database connected!")
}
