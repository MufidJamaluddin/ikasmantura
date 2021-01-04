package models

import (
	"backend/utils"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:53"`
	Username   string `gorm:"size:35,unique"`
	Email      string `gorm:"size:286,unique"`
	Password   string `gorm:"size:40"`
	Phone      string `gorm:"size:13"`
	ForceYear  string `gorm:"size:4"`
	Role       string `gorm:"size:6"`
	RefreshToken uuid.UUID `gorm:"type:binary(16)"`
	Address    UserAddress     `gorm:"foreignKey:UserId"`
	MyEvents   []UserEvent     `gorm:"foreignKey:UserId"`
	Classrooms []UserClassroom `gorm:"foreignKey:UserId"`
	utils.Created
	utils.Updated
}

type UserAddress struct {
	ID      uint `gorm:"primaryKey"`
	UserId  uint
	Street  string
	Suite   string `gorm:"size:35"`
	City    string `gorm:"size:35"`
	Zipcode string `gorm:"size:4"`
	State   string `gorm:"size:35"`
}

type UserClassroom struct {
	UserId      uint
	ClassroomId uint
	Classroom   Classroom `gorm:"foreignKey:ClassroomId"`
	User        User      `gorm:"foreignKey:UserId"`
}

func (User) CreateHistory() interface{} {
	return &UserHistory{}
}

type UserHistory struct {
	utils.History
	ID       uint   `gorm:"primaryKey"`
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
	ID      uint `gorm:"primaryKey"`
	Street  string
	Suite   string
	City    string
	Zipcode string
	utils.Updated
}
