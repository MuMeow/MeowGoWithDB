package gorm

import (
	cati "MeowGoWithDB/services/cat/interface"

	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB func
func ConnectDB(address string) (connect *gorm.DB) {
	connect, err := gorm.Open(mysql.Open(address), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	connect.AutoMigrate(&cati.Cat{})

	return connect
}
