package models

import (
	"time"
)

type Todo struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" gorm:"size:255;not null"`
	Content   string    `json:"content" gorm:"type:text"`
	Status    string    `json:"status" gorm:"type:enum('new', 'read', 'completed', 'cancelled');default:'new';not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime;not null"`
}
