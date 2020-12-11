package auth

import (
	authService "backend/services/auth"
	"backend/utils"
	"backend/viewmodels"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"os"
	"time"
)

// @author Mufid Jamaluddin

// Login godoc
// @Tags Authentication & Authorization
// @Summary Login
// @Description Login to IKA SMAN Situraja
// @Accept  json
// @Produce  json
// @Param q body viewmodels.LoginDto true "Pagination Options"
// @Success 200 {object} viewmodels.LoginDto
// @Failure 400 {object} string
// @Router /api/v1/auth [post]
func Login(c *fiber.Ctx) error {
	var (
		err       error
		tokenizer *jwt.Token
		loginData viewmodels.LoginDto
		expired   time.Time
		db        *gorm.DB
		ok        bool
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err = c.BodyParser(&loginData); err != nil {
		err = c.SendStatus(fiber.StatusUnauthorized)
		return err
	}

	// Throws Unauthorized error
	if err = authService.Login(db, &loginData); err != nil {
		err = c.SendStatus(fiber.StatusUnauthorized)
		return err
	}

	loginData.Password = ""
	loginData.Data.Password = ""

	// Create token
	tokenizer = jwt.New(jwt.SigningMethodHS256)

	// Expiration
	expired = time.Now().Add(time.Hour * 72)

	// Set claims
	claims := tokenizer.Claims.(jwt.MapClaims)
	claims["name"] = loginData.Username
	claims["email"] = loginData.Data.Email
	claims["id"] = loginData.Data.Id
	claims["admin"] = loginData.Data.IsAdmin
	claims["exp"] = expired.Unix()

	// Generate encoded token and send it as response.
	if loginData.Token,
		err = tokenizer.SignedString(utils.ToBytes(os.Getenv("SECRET_KEY"))); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	loginData.Expired = expired

	c.Cookie(&fiber.Cookie{
		Name:     "web_ika_id",
		Value:    loginData.Token,
		Expires:  expired,
		HTTPOnly: true,
		SameSite: "strict",
	})

	err = c.JSON(&loginData)
	return err
}

// Logout godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Authentication & Authorization
// @Summary Logout
// @Description Logout to IKA SMAN Situraja
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /api/v1/auth [delete]
func Logout(c *fiber.Ctx) error {
	var err error

	c.ClearCookie("web_ika_id")
	err = c.SendStatus(fiber.StatusOK)

	return err
}
