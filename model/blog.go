package model

import "time"

type Blog struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Slug      string    `gorm:"type:varchar(255);not null;unique;" json:"slug"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	IsActive  bool      `gorm:"not null" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
