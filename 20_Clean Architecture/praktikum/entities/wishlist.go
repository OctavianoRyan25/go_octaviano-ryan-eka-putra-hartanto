package entities

import (
	"time"

	"gorm.io/gorm"
)

type Wishlist struct {
	ID         int            `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Title      string         `json:"title"`
	IsAchieved bool           `json:"is_achieved"`
}

type RepositoryInterface interface {
	GetAll(wishlist *Wishlist) error
	Create(wishlist *Wishlist) error
}

type UseCaseInterface interface {
	GetAll(wishlist *Wishlist) (Wishlist, error)
	Create(wishlist *Wishlist) (Wishlist, error)
}
