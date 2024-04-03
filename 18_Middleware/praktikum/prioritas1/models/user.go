package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email"`
	Name      string         `json:"name"`
	Password  string         `json:"password"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeleteAt  gorm.DeletedAt `json:"delete_at"`
}

// type RegisterRequest struct {
// 	Email    string `json:"email"`
// 	Name     string `json:"name"`
// 	Password string `json:"password"`
// }

// type LoginRequest struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// type AuthResponse struct {
// 	Status  string `json:"status"`
// 	Message string `json:"message"`
// 	Data    any    `json:"data, omitempty"`
// }
