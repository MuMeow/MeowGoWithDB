package gorm

import (
	cati "MeowGoWithDB/services/cat/interface"

	"log"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// //Cat struct
// type Cat struct {
// 	ID        string `gorm:"column:id;primary_key" json:"id"`
// 	Name      string `gorm:"column:name" json:"name"`
// 	IsDeleted bool   `gorm:"default:false;column:isDeleted" json:"isDeleted"`
// 	CreatedAt int64  `gorm:"column:createdAt;autoCreateTime:milli" json:"createdAt"`
// 	UpdatedAt int64  `gorm:"column:updatedAt;autoUpdateTime:milli" json:"updatedAt"`
// }

type catInterface cati.Cat

// BeforeCreate func
func (cat *catInterface) BeforeCreate(db *gorm.DB) error {
	cat.ID = uuid.New().String()
	return nil
}

// TableName func
func (cat *catInterface) TableName() string {
	return "CattoHouse"
}

// ConnectDB func
func ConnectDB(address string) (connect *gorm.DB) {
	connect, err := gorm.Open(mysql.Open(address), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	connect.AutoMigrate(&cati.Cat{})

	return connect
}
