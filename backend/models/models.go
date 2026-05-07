package models

import (
	"time"

	"gorm.io/gorm"
)

type Meme struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Author    string         `json:"author"`
	Content   string         `json:"content"`
}
