package viewmodels

import "backend/utils"

type UserDto struct {
	Id           int            `query:"id,omitempty" json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	Name         string         `query:"name,omitempty" json:"name,omitempty" xml:"name,omitempty" form:"name,omitempty"`
	Username     string         `query:"username,omitempty" json:"username,omitempty" xml:"username,omitempty" form:"username,omitempty"`
	Role         string         `query:"role" json:"role" xml:"role" form:"role"`
	Email        string         `query:"email,omitempty" json:"email,omitempty" xml:"email,omitempty" form:"email,omitempty"`
	EmailValid   bool		    `query:"emailValid,omitempty" json:"emailValid,omitempty" xml:"emailValid,omitempty" form:"emailValid,omitempty"`
	Phone        string         `query:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty" form:"phone,omitempty"`
	ForceYear    string         `query:"forceYear,omitempty" json:"forceYear,omitempty" xml:"forceYear,omitempty" form:"forceYear,omitempty"`
	Job          string         `query:"job,omitempty" json:"job,omitempty" xml:"job,omitempty" form:"job,omitempty"`
	JobDesc      string         `query:"jobDesc,omitempty" json:"jobDesc,omitempty" xml:"jobDesc,omitempty" form:"jobDesc,omitempty"`
	Password     string         `query:"-" json:"password,omitempty" xml:"password,omitempty" form:"password,omitempty"`
	RefreshToken string         `query:"-" json:"-" xml:"-" form:"-"`
	ConfirmEmailToken string    `query:"-" json:"-" xml:"-" form:"-"`
	Address      UserAddressDto `query:"-" json:"address,omitempty" xml:"address,omitempty" form:"address,omitempty"`
	Classrooms   []int 		    `query:"-" json:"classrooms,omitempty" xml:"classrooms,omitempty" form:"classrooms,omitempty"`
	utils.Created
	utils.Updated
}

type UserAddressDto struct {
	ID      int    `query:"-" json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	Street  string `query:"-" json:"street,omitempty" xml:"street,omitempty" form:"street,omitempty"`
	Suite   string `query:"-" json:"suite,omitempty" xml:"suite,omitempty" form:"suite,omitempty"`
	City    string `query:"-" json:"city,omitempty" xml:"city,omitempty" form:"city,omitempty"`
	Zipcode string `query:"-" json:"zipcode,omitempty" xml:"zipcode,omitempty" form:"zipcode,omitempty"`
	State   string `query:"-" json:"state,omitempty" xml:"state,omitempty" form:"state,omitempty"`
}

type UserAvailabilityResponseDto struct {
	Exist *bool `query:"-" json:"exist,omitempty" xml:"exist,omitempty" form:"exist,omitempty"`
}

type UserAvailabilityDto struct {
	Username string `query:"username,omitempty" json:"username,omitempty" xml:"username,omitempty" form:"username,omitempty"`
	Email    string `query:"email,omitempty" json:"email,omitempty" xml:"email,omitempty" form:"email,omitempty"`
}

func (p *UserDto) GetId() uint {
	return uint(p.Id)
}

type UserParam struct {
	utils.GetParams
	UserDto
}

func (p *UserParam) GetModel() interface{} {
	return p
}
