package repo

import (
	"errors"
	"gorm.io/gorm"
	"libraryManagment/internal/domain"
	"time"
)

type LoanRepo struct {
	DB *gorm.DB
}

func NewLoanRepo(db *gorm.DB) *LoanRepo {
	return &LoanRepo{DB: db}
}

func (repo LoanRepo) BookExist(book *domain.Book) (bool, error) {
	if err := repo.DB.First(&book).Error; err != nil {
		return false, errors.New("book not found")
	}
	if book.Copies <= 0 {
		return false, errors.New("no copies available")
	}
	return true, nil

}
func (repo LoanRepo) GetUserLoanObj(userID uint, bookID uint) (*domain.Loan, error) {
	var loan *domain.Loan
	if err := repo.DB.Where("book_id = ? and user_id = ?", bookID, userID).First(&loan).Error; err != nil {
		return nil, err
	}
	return loan, nil
}
func (repo LoanRepo) ReturnLoan(loan *domain.Loan) error {
	loan.Returned = true
	if err := repo.DB.Save(loan).Error; err != nil {
		return err
	}
	return nil
}
func (repo LoanRepo) DecreaseCopies(book *domain.Book) error {
	book.Copies -= 1
	if err := repo.DB.Save(&book).Error; err != nil {
		return err
	}
	return nil
}
func (repo LoanRepo) IncreaseCopies(book *domain.Book) error {
	book.Copies += 1
	if err := repo.DB.Save(&book).Error; err != nil {
		return err
	}
	return nil
}
func (repo LoanRepo) GetBookByID(bookID uint) (*domain.Book, error) {
	var book *domain.Book
	if err := repo.DB.First(&book, bookID).Error; err != nil {
		return nil, err
	}
	return book, nil
}
func (repo LoanRepo) LoanBook(userID uint, bookID uint) (*domain.Loan, error) {

	// Create new loan
	loan := domain.Loan{
		UserID:   userID,
		BookID:   bookID,
		DueDate:  time.Now().AddDate(0, 0, 7), // One week until now for loan
		Returned: false,
	}

	if err := repo.DB.Create(&loan).Error; err != nil {
		return nil, err
	}

	return &loan, nil
}
func (repo LoanRepo) UserLoanList(userID uint) ([]*domain.Loan, error) {
	var loans []*domain.Loan
	if err := repo.DB.Where("user_id = ?", userID).Find(&loans).Error; err != nil {
		return nil, err
	}
	return loans, nil
}
func (repo LoanRepo) HasActiveLoans(userID uint) (bool, error) {
	var count int64

	err := repo.DB.Model(&domain.Loan{}).
		Where("user_id = ? AND return_date IS NULL", userID).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
