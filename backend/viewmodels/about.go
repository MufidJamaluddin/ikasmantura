package viewmodels

import (
	"backend/utils"
)

type AboutDto struct {
	Id          uint   `json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	Title       string `json:"title,omitempty" xml:"title,omitempty" form:"title,omitempty"`
	Description string `json:"description,omitempty" xml:"description,omitempty" form:"description,omitempty"`
	Vision      string `json:"vision,omitempty" xml:"vision,omitempty" form:"vision,omitempty"`
	Mission     string `json:"mission,omitempty" xml:"mission,omitempty" form:"mission,omitempty"`
	Facebook    string `json:"facebook,omitempty" xml:"facebook,omitempty" form:"facebook,omitempty"`
	Twitter     string `json:"twitter,omitempty" xml:"twitter,omitempty" form:"twitter,omitempty"`
	Instagram   string `json:"instagram,omitempty" xml:"instagram,omitempty" form:"instagram,omitempty"`
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
