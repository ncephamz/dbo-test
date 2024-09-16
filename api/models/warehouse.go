package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	Warehouse struct {
		gorm.Model
		Id        uint64    `gorm:"primary_key"`
		Status    string    `gorm:"type:varchar(14);not null"`
		Name      string    `gorm:"type:varchar(225);not null"`
		CreatedAt time.Time `gorm:"not null"`
		UpdatedAt time.Time `gorm:"null"`
	}
)
