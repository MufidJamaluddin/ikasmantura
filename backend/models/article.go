package models

import (
	"backend/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Article struct {
	ID             utils.UUID `gorm:"type:binary(16);primaryKey"`
	Title          string     `gorm:"size:35"`
	Body           string     `gorm:"type:TEXT"`
	Image          string     `gorm:"size:100"`
	Thumbnail      string     `gorm:"size:100"`
	ArticleTopicId uint
	ArticleTopic   ArticleTopic `gorm:"foreignKey:ArticleTopicId"`
	utils.Created
	utils.Updated
}

func (base *Article) BeforeCreate(scope *gorm.DB) (err error) {
	base.ID = utils.UUID(uuid.NewV1())
	return
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
