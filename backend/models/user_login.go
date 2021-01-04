package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type UserLogin struct {
	UserId      uint   	   `gorm:"primaryKey;autoIncrement:false"`
	RefreshToken UUID 	   `gorm:"type:binary(16);primaryKey"`
	RemoteIP 	string 	   `gorm:"size:45"`
	OSName 		string 	   `gorm:"size:35"`
	OSVersion 	string 	   `gorm:"size:10"`
	Device 		string 	   `gorm:"size:35"`
	DeviceType	deviceType `sql:"type:deviceType"`
	CreatedAt   time.Time
}

func (base *UserLogin) BeforeCreate(scope *gorm.DB) (err error) {
	base.RefreshToken = UUID(uuid.NewV1())
	return
}
