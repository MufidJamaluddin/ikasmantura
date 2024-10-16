package auth

import (
	authService "backend/services/auth"
	"backend/viewmodels"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

// @author Mufid Jamaluddin

// GetLoggedInUser godoc
// @Tags Authentication & Authorization
// @Summary GetLoggedInUser
// @Description GetLoggedInUser for get IKA SMAN Situraja users
// @Accept json
// @Produce json
// @Success 200 {object} viewmodels.AuthorizationModel
// @Failure 401 {object} string
// @Router /api/v1/auth [get]
func GetLoggedInUser(c *fiber.Ctx) error {
	var (
		err      error
		ok       bool
		authData *viewmodels.AuthorizationModel
	)

	if authData, ok = viewmodels.GetAuthorizationData(c); !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	err = c.JSON(authData)

	return err
}

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
		loginData viewmodels.LoginDto
		db        *gorm.DB
		ok        bool
		token     *string
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

	if token, err = DoLogin(c, &loginData.Data, &loginData.Expired); err != nil {
		return err
	}

	loginData.Token = *token
	loginData.RefreshToken = loginData.Data.RefreshToken

	err = c.JSON(&loginData.LoginResponseDto)
	return err
}

func RefreshLogin(c *fiber.Ctx) error {
	var (
		err           error
		db            *gorm.DB
		ok            bool
		token         *string
		refreshToken  string
		authData      *viewmodels.AuthorizationModel
		userData      viewmodels.UserDto
		responseLogin viewmodels.LoginResponseDto
		expired       *time.Time
	)

	if authData, ok = viewmodels.GetAuthorizationData(c); !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	refreshToken = c.Cookies(os.Getenv("COOKIE_REFRESH_TOKEN"))
	if refreshToken == "" {
		refreshToken = c.Get(os.Getenv("HEADER_REFRESH_TOKEN"))
	}

	userData.Id = int(authData.ID)
	userData.RefreshToken = refreshToken

	if refreshToken != "" && authData.Role != "" {

		if authData.ID == 0 {
			// c.ClearCookie(os.Getenv("COOKIE_TOKEN"))
			// c.ClearCookie(os.Getenv("COOKIE_REFRESH_TOKEN"))
			return c.SendStatus(fiber.StatusForbidden)
		}

		if err = authService.RefreshToken(db, &userData); err != nil {
			return err
		}

		if userData.Username != authData.Username &&
			userData.Role != authData.Role {
			// c.ClearCookie(os.Getenv("COOKIE_TOKEN"))
			// c.ClearCookie(os.Getenv("COOKIE_REFRESH_TOKEN"))
			log.Printf("User ID %s (username %s != %s)",
				userData.Id, authData.Username, userData.Username)
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	} else {
		userData.Role = authData.Role
		userData.Username = authData.Username
		userData.Email = authData.Email
		if refreshToken == "" {
			userData.RefreshToken = uuid.NewV4().String()
		}
	}

	if token, err = DoLogin(c, &userData, expired); err != nil {
		return err
	}

	responseLogin.Token = *token
	responseLogin.RefreshToken = userData.RefreshToken

	err = c.JSON(&responseLogin)
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

	c.Cookie(&fiber.Cookie{
		Name:     os.Getenv("COOKIE_REFRESH_TOKEN"),
		Value:    "",
		HTTPOnly: true,
		SameSite: "strict",
	})

	c.Cookie(&fiber.Cookie{
		Name:     os.Getenv("COOKIE_TOKEN"),
		Value:    "",
		HTTPOnly: true,
		SameSite: "strict",
	})

	// c.ClearCookie(os.Getenv("COOKIE_REFRESH_TOKEN"))
	// c.ClearCookie(os.Getenv("COOKIE_TOKEN"))

	err = c.SendStatus(fiber.StatusOK)

	return err
}
