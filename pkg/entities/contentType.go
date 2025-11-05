package entities

import (
	"encoding/json"
	"time"
)

type ContentType struct {
	ID          int64           `gorm:"primaryKey" json:"id"`
	Name        string          `gorm:"uniqueIndex" json:"name"`
	Slug        string          `json:"slug"`
	Icon        string          `json:"icon"`
	Description string          `json:"description"`
	Options     json.RawMessage `json:"options_json"`
	IsActive    bool            `json:"is_active"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`
	DeletedAt   *time.Time      `json:"deleted_at,omitempty"`
}
