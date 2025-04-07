package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"libraryManagment/internal/usecase"
)

type Server struct {
	e           *echo.Echo
	bookUseCase usecase.BookUseCase
	userUseCase usecase.UserUseCase
	loanUseCase usecase.LoanUseCase
}

// NewServer creates a new HTTP server instance
func NewServer(bookUseCase usecase.BookUseCase, userUseCase usecase.UserUseCase, loanUseCase usecase.LoanUseCase) *Server {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	return &Server{
		e:           e,
		bookUseCase: bookUseCase,
		userUseCase: userUseCase,
		loanUseCase: loanUseCase,
	}
}

// Start initializes the server and starts listening on the given port
func (s *Server) Start(port string) error {
	// Initialize handlers and routes
	NewBookHandler(s.e, s.bookUseCase)
	NewUserHandler(s.e, s.userUseCase)
	NewLoanHandler(s.e, s.loanUseCase)

	// Start server
	return s.e.Start(port)
}
