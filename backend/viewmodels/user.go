package viewmodels

import "backend/utils"

type UserDto struct {
	Id        uint           `query:"-" json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	Name      string         `query:"name,omitempty" json:"name,omitempty" xml:"name,omitempty" form:"name,omitempty"`
	Username  string         `query:"username,omitempty" json:"username,omitempty" xml:"username,omitempty" form:"username,omitempty"`
	IsAdmin   bool           `query:"isAdmin" json:"isAdmin" xml:"isAdmin" form:"isAdmin"`
	Email     string         `query:"email,omitempty" json:"email,omitempty" xml:"email,omitempty" form:"email,omitempty"`
	Phone     string         `query:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty" form:"phone,omitempty"`
	ForceYear string         `query:"forceYear,omitempty" json:"forceYear,omitempty" xml:"forceYear,omitempty" form:"forceYear,omitempty"`
	Password  string         `query:"-" json:"password,omitempty" xml:"password,omitempty" form:"password,omitempty"`
	Address   UserAddressDto `query:"-" json:"address,omitempty" xml:"address,omitempty" form:"address,omitempty"`
	utils.Created
	utils.Updated
}

type UserAddressDto struct {
	ID      uint   `query:"-" json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	Street  string `query:"-" json:"street,omitempty" xml:"street,omitempty" form:"street,omitempty"`
	Suite   string `query:"-" json:"suite,omitempty" xml:"suite,omitempty" form:"suite,omitempty"`
	City    string `query:"-" json:"city,omitempty" xml:"city,omitempty" form:"city,omitempty"`
	Zipcode string `query:"-" json:"zipcode,omitempty" xml:"zipcode,omitempty" form:"zipcode,omitempty"`
}

func (p *UserDto) GetId() uint {
	return p.Id
}

type UserParam struct {
	utils.GetParams
	UserDto
}

func (p *UserParam) GetModel() interface{} {
	return p
}
