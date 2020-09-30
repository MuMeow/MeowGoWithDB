package gorm

import (
	"log"

	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Cat struct
type Cat struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// BeforeCreate func
func (cat *Cat) BeforeCreate(db *gorm.DB) error {
	db.Statement.SetColumn("ID", uuid.NewV4().String())
	return nil
}

// TableName func
func (cat *Cat) TableName() string {
	return "CattoHouse"
}

// ConnectDB func
func ConnectDB() (connect *gorm.DB) {
	connect, err := gorm.Open(mysql.Open("root:mostori1234@tcp(192.168.1.10:3306)/meow_test?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	connect.AutoMigrate(&Cat{})

	return connect
}
