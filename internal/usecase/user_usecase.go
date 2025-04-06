package usecase

import (
	"libraryManagment/internal/domain"
	"libraryManagment/internal/repo"
)

type UserUseCase interface {
	FindByID(uint) (*domain.User, error)
	Register(user *domain.User) error
	Login(email string, password string) (*domain.User, error)
}
type userUseCase struct {
	userRepo *repo.UserRepo
}

func NewUserUseCase(Repo *repo.UserRepo) UserUseCase {
	return &userUseCase{userRepo: Repo}
}
func (u *userUseCase) FindByID(id uint) (*domain.User, error) {
	return u.userRepo.FindByID(id)
}
func (u *userUseCase) Register(user *domain.User) error {
	return u.userRepo.Register(user)
}
func (u *userUseCase) Login(email string, password string) (*domain.User, error) {
	return u.userRepo.Login(email, password)
}
