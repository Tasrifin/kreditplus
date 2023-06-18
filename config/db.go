package config

import (
	"fmt"

	"github.com/Tasrifin/kreditplus/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 8889
	DB_USER     = "user"
	DB_PASSWORD = "password"
	DB_NAME     = "db_kreditplus"
	APP_PORT    = ":7777"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(models.Customer{}, models.Limit{}, models.Transaction{})

	return db

}
