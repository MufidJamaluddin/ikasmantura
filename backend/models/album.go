package models

import (
	"backend/utils"
)

type Album struct {
	ID     uint         `gorm:"primaryKey"`
	Title  string       `gorm:"size:35"`
	Photos []AlbumPhoto `gorm:"foreignKey:AlbumId"`
	utils.Created
	utils.Updated
}

func (Album) CreateHistory() interface{} {
	return &AlbumHistory{}
}

type AlbumPhoto struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:35"`
	Image     string `gorm:"size:100"`
	Thumbnail string `gorm:"size:100"`
	AlbumId   uint
	Album     Album `gorm:"foreignKey:AlbumId"`
	utils.Created
	utils.Updated
}

type AlbumHistory struct {
	utils.History
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"size:35"`
	utils.Updated
}
