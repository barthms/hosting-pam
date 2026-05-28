package models

import "time"

type Role struct {
	ID          int32      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `gorm:"type:varchar(50);not null;uniqueIndex" json:"name"`
	Description string     `gorm:"type:text" json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	IsDeleted   *time.Time `json:"is_deleted,omitempty"`
}

func (Role) TableName() string { return "roles" }
