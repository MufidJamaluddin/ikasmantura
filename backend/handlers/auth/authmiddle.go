package auth

import (
	"backend/viewmodels"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	history "github.com/vcraescu/gorm-history/v2"
	"gorm.io/gorm"
	"strconv"
)

func AuthenticationHandler(c *fiber.Ctx, db *gorm.DB) error {
	var (
		currentUserId int
		userdata      *viewmodels.AuthorizationModel
		claims        jwt.MapClaims
	)

	if user := c.Locals("user").(*jwt.Token); user != nil {
		claims = user.Claims.(jwt.MapClaims)
		currentUserId = int(claims["id"].(float64))

		userdata = &viewmodels.AuthorizationModel{
			ID:       uint(currentUserId),
			Username: claims["name"].(string),
			Email:    claims["email"].(string),
			IsAdmin:  claims["admin"].(bool),
		}

		c.Locals("db", history.SetUser(db, history.User{
			ID:    strconv.Itoa(currentUserId),
			Email: claims["email"].(string),
		}))

		c.Locals("user", userdata)
	} else {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.Next()
}
