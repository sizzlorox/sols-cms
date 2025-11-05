package entities

import "time"

type User struct {
	ID           int64      `gorm:"primaryKey" json:"id"`
	Name         string     `json:"name"`
	Email        string     `gorm:"uniqueIndex" json:"email"`
	PasswordHash string     `json:"-"`
	RefreshToken string     `json:"-"`
	IsActive     bool       `json:"is_active"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
