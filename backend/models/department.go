package models

import (
	"backend/utils"
)

type Department struct {
	ID     uint   `gorm:"primarykey"`
	Name   string `gorm:"size:35"`
	Type   uint8
	UserId uint
	User   User `gorm:"foreignKey:UserId"`
	utils.Created
	utils.Updated
}

func (Department) CreateHistory() interface{} {
	return &DepartmentHistory{}
}

type DepartmentHistory struct {
	utils.History
	ID     uint `gorm:"primarykey"`
	UserId uint
	Name   string `gorm:"size:35"`
	Type   uint8
	utils.Updated
}
