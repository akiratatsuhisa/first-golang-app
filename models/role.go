package models

type Role struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Name      string     `json:"name" gorm:"size:255;not null"`
	UserRoles []UserRole `json:"userRoles"`
}
