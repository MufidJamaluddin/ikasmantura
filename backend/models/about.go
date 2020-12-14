package models

import (
	"backend/utils"
)

type About struct {
	ID          uint   `gorm:"primarykey"`
	Title       string `gorm:"size:35"`
	Description string
	Vision      string
	Mission     string
	Facebook    string `gorm:"size:35"`
	Twitter     string `gorm:"size:35"`
	Instagram   string `gorm:"size:35"`
	utils.Created
	utils.Updated
}

func (About) CreateHistory() interface{} {
	return &AboutHistory{}
}

type AboutHistory struct {
	utils.History
	ID          uint `gorm:"primarykey"`
	Description string
	Vision      string
	Mission     string
	utils.Updated
}
