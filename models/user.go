package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Mobile    string
	CreatedAt time.Time
}