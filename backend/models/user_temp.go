package models

import (
	"backend/utils"
)

type TempUser struct {
	ID           uint                `gorm:"primaryKey"`
	Name         string              `gorm:"size:53"`
	Username     string              `gorm:"size:35,unique"`
	Email        string              `gorm:"size:286,unique"`
	EmailValid   bool
	Password     string              `gorm:"size:40"`
	Phone        string              `gorm:"size:13"`
	ForceYear    string              `gorm:"size:4"`
	Job          string              `gorm:"size:35"`
	JobDesc      string              `gorm:"size:85"`
	RefreshToken utils.UUID          `gorm:"type:binary(16)"`
	ConfirmEmailToken   utils.UUID   `gorm:"type:binary(16)"`
	Address      TempUserAddress     `gorm:"foreignKey:UserId"`
	Classrooms   []TempUserClassroom `gorm:"foreignKey:UserId"`
}

type TempUserAddress struct {
	ID      uint `gorm:"primaryKey"`
	UserId  uint
	Street  string
	Suite   string `gorm:"size:35"`
	City    string `gorm:"size:35"`
	Zipcode string `gorm:"size:4"`
	State   string `gorm:"size:35"`
}

type TempUserClassroom struct {
	UserId      uint
	ClassroomId uint
	Classroom   Classroom `gorm:"foreignKey:ClassroomId"`
	User        TempUser  `gorm:"foreignKey:UserId"`
}
