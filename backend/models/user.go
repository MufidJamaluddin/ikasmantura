package models

import (
	"backend/utils"
)

type User struct {
	ID         uint   `gorm:"primarykey"`
	Name       string `gorm:"size:53"`
	Username   string `gorm:"size:35,unique"`
	Email      string `gorm:"size:286"`
	Password   string `gorm:"size:40"`
	Phone      string `gorm:"size:13"`
	IsAdmin    bool
	Address    UserAddress     `gorm:"foreignKey:UserId"`
	MyEvents   []UserEvent     `gorm:"foreignKey:UserId"`
	Classrooms []UserClassroom `gorm:"foreignKey:UserId"`
	utils.Created
	utils.Updated
}

type UserAddress struct {
	ID      uint `gorm:"primarykey"`
	UserId  uint
	Street  string
	Suite   string `gorm:"size:35"`
	City    string `gorm:"size:35"`
	Zipcode string `gorm:"size:4"`
}

type UserClassroom struct {
	UserId      uint
	ClassroomId uint8
	Classroom   Classroom `gorm:"foreignKey:ClassroomId"`
	User        User      `gorm:"foreignKey:UserId"`
}

func (User) CreateHistory() interface{} {
	return &UserHistory{}
}

type UserHistory struct {
	utils.History
	ID       uint   `gorm:"primarykey"`
	Name     string `gorm:"size:53"`
	Username string `gorm:"size:35"`
	Email    string `gorm:"size:286"`
	Password string `gorm:"size:40"`
	Phone    string `gorm:"size:13"`
	utils.Updated
}

func (UserAddress) CreateHistory() interface{} {
	return &UserAddressHistory{}
}

type UserAddressHistory struct {
	utils.History
	ID      uint `gorm:"primarykey"`
	Street  string
	Suite   string
	City    string
	Zipcode string
	utils.Updated
}
