package viewmodels

import "backend/utils"

type ClassroomDto struct {
	Id    uint   `query:"-" json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	Major string `query:"major,omitempty" json:"major,omitempty" xml:"major,omitempty" form:"major,omitempty"`
	Level string `query:"level,omitempty" json:"level,omitempty" xml:"level,omitempty" form:"level,omitempty"`
	Seq   uint8  `query:"seq,omitempty" json:"seq,omitempty" xml:"seq,omitempty" form:"seq,omitempty"`
	utils.Created
	utils.Updated
}

func (p *ClassroomDto) GetId() uint {
	return uint(p.Id)
}

type ClassroomParam struct {
	utils.GetParams
	ClassroomDto
}

func (p *ClassroomParam) GetModel() interface{} {
	return p
}
