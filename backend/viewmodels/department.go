package viewmodels

import "backend/utils"

type DepartmentDto struct {
	UserId       int    `query:"userId,omitempty" json:"userId,omitempty" xml:"userId,omitempty" form:"userId,omitempty"`
	Name         string `query:"name,omitempty" json:"name,omitempty" xml:"name,omitempty" form:"name,omitempty"`
	Type         int    `query:"type,omitempty" json:"type,omitempty" xml:"type,omitempty" form:"type,omitempty"`
	Id           int    `query:"-" json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	UserFullname string `query:"-" json:"userFullname,omitempty" xml:"userFullname,omitempty" form:"-"`
	utils.Created
	utils.Updated
}

func (p *DepartmentDto) GetId() uint {
	return uint(p.Id)
}

type DepartmentParam struct {
	utils.GetParams
	DepartmentDto
}

func (p *DepartmentParam) GetModel() interface{} {
	return p
}
