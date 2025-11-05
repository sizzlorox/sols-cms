package entities

import "time"

type AuditLog struct {
	ID         uint           `gorm:"primaryKey"`
	DomainID   *uint          `gorm:"index"`                  // nullable for global events
	UserID     *uint          `gorm:"index"`                  // null for system actions
	Action     string         `gorm:"type:varchar(64);index"` // e.g. "create", "update", "delete", "login"
	EntityType string         `gorm:"type:varchar(64);index"` // e.g. "entry", "content_type", "user"
	EntityID   *uint          `gorm:"index"`
	EntitySlug string         `gorm:"type:varchar(128)"` // optional human-readable ref
	Metadata   datatypes.JSON `gorm:"type:jsonb"`        // flexible extra context
	Diff       datatypes.JSON `gorm:"type:jsonb"`        // what changed
	IPAddress  string         `gorm:"type:varchar(45)"`  // optional
	UserAgent  string
	CreatedAt  time.Time

	Domain *Domain `gorm:"constraint:OnDelete:SET NULL"`
	User   *User   `gorm:"constraint:OnDelete:SET NULL"`
}
