package database

import (
	"github.com/ncephamz/dbo-test/api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config config.Config) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(config.Database.DSN), &gorm.Config{})
	if err != nil {
		return db, err
	}

	conn, err := db.DB()
	if err != nil {
		return db, err
	}

	conn.SetMaxIdleConns(config.Database.MaxIdleConns)
	conn.SetMaxOpenConns(config.Database.MaxOpenConns)

	return db, nil
}
