package models

import (
	"time"
)

type (
	Admin struct {
		Id        uint64    `gorm:"primary_key"`
		Username  string    `gorm:"type:varchar(14);not null"`
		Password  string    `gorm:"not null"`
		Name      string    `gorm:"type:varchar(50);not null"`
		Email     string    `gorm:"type:varchar(100);not null"`
		LastLogin string    `gorm:"null"`
		Token     string    `gorm:"null"`
		CreatedAt time.Time `gorm:"not null"`
		UpdatedAt time.Time `gorm:"null"`
		DeletedAt string    `gorm:"null"`
	}

	AdminLogin struct {
		Username string `json:"username"  binding:"required"`
		Password string `json:"password"  binding:"required"`
	}

	AdminLoginResponse struct {
		Token     string        `json:"token"`
		ExpiredAt time.Duration `json:"expired_at"`
	}
)
