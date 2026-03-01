package models

import "time"

type Chat struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Role      string
	Message   string
	CreatedAt time.Time
}