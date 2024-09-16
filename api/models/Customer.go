package models

import (
	"time"

	utils "github.com/ncephamz/dbo-test/api/pkg"
)

type (
	Customer struct {
		Id           uint64    `gorm:"primary_key"`
		PhoneNumber  string    `gorm:"type:varchar(14);not null"`
		Email        string    `gorm:"type:varchar(100);not null"`
		Password     string    `gorm:"not null"`
		Name         string    `gorm:"type:varchar(50);not null"`
		PhotoProfile string    `gorm:"type:varchar(225);null"`
		CreatedAt    time.Time `gorm:"not null"`
		UpdatedAt    time.Time `gorm:"null"`
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
)

func (c Customer) ToResponse() ResponseGetAllCustomer {
	return ResponseGetAllCustomer{
		Id:           utils.IntToString(c.Id),
		PhoneNumber:  c.PhoneNumber,
		Email:        c.Email,
		Name:         c.Name,
		PhotoProfile: c.PhotoProfile,
	}
}

func (req RequestCreateCustomer) ToModel() Customer {
	now := time.Now()
	password, _ := utils.HashPassword(req.Password)

	return Customer{
		Id:           utils.GenerateID(),
		PhoneNumber:  req.PhoneNumber,
		Email:        req.Email,
		Name:         req.Name,
		PhotoProfile: req.PhotoProfile,
		Password:     password,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}
