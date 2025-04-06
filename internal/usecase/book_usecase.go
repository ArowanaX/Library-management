package usecase

import (
	"errors"
	"libraryManagment/internal/domain"
	"libraryManagment/internal/repo"
)

type BookUseCase interface {
	AddBook(book *domain.Book) error
	EditBook(id uint, fields map[string]interface{}) error
	DeleteBook(id uint) error
	GetBook(id string) (*domain.Book, error)
	GetBookList() ([]domain.Book, error)
}

type bookUseCase struct {
	bookRepo *repo.BookRepo
}

func NewBookUseCase(repo *repo.BookRepo) BookUseCase {
	return &bookUseCase{bookRepo: repo}
}

func (u *bookUseCase) AddBook(book *domain.Book) error {

	if book.Title == "" {
		return errors.New("کتاب باید عنوان داشته باشد")
	}
	if book.Author == "" {
		return errors.New("نام نویسنده الزامی است")
	}
	if book.ISBN == "" {
		return errors.New("شابک کتاب الزامی است")
	}

	return u.bookRepo.CreateBook(book)
}

func (u *bookUseCase) EditBook(id uint, fields map[string]interface{}) error {
	return u.bookRepo.UpdateBook(id, fields)
}

func (u *bookUseCase) DeleteBook(id uint) error {
	return u.bookRepo.DeleteBook(id)
}

func (u *bookUseCase) GetBook(id string) (*domain.Book, error) {
	return u.bookRepo.GetBookById(id)
}

func (u *bookUseCase) GetBookList() ([]domain.Book, error) {
	return u.bookRepo.GetBookList()
}
