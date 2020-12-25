package cati

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Cat struct
type Cat struct {
	ID        string `gorm:"column:id;primary_key" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	IsDeleted bool   `gorm:"default:false;column:isDeleted" json:"isDeleted"`
	CreatedAt int64  `gorm:"column:createdAt;autoCreateTime:milli" json:"createdAt"`
	UpdatedAt int64  `gorm:"column:updatedAt;autoUpdateTime:milli" json:"updatedAt"`
}

// BeforeCreate func
func (cat *Cat) BeforeCreate(db *gorm.DB) error {
	cat.ID = uuid.New().String()
	return nil
}

// TableName func
func (cat *Cat) TableName() string {
	return "CattoHouse"
}
