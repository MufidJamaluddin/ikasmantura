package dto

import "time"

type LoginDto struct {
	Username string    `query:"username,omitempty" json:"username,omitempty" xml:"username,omitempty" form:"username,omitempty"`
	Password string    `query:"password,omitempty" json:"password,omitempty" xml:"password,omitempty" form:"password,omitempty"`
	Token    string    `query:"-" json:"-" xml:"-" form:"-"`
	Data     UserDto   `query:"-" json:"data" xml:"data" form:"-"`
	Expired  time.Time `query:"-" json:"-" xml:"-" form:"-"`
}
