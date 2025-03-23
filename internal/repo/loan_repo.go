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

func NewLoanRepo(db *gorm.DB) LoanRepo {
	return LoanRepo{}
}

func (repo LoanRepo) LoanBook(userID uint, bookID uint) (*domain.Loan, error) {
	var book domain.Book
	// Book exist
	if err := repo.DB.First(&book, bookID).Error; err != nil {
		return nil, errors.New("book not found")
	}
	if book.Copies <= 0 {
		return nil, errors.New("no copies available")
	}

	// Create new loan
	loan := domain.Loan{
		UserID:   uint(userID),
		BookID:   uint(bookID),
		DueDate:  time.Now().AddDate(0, 0, 7), // One week until now for loan
		Returned: false,
	}

	if err := repo.DB.Create(&loan).Error; err != nil {
		return nil, err
	}

	// Decrease on copy
	book.Copies -= 1
	if err := repo.DB.Save(&book).Error; err != nil {
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

func (repo LoanRepo) UserReturnLoan(bookID uint, userID uint) ([]*domain.Loan, error) {
	//Get loan and book
	var loans *domain.Loan
	if err := repo.DB.Where("book_id = ? and user_id = ?", bookID, userID).First(&loans).Error; err != nil {
		return nil, err
	}

	var book domain.Book
	// Book exist
	if err := repo.DB.First(&book, bookID).Error; err != nil {
		return nil, errors.New("book not found")
	}
	//Complete Loan
	loans.Returned = true
	if err := repo.DB.Save(loans).Error; err != nil {
		return nil, err
	}
	//++1 Copies
	book.Copies += 1
	if err := repo.DB.Save(&book).Error; err != nil {
		return nil, err
	}
	return []*domain.Loan{loans}, nil
}
