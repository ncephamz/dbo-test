package models

import (
	"time"

	utils "github.com/ncephamz/dbo-test/api/pkg"
	"gorm.io/gorm"
)

const (
	CART      Status = "CART"
	PAID      Status = "PAID"
	DELIVRY   Status = "DELIVERY"
	COMPLETED Status = "COMPLETED"
)

type (
	Status string
	Orders struct {
		gorm.Model
		Id                uint64    `gorm:"primary_key"`
		CustomerAddressId uint64    `gorm:"foreignkey;not null"`
		CustomerId        uint64    `gorm:"foreignkey;not null"`
		Status            string    `gorm:"type:varchar(25);not null"`
		CreatedAt         time.Time `gorm:"not null"`
		UpdatedAt         time.Time `gorm:"null"`
	}

	OrderAssosiationToCustomer struct {
		Orders
		Customer Customer `gorm:"foreignkey:customer_id;references:id"`
	}

	RequestAddToCart struct {
		CustomerAddressId string `json:"customer_address_id" binding:"required"`
		CustomerId        string `json:"customer_id" binding:"required"`
		ProductId         string `json:"product_id" binding:"required"`
		Qty               int    `json:"qty" binding:"required"`
	}

	ResponseGetOrders struct {
		Id           string `json:"id"`
		Status       string `json:"status"`
		CustomerName string `json:"customer_name"`
	}
)

func (OrderAssosiationToCustomer) TableName() string {
	return "orders"
}

func (req RequestAddToCart) ToModelOrder() Orders {
	return Orders{
		Id:                utils.GenerateID(),
		CustomerAddressId: utils.StringToUint64(req.CustomerAddressId),
		CustomerId:        utils.StringToUint64(req.CustomerId),
		Status:            string(CART),
	}
}

func (o OrderAssosiationToCustomer) ToResponse() ResponseGetOrders {
	return ResponseGetOrders{
		Id:           utils.IntToString(o.Orders.Id),
		Status:       o.Orders.Status,
		CustomerName: o.Customer.Name,
	}
}
