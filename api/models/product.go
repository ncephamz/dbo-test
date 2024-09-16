package models

import (
	"time"

	utils "github.com/ncephamz/dbo-test/api/pkg"
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

	ProductsAssosiationsToStoreWarehouse struct {
		Products
		StoreWarehouse StoresWarehouses `gorm:"foreignkey:product_id;references:id"`
	}

	ResponseGetProducts struct {
		Id          string  `json:"id"`
		ProductName string  `json:"product_name"`
		Image       string  `json:"image"`
		Uom         string  `json:"uom"`
		Description string  `json:"description"`
		Qty         int     `json:"qty"`
		Price       float32 `json:"price"`
	}
)

func (ProductsAssosiationsToStoreWarehouse) TableName() string {
	return "products"
}

func (p ProductsAssosiationsToStoreWarehouse) ToResponse() ResponseGetProducts {
	return ResponseGetProducts{
		Id:          utils.IntToString(p.Products.Id),
		ProductName: p.Products.ProductName,
		Image:       p.Products.Image,
		Uom:         p.Products.Uom,
		Description: p.Products.Description,
		Qty:         p.StoreWarehouse.Qty,
		Price:       p.StoreWarehouse.Price,
	}
}
