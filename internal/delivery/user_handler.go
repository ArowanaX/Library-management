package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"libraryManagment/internal/domain"
	"libraryManagment/internal/usecase"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(e *echo.Echo, userUseCase usecase.UserUseCase) {
	handler := &UserHandler{
		userUseCase: userUseCase,
	}

	// User routes
	userGroup := e.Group("/api/users")
	userGroup.POST("/register", handler.Register)
	userGroup.POST("/login", handler.Login)
	userGroup.GET("/:id", handler.GetUser)
}

// Register creates a new user
func (h *UserHandler) Register(c echo.Context) error {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	if err := h.userUseCase.Register(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, user)
}

// Login authenticates a user
func (h *UserHandler) Login(c echo.Context) error {
	loginData := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	user, err := h.userUseCase.Login(loginData.Email, loginData.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid credentials",
		})
	}
	return c.JSON(http.StatusOK, user)
}

// GetUser returns a user by ID
func (h *UserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID format",
		})
	}

	user, err := h.userUseCase.FindByID(uint(idUint))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "User not found",
		})
	}
	return c.JSON(http.StatusOK, user)
}
