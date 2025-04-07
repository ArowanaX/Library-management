package repo

import (
	"gorm.io/gorm"
	"libraryManagment/internal/domain"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{DB: db}
}

func (repo UserRepo) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := repo.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (repo UserRepo) Register(user *domain.User) error {
	return repo.DB.Create(&user).Error
}
func (repo UserRepo) Login(email string, password string) (*domain.User, error) {
	var user domain.User
	if err := repo.DB.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
