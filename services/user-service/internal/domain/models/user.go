package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" size:"55" notnull:"true"`
	Fmaily   string `json:"family" size:"55" notnull:"true"`
	Mobile   string `json:"mobile" size:"11" notnull:"true"`
	Email    string `json:"email" size:"55" notnull:"true"`
	Username string `json:"username" size:"55" notnull:"true"`
	Password string `json:"password" size:"55" notnull:"true"`
	RoleID   uint   `json:"role_id" notnull:"true"`
}
