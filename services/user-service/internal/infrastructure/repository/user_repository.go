package repository

import (
	"github.com/amanc1361/bilan-rekar/user-service/internal/domain/models"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GormUserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *GormUserRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *GormUserRepository) DeleteUser(id int) error {
	return r.db.Delete(&models.User{}, id).Error
}
func (r *GormUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("user_name = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *GormUserRepository) GetUserByEmailAndPassword(email, password string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *GormUserRepository) GetUserByMobile(mobile string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("mobile = ?", mobile).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
