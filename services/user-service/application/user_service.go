package application

import (
	"github.com/bilan-rekar/services/user-service/internal/domain/interfaces"
	"github.com/bilan-rekar/services/user-service/internal/domain/models"
)

type UserService struct {
	userRepo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *UserService) GetUser(id int) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}
func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepo.CreateUser(user)
}
func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepo.UpdateUser(user)
}
func (s *UserService) DeleteUser(id int) error {
	return s.userRepo.DeleteUser(id)
}
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(email)
}
func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	return s.userRepo.GetUserByUsername(username)
}
func (s *UserService) GetUserByEmailAndPassword(email, password string) (*models.User, error) {
	return s.userRepo.GetUserByEmailAndPassword(email, password)
}
func (s *UserService) GetUserByMobile(mobile string) (*models.User, error) {
	return s.userRepo.GetUserByMobile(mobile)
}
