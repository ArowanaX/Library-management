package repo

import (
	"gorm.io/gorm"
	"libraryManagment/internal/domain"
)

type LoanRepo struct {
	DB *gorm.DB
}

func NewLoanRepo(db *gorm.DB) *LoanRepo {
	return LoanRepo{}
}

func (repo LoanRepo) LoanBook(user_id int, book_id int)(*domain.Loan, error) {

}
func ()