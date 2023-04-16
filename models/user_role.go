package models

import (
	"time"
)

type UserRole struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"userId"`
	User      User      `json:"user"`
	RoleID    uint      `json:"roleId"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime;not null"`
}
