package auth

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	PasswordHash string
	RoleID       uint
	Role         Role
	// Profile fields...
}

type Role struct {
	gorm.Model
	Name string `gorm:"unique"`
}
