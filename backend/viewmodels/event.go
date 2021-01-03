package viewmodels

import (
	"backend/utils"
	"time"
)

type EventDto struct {
	UserId               uint         `query:"userId,omitempty" json:"userId,omitempty" xml:"userId,omitempty" form:"userId,omitempty"`
	Organizer            string       `query:"organizer,omitempty" json:"organizer,omitempty" xml:"organizer,omitempty" form:"organizer,omitempty"`
	Title                string       `query:"title,omitempty" json:"title,omitempty" xml:"title,omitempty" form:"title,omitempty"`
	IsMyEvent            bool         `query:"myEvent,omitempty" json:"myEvent,omitempty" xml:"myEvent,omitempty" form:"myEvent,omitempty"`
	Id                   string       `query:"-" json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	Description          string       `query:"-" json:"description,omitempty" xml:"description,omitempty" form:"description,omitempty"`
	Image                string       `query:"-" json:"image,omitempty" xml:"image,omitempty" form:"-"`
	Thumbnail            string       `query:"-" json:"thumbnail,omitempty" xml:"thumbnail,omitempty" form:"-"`
	Start                time.Time    `query:"-" json:"start,omitempty" xml:"start,omitempty" form:"start,omitempty"`
	End                  time.Time    `query:"-" json:"end,omitempty" xml:"end,omitempty" form:"end,omitempty"`
	CurrentUserId        uint         `query:"-" json:"-" xml:"-" form:"-"`
	CurrentUserRegisData UserEventDto `query:"-" json:"registration" xml:"registration" form:"-"`
	CreatedByName        string       `query:"-" json:"createdByName,omitempty" xml:"createdByName,omitempty" form:"-"`
	utils.Created
	utils.Updated
}

type UserEventDto struct {
	ID      uint `gorm:"primarykey"`
	UserId  uint
	EventId uint
	utils.Created
}

type EventParam struct {
	utils.GetParams
	EventDto
	StartFrom *time.Time `json:"createdAt_gte" xml:"createdAt_gte" form:"createdAt_gte"`
	EndTo     *time.Time `json:"createdAt_lte" xml:"createdAt_lte" form:"createdAt_lte"`
}

func (p *EventParam) GetModel() interface{} {
	return p
}

type UserEventDetailDto struct {
	UserId       uint
	UserFullName string
	UserEmail    string
	EventId      uint
	Organizer    string
	EventName    string
	Description  string
	TicketId     uint
	StartStr     string
	EndStr       string
}
