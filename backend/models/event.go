package models

import (
	"backend/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Event struct {
	ID           utils.UUID `gorm:"type:binary(16);primaryKey"`
	Organizer    string     `gorm:"size:35"`
	Title        string     `gorm:"size:35"`
	Description  string     `gorm:"size:256"`
	Image        string     `gorm:"size:100"`
	Thumbnail    string     `gorm:"size:100"`
	Start        time.Time
	End          time.Time
	Participants []UserEvent `gorm:"foreignKey:EventId"`
	utils.Created
	utils.Updated
}

func (base *Event) BeforeCreate(scope *gorm.DB) (err error) {
	base.ID = utils.UUID(uuid.NewV1())
	return
}

func (Event) CreateHistory() interface{} {
	return &EventHistory{}
}

type UserEvent struct {
	ID      uint `gorm:"primaryKey"`
	UserId  uint
	EventId utils.UUID `gorm:"type:binary(16)"`
	User    User  `gorm:"foreignKey:UserId"`
	Event   Event `gorm:"foreignKey:EventId"`
	utils.Created
}

type EventHistory struct {
	utils.History
	ID          uint   `gorm:"primaryKey"`
	Organizer   string `gorm:"size:35"`
	Title       string `gorm:"size:35"`
	Description string `gorm:"size:256"`
	Image       string `gorm:"size:100"`
	Start       time.Time
	End         time.Time
	utils.Updated
}
