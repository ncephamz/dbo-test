package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	OrdersDetails struct {
		gorm.Model
		Id               uint64    `gorm:"primary_key"`
		OrderId          uint64    `gorm:"foreignkey;not null"`
		StoreWarehouseId uint64    `gorm:"foreignkey;not null"`
		Qty              int       `gorm:"type:int;not null"`
		CreatedAt        time.Time `gorm:"not null"`
		UpdatedAt        time.Time `gorm:"null"`
	}
)
