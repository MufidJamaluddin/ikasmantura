package dto

import (
	"backend/utils"
)

type AboutDto struct {
	Id          uint   `query:"-" json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	Description string `query:"description,omitempty" json:"description,omitempty" xml:"description,omitempty" form:"description,omitempty"`
	Vision      string `query:"vision,omitempty" json:"vision,omitempty" xml:"vision,omitempty" form:"vision,omitempty"`
	Mission     string `query:"mission,omitempty" json:"mission,omitempty" xml:"mission,omitempty" form:"mission,omitempty"`
	utils.Created
	utils.Updated
}

func (p *AboutDto) GetId() uint {
	return p.Id
}

type AlbumParam struct {
	utils.GetParams
	AlbumDto
}

func (p *AlbumParam) GetModel() interface{} {
	return p
}
