package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	Orders struct {
		gorm.Model
		Id                uint64    `gorm:"primary_key"`
		CustomerAddressId uint64    `gorm:"foreignkey;not null"`
		CustomerId        uint64    `gorm:"foreignkey;not null"`
		Status            string    `gorm:"type:varchar(25);not null"`
		CreatedAt         time.Time `gorm:"not null"`
		UpdatedAt         time.Time `gorm:"null"`
	}
)
