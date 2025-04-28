package services

import (
	"github.com/ark1500j/repo-pattern/internal/models"
	"github.com/ark1500j/repo-pattern/internal/repositories"
)

type UserService interface {
	CreateUser(user *models.User) (models.User, error)
	GetUsers() ([]models.User, error)
	GetUser(id uint) (models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) CreateUser(user *models.User) (models.User, error) {
	return s.repo.Create(*user)
}

func (s *userService) GetUser(id uint) (models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) GetUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.DeleteByID(id)
}
