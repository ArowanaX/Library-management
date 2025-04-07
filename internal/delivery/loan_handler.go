package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"libraryManagment/internal/usecase"
)

type LoanHandler struct {
	loanUseCase usecase.LoanUseCase
}

func NewLoanHandler(e *echo.Echo, loanUseCase usecase.LoanUseCase) {
	handler := &LoanHandler{
		loanUseCase: loanUseCase,
	}

	// Loan routes
	loanGroup := e.Group("/api/loans")
	loanGroup.GET("/user/:user_id", handler.GetUserLoans)
	loanGroup.POST("/borrow", handler.BorrowBook)
	loanGroup.POST("/return", handler.ReturnBook)
}

// GetUserLoans returns all loans for a specific user
func (h *LoanHandler) GetUserLoans(c echo.Context) error {
	userID := c.Param("user_id")
	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid user ID format",
		})
	}

	loans, err := h.loanUseCase.UserLoanList(uint(userIDUint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, loans)
}

// BorrowBook creates a new loan
func (h *LoanHandler) BorrowBook(c echo.Context) error {
	// Parse loan request
	loanRequest := struct {
		UserID uint `json:"user_id"`
		BookID uint `json:"book_id"`
	}{}

	if err := c.Bind(&loanRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	loan, err := h.loanUseCase.LoanBook(loanRequest.UserID, loanRequest.BookID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, loan)
}

// ReturnBook marks a loan as returned
func (h *LoanHandler) ReturnBook(c echo.Context) error {
	// Parse return request
	returnRequest := struct {
		UserID uint `json:"user_id"`
		BookID uint `json:"book_id"`
	}{}

	if err := c.Bind(&returnRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	err := h.loanUseCase.ReturnLoan(returnRequest.UserID, returnRequest.BookID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Book returned successfully",
	})
}
