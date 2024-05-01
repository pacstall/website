package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"pacstall.dev/webserver/config"
)

var connectionString = ""

func Connect(config config.DatabaseConfiguration) *gorm.DB {
	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	if connectionString == "" {
		panic("database connection has not been initialized")
	}

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	defer postConnect(db)

	return db
}

func postConnect(database *gorm.DB) {
	database.AutoMigrate(&ShortenedLink{})
}
