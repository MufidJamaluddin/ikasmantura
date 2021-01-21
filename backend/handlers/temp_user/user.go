package temp_user

import (
	"backend/services/email"
	tempUserService "backend/services/temp_user"
	"backend/utils"
	"backend/viewmodels"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/url"
	"path"
)

// @author Mufid Jamaluddin

// SearchTempUser godoc
// @Tags User Management
// @Summary Search Department
// @Description Search Department
// @Accept  json
// @Produce  json
// @Param q query viewmodels.UserParam true "Pagination Options"
// @Success 200 {object} []viewmodels.UserDto
// @Failure 400 {object} string
// @Router /api/v1/temp_users [get]
func SearchTempUser(c *fiber.Ctx) error {
	var (
		data      viewmodels.UserParam
		err       error
		total     uint
		callback  func(userDto *viewmodels.UserDto)
		isStarted bool
		db        *gorm.DB
		ok        bool
		counter   uint
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err = c.QueryParser(&data); err != nil {
		return err
	}

	if total, err = tempUserService.GetTotal(db, &data); err != nil {
		return err
	}

	c.Response().Header.Add("X-Total-Count", fmt.Sprintf("%v", total))

	// RESPONSE ARRAY JSON DATA
	// HEMAT MEMORY, NGGAK PERLU ALOKASI ARRAY, KIRIM AJA KE CLIENT SECARA MENGALIR
	isStarted = false
	counter = data.Start
	callback = func(dt *viewmodels.UserDto) {
		var (
			response []byte
			e        error
		)
		if isStarted {
			_, _ = c.Write([]byte(","))
		}
		if dt == nil {
			response = []byte("{}")
			_, _ = c.Write(response)
		} else if response, e = json.Marshal(dt); e == nil {
			_, _ = c.Write(response)
		}
		isStarted = true
		counter++
	}

	_, err = c.Write(utils.ToBytes("["))
	err = tempUserService.Find(db, &data, callback)
	_, err = c.Write(utils.ToBytes("]"))
	// END RESPONSE ARRAY JSON DATA

	if data.Start < counter {
		c.Response().Header.Add("Content-Range",
			fmt.Sprintf("items %v-%v/%v", data.Start, counter, total))

		if total == counter {
			c.Response().Header.SetStatusCode(fiber.StatusOK)
		} else {
			c.Response().Header.SetStatusCode(fiber.StatusPartialContent)
		}
	}

	return err
}

// GetOneTempUser godoc
// @Tags User Management
// @Summary Get one data by id
// @Description Get data by id
// @Param id path int true "User ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} viewmodels.UserDto
// @Failure 400 {object} string
// @Router /api/v1/temp_users/{id} [get]
func GetOneTempUser(c *fiber.Ctx) error {
	var (
		data viewmodels.UserDto
		err  error
		id   uint
		db   *gorm.DB
		ok   bool
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if id, err = utils.ToUint(c.Params("id")); err != nil {
		return err
	}

	if err = tempUserService.FindById(db, id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusOK)
	err = c.JSON(&data)

	return err
}

// UpdateTempUser godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags User Management
// @Summary Update user
// @Description Update user
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param q body viewmodels.UserDto true "New User Data"
// @Success 202 {object} viewmodels.UserDto
// @Failure 400 {object} string
// @Router /api/v1/temp_users/{id} [put]
func UpdateTempUser(c *fiber.Ctx) error {
	var (
		data viewmodels.UserDto
		id   uint
		err  error
		db   *gorm.DB
		ok   bool

		tempUsername string
		tempPassword string

		authData *viewmodels.AuthorizationModel
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if authData, ok = viewmodels.GetAuthorizationData(c); !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if id, err = utils.ToUint(c.Params("id")); err != nil {
		return err
	}

	if err = tempUserService.FindById(db, id, &data); err != nil {
		_ = c.SendStatus(fiber.StatusNotFound)
		return err
	}

	tempUsername = data.Username
	tempPassword = data.Password

	if err = c.BodyParser(&data); err != nil {
		return err
	}

	data.Username = tempUsername
	data.Password = tempPassword
	data.UpdatedBy = authData.ID

	if err = tempUserService.Update(db, id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)
	err = c.JSON(&data)

	return err
}

// CheckAvailabilityUser godoc
// @Tags User Management
// @Summary Add new user
// @Description Add new user
// @Accept  json
// @Produce  json
// @Param q body viewmodels.UserAvailabilityDto true "New User Data"
// @Success 202 {object} viewmodels.UserAvailabilityResponseDto
// @Failure 400 {object} string
// @Router /api/v1/register/availability [post]
func CheckAvailabilityUser(c *fiber.Ctx) error {
	var (
		availabilityReq viewmodels.UserAvailabilityDto
		availabilityRes viewmodels.UserAvailabilityResponseDto
		err             error
		db              *gorm.DB
		ok              bool
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err = c.BodyParser(&availabilityReq); err != nil {
		return err
	}

	if err = tempUserService.IsUsernameOrEmailAvailable(db, &availabilityReq, &availabilityRes); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)
	err = c.JSON(&availabilityRes)

	return err
}

// VerifyUser godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags User Management
// @Summary Add new user
// @Description Add new user
// @Accept  json
// @Produce  json
// @Param q body viewmodels.UserDto true "New User Data"
// @Success 202 {object} viewmodels.UserDto
// @Failure 400 {object} string
// @Router /api/v1/verify_user [post]
func VerifyUser(c *fiber.Ctx) error {
	var (
		data viewmodels.UserDto
		err  error
		db   *gorm.DB
		ok   bool
		id   uint
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if id, err = utils.ToUint(c.Params("id")); err != nil {
		return err
	}

	if err = c.BodyParser(&data); err != nil {
		return err
	}

	if err = tempUserService.Verify(db, id, &data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	c.Status(fiber.StatusAccepted)
	err = c.JSON(&data)

	return err
}

// ConfirmTempUserEmail godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags User Management
// @Summary Add new user
// @Description Add new user
// @Accept  json
// @Produce  json
// @Success 202
// @Failure 400 {object} string
// @Router /api/v1/confirms/tu_emails/{username}/{token} [post]
func ConfirmTempUserEmail(c *fiber.Ctx) (err error) {
	var (
		db                   *gorm.DB
		ok                   bool
		username             string
		confirmEmailToken    string
		confirmEmailTokenUid utils.UUID
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	username = c.Path("username")
	confirmEmailToken = c.Path("token")

	if confirmEmailTokenUid, err = utils.FromBase64UUID(confirmEmailToken); err != nil {
		log.Println(err.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if username == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err = tempUserService.ConfirmEmail(db, username, confirmEmailTokenUid); err != nil {
		log.Println(err.Error())
		return c.SendStatus(fiber.StatusForbidden)
	}

	return c.SendStatus(fiber.StatusAccepted)
}

// SaveTempUser godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags User Management
// @Summary Add new user
// @Description Add new user
// @Accept  json
// @Produce  json
// @Param q body viewmodels.UserDto true "New User Data"
// @Success 202 {object} viewmodels.UserDto
// @Failure 400 {object} string
// @Router /api/v1/temp_users [post]
func SaveTempUser(c *fiber.Ctx) error {
	var (
		data            viewmodels.UserDto
		availabilityReq viewmodels.UserAvailabilityDto
		availabilityRes viewmodels.UserAvailabilityResponseDto
		err             error
		db              *gorm.DB
		ok              bool
		confirmUrl      *url.URL
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err = c.BodyParser(&data); err != nil {
		return err
	}

	availabilityReq.Username = data.Username
	availabilityReq.Email = data.Email

	if err = tempUserService.IsUsernameOrEmailAvailable(db, &availabilityReq, &availabilityRes); err != nil {
		return err
	}

	if *availabilityRes.Exist {
		return c.Status(fiber.StatusConflict).
			SendString("Username atau Email telah terdaftar!")
	}

	if err = tempUserService.Save(db, &data); err != nil {
		return err
	}

	confirmUrl = utils.GetBasePath()
	confirmUrl.Path = path.Join(confirmUrl.Path,
		fmt.Sprintf("register_confirm/%v/%v",
			data.Username, data.ConfirmEmailToken))

	emailMsg := &viewmodels.EmailMessage{}
	emailMsg.Header = "Registrasi Data Alumni"
	emailMsg.Title = "Registrasi Anggota Ikatan Alumni SMAN Situraja"
	emailMsg.To = []string{data.Email}
	emailMsg.Message = template.HTML(fmt.Sprintf(
		"Registrasi %v (Username %v - Email %v) Sukses! "+
			"<br/><br/>Mohon Tunggu Kabar dari Kepengurusan IKA SMAN Situraja! "+
			"<br/><br/><i>Tekan tombol Verifikasi Email dibawah ini untuk verifikasi pendaftaran anda</i>"+
			"<br/><br/><a href=\"%v\">"+
			"<button style=\"background-color:#212529;color:#ffff;\">Verifikasi Email</button></a>",
		data.Name,
		data.Username,
		data.Email,
		confirmUrl.String()))

	email.SendMessage(emailMsg)

	c.Status(fiber.StatusAccepted)
	err = c.JSON(&data)

	return err
}

// DeleteTempUser godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags User Management
// @Summary Delete one user by id
// @Description Delete one user by id
// @Param id path int true "User ID"
// @Accept  json
// @Produce  json
// @Success 202 {object} viewmodels.UserDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/user/{id} [delete]
func DeleteTempUser(c *fiber.Ctx) error {
	var (
		data viewmodels.UserDto
		err  error
		id   uint
		db   *gorm.DB
		ok   bool

		authData *viewmodels.AuthorizationModel
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if authData, ok = viewmodels.GetAuthorizationData(c); !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if id, err = utils.ToUint(c.Params("id")); err != nil {
		return err
	}

	data.UpdatedBy = authData.ID
	if err = tempUserService.Delete(db, id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)
	err = c.JSON(&data)

	return err
}
