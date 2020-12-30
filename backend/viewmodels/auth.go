package viewmodels

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type AuthorizationModel struct {
	ID       uint   `json:"id,omitempty"`
	Username string `json:"name,omitempty"`
	FullName string `json:"fullName,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
	Seq      uint64 `json:"pi,omitempty"`
	Exp      int64  `json:"exp,omitempty"`
}

func GetAuthorizationData(ctx *fiber.Ctx) (*AuthorizationModel, bool) {
	var (
		authData *AuthorizationModel
		ok       bool
		user     interface{}
	)
	if user = ctx.Locals("user"); user == nil {
		return nil, false
	}
	authData, ok = user.(*AuthorizationModel)
	return authData, ok
}

type LoginRequestDto struct {
	Email string `query:"email,omitempty" json:"email,omitempty" xml:"email,omitempty" form:"email,omitempty"`
	Username string `query:"username,omitempty" json:"username,omitempty" xml:"username,omitempty" form:"username,omitempty"`
	Password string `query:"password,omitempty" json:"password,omitempty" xml:"password,omitempty" form:"password,omitempty"`
}

type LoginResponseDto struct {
	Token   	 string    `query:"-" json:"token" xml:"token" form:"-"`
	RefreshToken string	   `query:"-" json:"refreshToken" xml:"refreshToken" form:"-"`
	Expired 	 time.Time `query:"-" json:"-" xml:"-" form:"-"`
}

type LoginDto struct {
	LoginRequestDto
	LoginResponseDto
	Data UserDto `query:"-" json:"-" xml:"-" form:"-"`
}
