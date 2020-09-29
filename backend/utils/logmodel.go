package utils

import "time"

type Created struct {
	CreatedBy uint      `gorm:"<-:create" query:"-" json:"createdBy,omitempty" xml:"createdBy,omitempty" form:"createdBy,omitempty"`
	CreatedAt time.Time `gorm:"<-:create" query:"-" json:"createdAt,omitempty" xml:"createdAt,omitempty" form:"createdAt,omitempty"`
}

type Updated struct {
	UpdatedBy uint      `query:"-" json:"updatedBy,omitempty" xml:"updatedBy,omitempty" form:"updatedBy,omitempty"`
	UpdatedAt time.Time `query:"-" json:"updatedAt,omitempty" xml:"updatedAt,omitempty" form:"updatedAt,omitempty"`
}

func (p *Created) SetCreated(userId uint, createdAt time.Time) {
	p.CreatedBy = userId
	p.CreatedAt = createdAt
}

func (p *Created) GetCreatedBy() uint {
	return p.CreatedBy
}

func (p *Created) GetCreatedAt() time.Time {
	return p.CreatedAt
}

func (p *Updated) SetUpdated(userId uint, updatedAt time.Time) {
	p.UpdatedBy = userId
	p.UpdatedAt = updatedAt
}

func (p *Updated) GetUpdatedBy() uint {
	return p.UpdatedBy
}

func (p *Updated) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

func FillCreated(from interface{}, to interface{}) {
	in, okIn := from.(ICreated)
	out, okOut := to.(ICreated)

	if okIn && okOut {
		out.SetCreated(in.GetCreatedBy(), in.GetCreatedAt())
	}
}

func FillUpdated(from interface{}, to interface{}) {
	in, okIn := from.(IUpdated)
	out, okOut := to.(IUpdated)

	if okIn && okOut {
		out.SetUpdated(in.GetUpdatedBy(), in.GetUpdatedAt())
	}
}
