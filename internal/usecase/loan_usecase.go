package usecase

import (
	"errors"
	"libraryManagment/internal/domain"
	"libraryManagment/internal/repo"
)

type LoanUseCase interface {
	UserLoanList(userID uint) ([]*domain.Loan, error)
	LoanBook(userID uint, bookID uint) (*domain.Loan, error)
	ReturnLoan(userID uint, bookID uint) error
}

type loanUseCase struct {
	loanRepo *repo.LoanRepo
}

func NewLoanUseCase(repo *repo.LoanRepo) LoanUseCase {
	return &loanUseCase{loanRepo: repo}
}

func (l *loanUseCase) UserLoanList(userID uint) ([]*domain.Loan, error) {
	return l.loanRepo.UserLoanList(userID)
}

func (l *loanUseCase) LoanBook(userID uint, bookID uint) (*domain.Loan, error) {
	// Get book information from database
	book, err := l.loanRepo.GetBookByID(bookID)
	if err != nil {
		return nil, err
	}

	// Check if the book exists in inventory
	bookExists, err := l.loanRepo.BookExist(book)
	if err != nil {
		return nil, err
	}
	if !bookExists {
		return nil, errors.New("book does not exist in inventory")
	}

	// Check if user has any active loans that would prevent a new loan
	hasActiveLoans, err := l.loanRepo.HasActiveLoans(userID)
	if err != nil {
		return nil, err
	}
	if hasActiveLoans {
		return nil, err
	}

	// Process the loan after all validations passed
	loan, err := l.loanRepo.LoanBook(userID, bookID)
	if err != nil {
		return nil, err
	}

	// Update book inventory by decreasing available copies
	err = l.loanRepo.DecreaseCopies(book)
	if err != nil {
		// Consider rollback of loan creation here if decreasing copies fails
		return nil, err
	}

	return loan, nil
}
func (l *loanUseCase) ReturnLoan(userID uint, bookID uint) error {
	loan, err := l.loanRepo.GetUserLoanObj(userID, bookID)
	if err != nil {
		return err
	}
	err = l.loanRepo.ReturnLoan(loan)
	if err != nil {
		return err
	}
	book, err := l.loanRepo.GetBookByID(bookID)
	if err != nil {
		return err
	}
	err = l.loanRepo.IncreaseCopies(book)
	if err != nil {
		return err
	}
	return nil
}
