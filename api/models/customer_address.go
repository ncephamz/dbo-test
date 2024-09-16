package models

import (
	"time"

	utils "github.com/ncephamz/dbo-test/api/pkg"
)

type (
	CustomerAddress struct {
		Id          uint64    `gorm:"primary_key"`
		CustomerId  uint64    `gorm:"foreignkey;not null"`
		Province    string    `gorm:"type:varchar(25);not null"`
		City        string    `gorm:"type:varchar(50);not null"`
		District    string    `gorm:"type:varchar(50);not null"`
		SubDistrict string    `gorm:"type:varchar(50);not null"`
		Zipcode     string    `gorm:"type:varchar(6);not null"`
		Address     string    `gorm:"type:text;not null"`
		Note        string    `gorm:"type:text;null"`
		GoogleMap   string    `gorm:"type:text;null"`
		IsMain      bool      `gorm:"type:boolean;not null"`
		CreatedAt   time.Time `gorm:"not null"`
		UpdatedAt   time.Time `gorm:"null"`
	}

	RequestCreateCustomerAddress struct {
		Id          string `json:"id"`
		Province    string `json:"province" binding:"required"`
		City        string `json:"city" binding:"required"`
		District    string `json:"district" binding:"required"`
		SubDistrict string `json:"sub_district" binding:"required"`
		Zipcode     string `json:"zipcode" binding:"required"`
		Address     string `json:"address" binding:"required"`
		Note        string `json:"note"`
		GoogleMap   string `json:"google_map"`
		IsMain      bool   `json:"is_main" binding:"required"`
	}
)

func (CustomerAddress) TableName() string {
	return "customers_addresses"
}

func (req RequestCreateCustomerAddress) ToModel(customerId uint64) CustomerAddress {
	var (
		now   = time.Now()
		newId = utils.GenerateID()
	)

	if req.Id != "" {
		newId = utils.StringToUint64(req.Id)
	}

	return CustomerAddress{
		Id:          newId,
		CustomerId:  customerId,
		Province:    req.Province,
		City:        req.City,
		District:    req.District,
		SubDistrict: req.SubDistrict,
		Zipcode:     req.Zipcode,
		Address:     req.Address,
		Note:        req.Note,
		GoogleMap:   req.GoogleMap,
		IsMain:      req.IsMain,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
