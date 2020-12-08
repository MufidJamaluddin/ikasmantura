package viewmodels

import (
	"backend/utils"
	"time"
)

type ArticleDto struct {
	UserId        uint   `query:"userId,omitempty" json:"userId,omitempty" xml:"userId,omitempty" form:"userId,omitempty"`
	Title         string `query:"title,omitempty" json:"title,omitempty" xml:"title,omitempty" form:"title,omitempty"`
	Id            uint   `query:"-" json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	Body          string `query:"-" json:"body,omitempty" xml:"body,omitempty" form:"body,omitempty"`
	Image         string `query:"-" json:"image,omitempty" xml:"image,omitempty" form:"-"`
	Thumbnail     string `query:"-" json:"thumbnail,omitempty" xml:"thumbnail,omitempty" form:"thumbnail,omitempty"`
	TopicId       uint   `query:"-" json:"topicId,omitempty" xml:"topicId,omitempty" form:"topicId,omitempty"`
	CurrentUserId uint   `query:"-" json:"-" xml:"-" form:"-"`
	CreatedByName string `query:"-" json:"createdByName" xml:"createdByName" form:"-"`
	utils.Created
	utils.Updated
}

func (p *ArticleDto) GetId() uint {
	return p.Id
}

type ArticleParam struct {
	utils.GetParams
	ArticleDto
	StartFrom *time.Time `json:"createdAt_gte" xml:"createdAt_gte" form:"createdAt_gte"`
	EndTo     *time.Time `json:"createdAt_lte" xml:"createdAt_lte" form:"createdAt_lte"`
}

func (p *ArticleParam) GetModel() interface{} {
	return p
}

type ArticleTopicDto struct {
	Name        string `query:"name,omitempty" json:"name,omitempty" xml:"name,omitempty" form:"title,omitempty"`
	Id          uint   `query:"-" json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	Icon        string `query:"-" json:"icon,omitempty" xml:"icon,omitempty" form:"icon,omitempty"`
	Description string `query:"-" json:"description,omitempty" xml:"description,omitempty" form:"description,omitempty"`
	utils.Created
	utils.Updated
}

func (p *ArticleTopicDto) GetId() uint {
	return p.Id
}

type ArticleTopicParam struct {
	utils.GetParams
	ArticleTopicDto
}

func (p *ArticleTopicParam) GetModel() interface{} {
	return p
}
