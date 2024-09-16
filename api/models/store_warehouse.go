package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	StoresWarehouses struct {
		gorm.Model
		Id                  uint64    `gorm:"primary_key"`
		WarehouseCoverageId uint64    `gorm:"foreignkey;not null"`
		StoreId             uint64    `gorm:"foreignkey;not null"`
		ProductId           uint64    `gorm:"foreignkey;not null"`
		Qty                 int       `gorm:"type:int;not null"`
		Price               float32   `gorm:"type:numeric(10,3);not null"`
		CreatedAt           time.Time `gorm:"not null"`
		UpdatedAt           time.Time `gorm:"null"`
	}

	StoreWarehouseAssosiationToWareHouseCoverage struct {
		StoresWarehouses
		Coverage WarehousesCoverages `gorm:"foreignkey:warehouse_id;references:id"`
	}

	StoreWarehouseAssosiationToProduct struct {
		StoresWarehouses
		Product  Products            `gorm:"foreignkey:product_id;references:id"`
		Coverage WarehousesCoverages `gorm:"foreignkey:warehouse_id;references:id"`
	}
)

func (StoreWarehouseAssosiationToWareHouseCoverage) TableName() string {
	return "stores_warehouses"
}

func (StoreWarehouseAssosiationToProduct) TableName() string {
	return "stores_warehouses"
}
