package models

import (
	"time"

	utils "github.com/ncephamz/dbo-test/api/pkg"
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

	OrderDetailAssosiationToProducts struct {
		OrdersDetails
		Products StoreWarehouseAssosiationToProduct `gorm:"foreignkey:store_warehouse_id;references:id"`
	}

	ResponseOrderDetail struct {
		ProductName string  `json:"product_name"`
		Image       string  `json:"image"`
		Qty         int     `json:"qty"`
		Price       float32 `json:"price"`
		SubTotal    float32 `json:"sub_total"`
	}
)

func (OrderDetailAssosiationToProducts) TableName() string {
	return "orders_details"
}

func ToModelOrderDetail(orderId, storeWarehouseId uint64, qty int) OrdersDetails {
	return OrdersDetails{
		Id:               utils.GenerateID(),
		OrderId:          orderId,
		StoreWarehouseId: storeWarehouseId,
		Qty:              qty,
		CreatedAt:        time.Now(),
	}
}
