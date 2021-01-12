package viewmodels

import (
	"backend/utils"
	"time"
)

type EventDto struct {
	UserId               int          `query:"userId,omitempty" json:"userId,omitempty" xml:"userId,omitempty" form:"userId,omitempty"`
	Organizer            string       `query:"organizer,omitempty" json:"organizer,omitempty" xml:"organizer,omitempty" form:"organizer,omitempty"`
	Title                string       `query:"title,omitempty" json:"title,omitempty" xml:"title,omitempty" form:"title,omitempty"`
	IsMyEvent            bool         `query:"myEvent,omitempty" json:"myEvent,omitempty" xml:"myEvent,omitempty" form:"myEvent,omitempty"`
	Id                   string       `query:"-" json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	Description          string       `query:"-" json:"description,omitempty" xml:"description,omitempty" form:"description,omitempty"`
	Image                string       `query:"-" json:"image,omitempty" xml:"image,omitempty" form:"-"`
	Thumbnail            string       `query:"-" json:"thumbnail,omitempty" xml:"thumbnail,omitempty" form:"-"`
	Start                time.Time    `query:"-" json:"start,omitempty" xml:"start,omitempty" form:"start,omitempty"`
	End                  time.Time    `query:"-" json:"end,omitempty" xml:"end,omitempty" form:"end,omitempty"`
	CurrentUserId        int          `query:"-" json:"-" xml:"-" form:"-"`
	CurrentUserRegisData UserEventDto `query:"-" json:"registration" xml:"registration" form:"-"`
	CreatedByName        string       `query:"-" json:"createdByName,omitempty" xml:"createdByName,omitempty" form:"-"`
	utils.Created
	utils.Updated
}

type UserEventDto struct {
	ID      int `gorm:"primarykey"`
	UserId  int
	EventId int
	utils.Created
}

type EventParam struct {
	utils.GetParams
	EventDto
	StartFrom *time.Time `json:"start_gte" xml:"start_gte" form:"start_gte"`
	EndTo     *time.Time `json:"start_lte" xml:"start_lte" form:"start_lte"`
}

func (p *EventParam) GetModel() interface{} {
	return p
}

type UserEventDetailDto struct {
	UserId       int
	UserFullName string
	UserEmail    string
	EventId      int
	Organizer    string
	EventName    string
	Description  string
	TicketId     int
	StartStr     string
	EndStr       string
}
