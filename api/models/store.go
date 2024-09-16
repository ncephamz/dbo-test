package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	Stores struct {
		gorm.Model
		Id        uint64    `gorm:"primary_key"`
		Level     string    `gorm:"type:varchar(14);not null"`
		Name      string    `gorm:"type:varchar(225);not null"`
		CreatedAt time.Time `gorm:"not null"`
		UpdatedAt time.Time `gorm:"null"`
	}
)
