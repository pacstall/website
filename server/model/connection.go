package model

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/log"
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

	err := retry(5, func() (err error) {
		database, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		return
	})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	log.Info("connected to database.")

	defer postConnect()

	return database
}

func retry(trials int, fn func() error) error {
	var err error
	for i := 0; i < trials; i += 1 {
		if err = fn(); err != nil {
			log.Warn("failed to connect to database. retrying...")
			time.Sleep(3 * time.Second)
		} else {
			return nil
		}
	}

	return err
}

func postConnect() {
	database.AutoMigrate(&ShortenedLink{})
}
