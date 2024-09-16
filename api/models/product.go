package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	Products struct {
		gorm.Model
		Id          uint64    `gorm:"primary_key"`
		StoreId     uint64    `gorm:"foreignkey;not null"`
		ProductName string    `gorm:"type:varchar(225);not null"`
		Image       string    `gorm:"type:varchar(225);null"`
		Uom         string    `gorm:"type:varchar(50);not null"`
		Description string    `gorm:"type:text;not null"`
		CreatedAt   time.Time `gorm:"not null"`
		UpdatedAt   time.Time `gorm:"null"`
	}
)
