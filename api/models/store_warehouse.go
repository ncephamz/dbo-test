package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	StoresWarehouses struct {
		gorm.Model
		Id          uint64    `gorm:"primary_key"`
		WarehouseId uint64    `gorm:"foreignkey;not null"`
		StoreId     uint64    `gorm:"foreignkey;not null"`
		ProductId   uint64    `gorm:"foreignkey;not null"`
		Qty         int       `gorm:"type:int;not null"`
		Price       float32   `gorm:"type:numeric(10,3);not null"`
		CreatedAt   time.Time `gorm:"not null"`
		UpdatedAt   time.Time `gorm:"null"`
	}
)
