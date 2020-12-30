package models

import "backend/utils"

type Article struct {
	ID             uint   `gorm:"primaryKey"`
	Title          string `gorm:"size:35"`
	Body           string
	Image          string `gorm:"size:100"`
	Thumbnail      string `gorm:"size:100"`
	ArticleTopicId uint
	ArticleTopic   ArticleTopic `gorm:"foreignKey:ArticleTopicId"`
	utils.Created
	utils.Updated
}

type ArticleTopic struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:35"`
	Icon        string    `gorm:"size:35"`
	Description string    `gorm:"size:35"`
	Articles    []Article `gorm:"foreignKey:ArticleTopicId"`
	utils.Created
	utils.Updated
}

func (ArticleTopic) CreateHistory() interface{} {
	return &ArticleTopicHistory{}
}

type ArticleTopicHistory struct {
	utils.History
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:35"`
	utils.Updated
}
