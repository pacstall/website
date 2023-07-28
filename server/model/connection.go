package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"pacstall.dev/webserver/config"
)

var database *gorm.DB = nil
var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	config.Database.User,
	config.Database.Password,
	config.Database.Host,
	config.Database.Port,
	config.Database.Name,
)

func Instance() *gorm.DB {
	if database != nil {
		return database
	}

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	defer postConnect()

	database = db
	return database
}

func postConnect() {
	database.AutoMigrate(&ShortenedLink{})
}
