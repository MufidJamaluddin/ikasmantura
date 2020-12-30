package models

import (
	"database/sql/driver"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"time"
)

type deviceType string

const (
	mobile = "mobile"
	tablet = "tablet"
	desktop = "desktop"
	bot = "bot"
)

func (d *deviceType) Scan(value interface{}) error {
	*d = deviceType(value.([]byte))
	return nil
}

func (d deviceType) Value() (driver.Value, error) {
	return utils.ToString(d), nil
}

type UserLogin struct {
	UserId      uint   	   `gorm:"primaryKey;autoIncrement:false"`
	Seq 		uint64 	   `gorm:"primaryKey;autoIncrement:true"`
	RefreshToken uuid.UUID `gorm:"type:uuid;"`
	RemoteIP 	string 	   `gorm:"size:45"`
	OSName 		string 	   `gorm:"size:35"`
	OSVersion 	string 	   `gorm:"size:10"`
	Device 		string 	   `gorm:"size:35"`
	DeviceType	deviceType `sql:"type:deviceType"`
	CreatedAt   time.Time
}

func (base *UserLogin) BeforeCreate(scope *gorm.DB) (err error) {
	base.RefreshToken = uuid.NewV4()
	return
}
