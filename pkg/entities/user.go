package entities

import "time"

type User struct {
	ID           int64      `gorm:"primaryKey" json:"id"`
	Name         string     `json:"name"`
	Email        string     `gorm:"uniqueIndex" json:"email"`
	PasswordHash string     `json:"-"`
	RefreshToken string     `json:"-"`
	IsSuperAdmin bool       `json:"is_super_admin"`
	IsActive     bool       `json:"is_active"`
	LastLoginAt  *time.Time `json:"last_login_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`

	Memberships []Membership `gorm:"foreignKey:UserID" json:"memberships,omitempty"`
}

type Permission struct {
	ID        int64      `gorm:"primaryKey" json:"id"`
	Key       string     `gorm:"uniqueIndex" json:"key"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type RolePermission struct {
	ID           int64      `gorm:"primaryKey" json:"id"`
	RoleID       int64      `gorm:"index;not null" json:"role_id"`
	PermissionID int64      `gorm:"index;not null" json:"permission_id"`
	IsActive     bool       `json:"is_active"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`

	Role       Role       `gorm:"constraint:OnDelete:CASCADE"`
	Permission Permission `gorm:"constraint:OnDelete:CASCADE"`
}

type Role struct {
	ID        int64      `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"uniqueIndex" json:"name"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Permissions []Permission `gorm:"many2many:role_permissions;joinForeignKey:RoleID;JoinReferences:PermissionID" json:"permissions,omitempty"`
}

type Membership struct {
	ID        int64      `gorm:"primaryKey" json:"id"`
	UserID    int64      `gorm:"index;not null" json:"user_id"`
	DomainID  int64      `gorm:"index;not null" json:"domain_id"`
	RoleID    int64      `gorm:"index;not null" json:"role_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Role   Role
	User   User
	Domain Domain
}

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
