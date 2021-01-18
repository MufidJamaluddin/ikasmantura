package user

import (
	userService "backend/services/user"
	"backend/utils"
	"backend/viewmodels"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @author Mufid Jamaluddin

// SearchUser godoc
// @Tags User Management
// @Summary Search Department
// @Description Search Department
// @Accept  json
// @Produce  json
// @Param q query viewmodels.UserParam true "Pagination Options"
// @Success 200 {object} []viewmodels.UserDto
// @Failure 400 {object} string
// @Router /api/v1/users [get]
func SearchUser(c *fiber.Ctx) error {
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

	if total, err = userService.GetTotal(db, &data); err != nil {
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
	err = userService.Find(db, &data, callback)
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

// GetOneUser godoc
// @Tags User Management
// @Summary Get one data by id
// @Description Get data by id
// @Param id path int true "User ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} viewmodels.UserDto
// @Failure 400 {object} string
// @Router /api/v1/users/{id} [get]
func GetOneUser(c *fiber.Ctx) error {
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

	if err = userService.FindById(db, id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusOK)
	err = c.JSON(&data)

	return err
}

// UpdateUser godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags User
// @Summary Update user
// @Description Update user
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param q body viewmodels.UserDto true "New User Data"
// @Success 202 {object} viewmodels.UserDto
// @Failure 400 {object} string
// @Router /api/v1/users/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
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

	if err = userService.FindById(db, id, &data); err != nil {
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

	if err = userService.Update(db, id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)
	err = c.JSON(&data)

	return err
}

// SaveUser godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags User
// @Summary Add new user
// @Description Add new user
// @Accept  json
// @Produce  json
// @Param q body viewmodels.UserDto true "New User Data"
// @Success 202 {object} viewmodels.UserDto
// @Failure 400 {object} string
// @Router /api/v1/users [post]
func SaveUser(c *fiber.Ctx) error {
	var (
		data viewmodels.UserDto
		err  error
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

	if err = c.BodyParser(&data); err != nil {
		return err
	}

	data.CreatedBy = authData.ID
	data.UpdatedBy = authData.ID
	if err = userService.Save(db, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)
	err = c.JSON(&data)

	return err
}

// DeleteUser godoc
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
func DeleteUser(c *fiber.Ctx) error {
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
	if err = userService.Delete(db, id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)
	err = c.JSON(&data)

	return err
}
