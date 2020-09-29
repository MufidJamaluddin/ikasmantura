package models

import (
	"backend/utils"
)

type About struct {
	ID          uint `gorm:"primarykey"`
	Description string
	Vision      string
	Mission     string
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
