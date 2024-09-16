package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	WarehousesCoverages struct {
		gorm.Model
		Id          uint64    `gorm:"primary_key"`
		WarehouseId uint64    `gorm:"foreignkey;not null"`
		Province    string    `gorm:"type:varchar(25);not null"`
		City        string    `gorm:"type:varchar(50);not null"`
		District    string    `gorm:"type:varchar(50);not null"`
		SubDistrict string    `gorm:"type:varchar(50);not null"`
		Zipcode     string    `gorm:"type:varchar(6);not null"`
		Tax         float32   `gorm:"type:numeric(10,3);not null"`
		DeliveryFee float32   `gorm:"type:numeric(10,3);not null"`
		ServiceFee  float32   `gorm:"type:numeric(10,3);not null"`
		CreatedAt   time.Time `gorm:"not null"`
		UpdatedAt   time.Time `gorm:"null"`
	}
)
