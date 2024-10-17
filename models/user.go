package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Username  string         `json:"username" gorm:"size:255;not null"`
	Password  string         `json:"password" gorm:"size:255;not null"`
	TotpKey   sql.NullString `json:"totpKey" gorm:"size:255"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime;not null"`
	UserRoles []UserRole     `json:"userRoles"`
}
