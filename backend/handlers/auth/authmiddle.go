package auth

import (
	"backend/utils"
	"backend/viewmodels"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
	history "github.com/vcraescu/gorm-history/v2"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
	"time"
)

func DoLogin(
	c *fiber.Ctx,
	userData *viewmodels.UserDto,
	expired *time.Time,
) (*string, error) {

	var (
		err          error
		token        string
		refreshToken string
		tokenizer    *jwt.Token
	)

	// Create token
	tokenizer = jwt.New(jwt.SigningMethodHS256)

	// Expiration
	*expired = time.Now().Add(5 * 24 * time.Hour)

	// Set claims
	claims := tokenizer.Claims.(jwt.MapClaims)
	claims["name"] = userData.Username
	claims["fullName"] = userData.Name
	claims["email"] = userData.Email
	claims["id"] = userData.Id
	claims["role"] = userData.Role
	claims["ip"] = c.Context().RemoteIP().To16().String()
	claims["exp"] = expired.Unix()

	// Generate encoded token and send it as response.
	if token, err = tokenizer.SignedString(utils.ToBytes(os.Getenv("SECRET_KEY"))); err != nil {
		return nil, c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:     os.Getenv("COOKIE_TOKEN"),
		Value:    token,
		Expires:  *expired,
		HTTPOnly: true,
		SameSite: "strict",
	})

	if refreshToken != "" {
		c.Cookie(&fiber.Cookie{
			Name:     os.Getenv("COOKIE_REFRESH_TOKEN"),
			Value:    refreshToken,
			Expires:  (*expired).Add(30 * 24 * time.Hour),
			HTTPOnly: true,
			SameSite: "strict",
		})
	}

	return &token, err
}

func AuthorizationHandler(c *fiber.Ctx, db *gorm.DB, pageRoles []string) error {
	var (
		currentUserId int
		userdata      *viewmodels.AuthorizationModel
		claims        jwt.MapClaims
		tokenRole     string
		pageRole      string
		exp           int64
		authorized    = false
		ip            string
	)

	if user := c.Locals("user").(*jwt.Token); user != nil {
		claims = user.Claims.(jwt.MapClaims)

		tokenRole = claims["role"].(string)
		if len(pageRoles) == 0 {
			authorized = true
		} else {
			for _, pageRole = range pageRoles {
				authorized = authorized || tokenRole == pageRole
			}
		}

		if !authorized {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		currentUserId = int(claims["id"].(float64))
		exp = int64(claims["exp"].(float64))
		ip = claims["ip"].(string)

		if strings.Compare(ip, c.Context().RemoteIP().To16().String()) != 0 {
			c.ClearCookie(os.Getenv("COOKIE_TOKEN"))
			c.Status(fiber.StatusUnauthorized)

			return errors.New(
				fmt.Sprintf("User ID %s IP's changed from %s to %s",
					currentUserId, ip,
					c.IP()))
		}

		userdata = &viewmodels.AuthorizationModel{
			ID:       uint(currentUserId),
			Username: claims["name"].(string),
			Email:    claims["email"].(string),
			Role:     claims["role"].(string),
			FullName: claims["fullName"].(string),
			Exp:      exp,
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
