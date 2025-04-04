package repo

import (
	"gorm.io/gorm"
	"libraryManagment/internal/domain"
)

type BookRepo struct {
	DB *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepo {
	return &BookRepo{DB: db}
}
func (repo *BookRepo) CreateBook(book *domain.Book) error {
	return repo.DB.Create(book).Error
}
func (repo *BookRepo) UpdateBook(id uint, fields map[string]interface{}) error {
	return repo.DB.Model(&domain.Book{}).Where("id = ?", id).Updates(fields).Error
}
func (repo *BookRepo) DeleteBook(id uint) error {
	return repo.DB.Delete(&domain.Book{}, id).Error
}
func (repo *BookRepo) GetBookById(id string) (*domain.Book, error) {
	var book domain.Book
	err := repo.DB.Where("id = ?", id).First(&book).Error
	return &book, err
}
func (repo *BookRepo) GetBookList() ([]domain.Book, error) {
	var books []domain.Book
	err := repo.DB.Find(&books).Error
	return books, err
}
