package auth

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
	"backend/viewmodels"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
	ua2 "github.com/mileusna/useragent"
	uuid "github.com/satori/go.uuid"
	history "github.com/vcraescu/gorm-history/v2"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetUserAgentData(c *fiber.Ctx) *ua2.UserAgent {
	var (
		userAgentStr string
		userAgent ua2.UserAgent
	)
	userAgentStr = string(c.Context().UserAgent())
	userAgent = ua2.Parse(userAgentStr)
	return &userAgent
}

func SaveUserLogin(c *fiber.Ctx, userId uint) (*models.UserLogin, error) {

	var (
		err 	  		error
		ok 		  		bool
		db		  		*gorm.DB
		userAgent 		*ua2.UserAgent
		userLoginData	models.UserLogin
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return nil, c.SendStatus(fiber.StatusInternalServerError)
	}

	userAgent = GetUserAgentData(c)

	userLoginData.UserId = userId
	userLoginData.RemoteIP = c.Context().RemoteIP().String()
	userLoginData.Device = strings.Trim(userAgent.Device, " ")[0:35]
	userLoginData.OSName = strings.Trim(userAgent.OS, " ")[0:35]
	userLoginData.OSVersion = strings.Trim(userAgent.OSVersion, " ")[0:10]

	if userAgent.Bot {
		userLoginData.DeviceType = "bot"
	} else if userAgent.Mobile {
		userLoginData.DeviceType = "mobile"
	} else if userAgent.Tablet {
		userLoginData.DeviceType = "tablet"
	} else if userAgent.Desktop {
		userLoginData.DeviceType = "desktop"
	}

	err = repository.Save(db, &userLoginData)

	return &userLoginData, err
}

func DoLogin(
	c *fiber.Ctx,
	userLogin *models.UserLogin,
	userData *viewmodels.UserDto,
	expired *time.Time,
	) (*string, error) {

	var (
		err 	  		error
		token     		string
		refreshToken    string
		tokenizer 		*jwt.Token
	)

	// Create token
	tokenizer = jwt.New(jwt.SigningMethodHS256)

	// Expiration
	*expired = time.Now().Add(time.Hour * 3)

	// Set claims
	claims := tokenizer.Claims.(jwt.MapClaims)
	claims["name"] = userData.Username
	claims["fullName"] = userData.Name
	claims["email"] = userData.Email
	claims["id"] = userData.Id
	claims["role"] = userData.Role
	claims["ip"] = c.Context().RemoteIP().String()
	claims["exp"] = expired.Unix()

	// Generate encoded token and send it as response.
	if token, err = tokenizer.SignedString(utils.ToBytes(os.Getenv("SECRET_KEY")));
	err != nil {
		return nil, c.SendStatus(fiber.StatusInternalServerError)
	}

	if userLogin.RefreshToken != uuid.Nil {
		refreshToken = userLogin.RefreshToken.String()
	} else {
		refreshToken = ""
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
			Expires:  (*expired).Add(5 * 24 * time.Hour),
			HTTPOnly: true,
			SameSite: "strict",
		})
	}

	return &token, err
}

func AuthorizationHandler(c *fiber.Ctx, db *gorm.DB, pageRoles []string) error {
	var (
		currentUserId int
		seq           uint64
		userdata      *viewmodels.AuthorizationModel
		claims        jwt.MapClaims
		remoteIp      string
		tokenRole     string
		pageRole      string
		exp           int64
		authorized    = false
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
		seq = uint64(claims["pi"].(float64))
		exp = int64(claims["exp"].(float64))
		remoteIp = c.Context().RemoteIP().String()

		if claims["ip"].(string) != c.Context().RemoteIP().String() {
			c.ClearCookie(os.Getenv("COOKIE_TOKEN"))
			c.Status(fiber.StatusUnauthorized)
			return errors.New(
				fmt.Sprintf("User ID %s IP's changed from %s to %s",
					currentUserId, claims["ip"], remoteIp))
		}

		userdata = &viewmodels.AuthorizationModel{
			ID:       uint(currentUserId),
			Username: claims["name"].(string),
			Email:    claims["email"].(string),
			Role:  	  claims["role"].(string),
			FullName: claims["fullName"].(string),
			Exp:      exp,
			Seq:      seq,
		}

		c.Locals("db", history.SetUser(db, history.User{
			ID:    strconv.Itoa(currentUserId),
			Email: claims["email"].(string),
		}))

		c.Locals("user", userdata)
	} else {
		c.ClearCookie(os.Getenv("COOKIE_TOKEN"))
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.Next()
}
