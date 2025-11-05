package entities

import "time"

type Domain struct {
	ID        int64      `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"uniqueIndex" json:"name"`
	Slug      string     `gorm:"uniqueIndex" json:"slug"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type CreateDomainDTO struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Slug     string `json:"slug" validate:"required"`
	IsActive bool   `json:"is_active" default:"true"`
}

type UpdateDomainDTO struct {
	ID   int64   `json:"id"`
	Name *string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Slug *string `json:"slug,omitempty"`
}

func (d *Domain) IsEmpty() bool {
	return Domain{} == *d
}
