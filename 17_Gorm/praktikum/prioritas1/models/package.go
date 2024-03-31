package models

import "time"

type Package struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	Name             string    `json:"name"`
	Sender           string    `json:"sender"`
	Receiver         string    `json:"receiver"`
	SenderLocation   string    `json:"sender_location"`
	ReceiverLocation string    `json:"receiver_location"`
	Fee              int       `json:"fee"`
	Weight           int       `json:"weight"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
type PackageResponse struct {
	Name             string    `json:"name"`
	Sender           string    `json:"sender"`
	Receiver         string    `json:"receiver"`
	SenderLocation   string    `json:"sender_location"`
	ReceiverLocation string    `json:"receiver_location"`
	Fee              int       `json:"fee"`
	Weight           int       `json:"weight"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
