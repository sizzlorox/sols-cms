package entities

import (
	"encoding/json"
	"time"
)

type Field struct {
	ID            int64           `gorm:"primaryKey" json:"id"`
	ContentTypeID int64           `json:"content_type_id"`
	Name          string          `json:"name"`
	APIName       string          `json:"api_name"`
	Type          string          `json:"type"`
	Required      bool            `json:"required"`
	Unique        bool            `json:"unique"`
	Localized     bool            `json:"localized"`
	DefaultJSON   json.RawMessage `json:"default_json"`
	ConfigJSON    json.RawMessage `json:"config_json"`
	OrderIndex    int             `json:"order_index"`
	IsActive      bool            `json:"is_active"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     *time.Time      `json:"updated_at,omitempty"`
	DeletedAt     *time.Time      `json:"deleted_at,omitempty"`
}
