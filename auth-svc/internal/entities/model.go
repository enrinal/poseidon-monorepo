package entities

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        string         `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" sql:"index"`
}
