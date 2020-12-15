package gorm

import (
	"log"
	"os"

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
	connect, err := gorm.Open(mysql.Open(os.Getenv("DB_ADDRESS")), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	connect.AutoMigrate(&Cat{})

	return connect
}
