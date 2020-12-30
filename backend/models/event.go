package models

import (
	"backend/utils"
	"time"
)

type Event struct {
	ID           uint   `gorm:"primaryKey"`
	Organizer    string `gorm:"size:35"`
	Title        string `gorm:"size:35"`
	Description  string `gorm:"size:256"`
	Image        string `gorm:"size:100"`
	Thumbnail    string `gorm:"size:100"`
	Start        time.Time
	End          time.Time
	Participants []UserEvent `gorm:"foreignKey:EventId"`
	utils.Created
	utils.Updated
}

func (Event) CreateHistory() interface{} {
	return &EventHistory{}
}

type UserEvent struct {
	ID      uint `gorm:"primaryKey"`
	UserId  uint
	EventId uint
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
