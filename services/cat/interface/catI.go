package cati

//Cat struct
type Cat struct {
	ID        string `gorm:"column:id;primary_key" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	IsDeleted bool   `gorm:"default:false;column:isDeleted" json:"isDeleted"`
	CreatedAt int64  `gorm:"column:createdAt;autoCreateTime:milli" json:"createdAt"`
	UpdatedAt int64  `gorm:"column:updatedAt;autoUpdatedTime:milli" json:"updatedAt"`
}
