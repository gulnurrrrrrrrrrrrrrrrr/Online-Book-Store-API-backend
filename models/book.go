package models

import "time"

type Book struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `gorm:"size:255;not null" json:"title"`
	AuthorID   uint      `gorm:"not null" json:"author_id"`
	CategoryID uint      `gorm:"not null" json:"category_id"`
	Price      float64   `gorm:"not null" json:"price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
