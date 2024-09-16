package models

import (
	"time"

	utils "github.com/ncephamz/dbo-test/api/pkg"
	"gorm.io/gorm"
)

type (
	Customer struct {
		gorm.Model
		Id           uint64    `gorm:"primary_key"`
		PhoneNumber  string    `gorm:"type:varchar(14);not null"`
		Email        string    `gorm:"type:varchar(100);not null"`
		Password     string    `gorm:"not null"`
		Name         string    `gorm:"type:varchar(50);not null"`
		PhotoProfile string    `gorm:"type:varchar(225);null"`
		CreatedAt    time.Time `gorm:"not null"`
		UpdatedAt    time.Time `gorm:"null"`
	}

	CustomerAssosiationToAddress struct {
		Customer
		Address CustomerAddress `gorm:"foreignkey:customer_id;references:id"`
	}

	RequestCreateCustomer struct {
		PhoneNumber  string                       `json:"phone_number" binding:"required"`
		Email        string                       `json:"email" binding:"required"`
		Password     string                       `json:"password" binding:"required"`
		Name         string                       `json:"name" binding:"required"`
		PhotoProfile string                       `json:"photo_profile"`
		Address      RequestCreateCustomerAddress `json:"address" binding:"required"`
	}

	ResponseGetAllCustomer struct {
		Id           string `json:"id"`
		PhoneNumber  string `json:"phone_number"`
		Email        string `json:"email"`
		Name         string `json:"name"`
		PhotoProfile string `json:"photo_profile"`
	}

	ResponseDetailCustomer struct {
		Id           string                       `json:"id"`
		PhoneNumber  string                       `json:"phone_number"`
		Email        string                       `json:"email"`
		Name         string                       `json:"name"`
		PhotoProfile string                       `json:"photo_profile"`
		CreatedAt    time.Time                    `json:"created_at"`
		Address      RequestCreateCustomerAddress `json:"address"`
	}
)

func (CustomerAssosiationToAddress) TableName() string {
	return "customers"
}

func (c Customer) ToResponse() ResponseGetAllCustomer {
	return ResponseGetAllCustomer{
		Id:           utils.IntToString(c.Id),
		PhoneNumber:  c.PhoneNumber,
		Email:        c.Email,
		Name:         c.Name,
		PhotoProfile: c.PhotoProfile,
	}
}

func (req RequestCreateCustomer) ToModel(id string) Customer {
	var (
		now         = time.Now()
		password, _ = utils.HashPassword(req.Password)
		newId       = utils.GenerateID()
	)

	if id != "" {
		newId = utils.StringToUint64(id)
	}

	return Customer{
		Id:           newId,
		PhoneNumber:  req.PhoneNumber,
		Email:        req.Email,
		Name:         req.Name,
		PhotoProfile: req.PhotoProfile,
		Password:     password,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func (c CustomerAssosiationToAddress) ToResponseDetail() ResponseDetailCustomer {
	address := RequestCreateCustomerAddress{
		Id:          utils.IntToString(c.Address.Id),
		Province:    c.Address.Province,
		City:        c.Address.City,
		District:    c.Address.District,
		SubDistrict: c.Address.SubDistrict,
		Zipcode:     c.Address.Zipcode,
		Address:     c.Address.Address,
		Note:        c.Address.Note,
		GoogleMap:   c.Address.GoogleMap,
		IsMain:      c.Address.IsMain,
	}

	return ResponseDetailCustomer{
		Id:           utils.IntToString(c.Customer.Id),
		PhoneNumber:  c.Customer.PhoneNumber,
		Email:        c.Customer.Email,
		Name:         c.Customer.Name,
		PhotoProfile: c.Customer.PhotoProfile,
		Address:      address,
		CreatedAt:    c.Customer.CreatedAt,
	}
}
