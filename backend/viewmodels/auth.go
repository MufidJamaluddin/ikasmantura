package viewmodels

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type AuthorizationModel struct {
	ID       uint
	Username string
	Email    string
	IsAdmin  bool
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
	Username string `query:"username,omitempty" json:"username,omitempty" xml:"username,omitempty" form:"username,omitempty"`
	Password string `query:"password,omitempty" json:"password,omitempty" xml:"password,omitempty" form:"password,omitempty"`
}

type LoginResponseDto struct {
	Token   string    `query:"-" json:"-" xml:"-" form:"-"`
	Expired time.Time `query:"-" json:"-" xml:"-" form:"-"`
}

type LoginDto struct {
	LoginRequestDto
	LoginResponseDto
	Data UserDto `query:"-" json:"data" xml:"data" form:"-"`
}
