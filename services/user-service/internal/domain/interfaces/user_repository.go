package interfaces

import "github.com/amanc1361/bilan-rekar/user-service/internal/domain/models"

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmailAndPassword(email, password string) (*models.User, error)
	GetUserByMobile(mobile string) (*models.User, error)
}
