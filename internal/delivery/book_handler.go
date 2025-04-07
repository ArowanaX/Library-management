package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"libraryManagment/internal/domain"
	"libraryManagment/internal/usecase"
)

type BookHandler struct {
	bookUseCase usecase.BookUseCase
}

func NewBookHandler(e *echo.Echo, bookUseCase usecase.BookUseCase) {
	handler := &BookHandler{
		bookUseCase: bookUseCase,
	}

	// Book routes
	bookGroup := e.Group("/api/books")
	bookGroup.GET("", handler.GetAllBooks)
	bookGroup.GET("/:id", handler.GetBook)
	bookGroup.POST("", handler.CreateBook)
	bookGroup.PUT("/:id", handler.UpdateBook)
	bookGroup.DELETE("/:id", handler.DeleteBook)
}

// GetAllBooks returns all books
func (h *BookHandler) GetAllBooks(c echo.Context) error {
	books, err := h.bookUseCase.GetBookList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, books)
}

// GetBook returns a specific book by ID
func (h *BookHandler) GetBook(c echo.Context) error {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{})
	}
	book, err := h.bookUseCase.GetBook(uint(idUint))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Book not found",
		})
	}
	return c.JSON(http.StatusOK, book)
}

// CreateBook creates a new book
func (h *BookHandler) CreateBook(c echo.Context) error {
	book := new(domain.Book)
	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	if err := h.bookUseCase.AddBook(book); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, book)
}

// UpdateBook updates an existing book
func (h *BookHandler) UpdateBook(c echo.Context) error {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID format",
		})
	}

	// Get fields to update
	updateFields := make(map[string]interface{})
	if err := c.Bind(&updateFields); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	// Update book
	if err := h.bookUseCase.EditBook(uint(idUint), updateFields); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Book updated successfully",
	})
}

// DeleteBook deletes a book by ID
func (h *BookHandler) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID format",
		})
	}

	if err := h.bookUseCase.DeleteBook(uint(idUint)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Book deleted successfully",
	})
}
