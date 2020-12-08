package controllers

import (
	"backend/services"
	"backend/utils"
	"backend/viewmodels"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// @author Mufid Jamaluddin
type UserController struct {
	Service services.UserService
}

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
func (p *UserController) SearchUser(c *fiber.Ctx) error {
	var (
		data     viewmodels.UserParam
		err      error
		total    uint
		callback func(userDto *viewmodels.UserDto)
		counter  uint
	)

	if err = c.QueryParser(&data); err != nil {
		return err
	}

	if total, err = p.Service.GetTotal(&data); err != nil {
		return err
	}

	c.Response().Header.Add("X-Total-Count", fmt.Sprintf("%v", total))

	// RESPONSE ARRAY JSON DATA
	// HEMAT MEMORY, NGGAK PERLU ALOKASI ARRAY, KIRIM AJA KE CLIENT SECARA MENGALIR
	counter = data.Start
	callback = func(dt *viewmodels.UserDto) {
		var (
			response []byte
			e        error
		)
		if dt == nil {
			response = []byte("{}")
			_, _ = c.Write(response)
		} else if response, e = json.Marshal(dt); e == nil {
			_, _ = c.Write(response)
		}
		counter++
		if counter < data.End {
			_, _ = c.Write([]byte(","))
		}
	}

	_, err = c.Write(utils.ToBytes("["))
	err = p.Service.Find(&data, callback)
	if counter < data.End {
		_, _ = c.Write([]byte("{}"))
	}
	_, err = c.Write(utils.ToBytes("]"))
	// END RESPONSE ARRAY JSON DATA

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
func (p *UserController) GetOneUser(c *fiber.Ctx) error {
	var (
		data viewmodels.UserDto
		err  error
		id   uint
	)

	if id, err = utils.ToUint(c.Params("id")); err != nil {
		return err
	}

	if err = p.Service.FindById(id, &data); err != nil {
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
func (p *UserController) UpdateUser(c *fiber.Ctx) error {
	var (
		data viewmodels.UserDto
		id   uint
		err  error

		currentUserId uint
	)

	if user := c.Locals("user").(*jwt.Token); user != nil {
		claims := user.Claims.(jwt.MapClaims)
		currentUserId = uint(claims["id"].(float64))
	} else {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if id, err = utils.ToUint(c.Params("id")); err != nil {
		return err
	}

	if err = c.BodyParser(&data); err != nil {
		return err
	}

	data.UpdatedBy = currentUserId
	if err = p.Service.Update(id, &data); err != nil {
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
func (p *UserController) SaveUser(c *fiber.Ctx) error {
	var (
		data viewmodels.UserDto
		err  error

		currentUserId uint
	)

	if user := c.Locals("user").(*jwt.Token); user != nil {
		claims := user.Claims.(jwt.MapClaims)
		currentUserId = uint(claims["id"].(float64))
	} else {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if err = c.BodyParser(&data); err != nil {
		return err
	}

	data.CreatedBy = currentUserId
	data.UpdatedBy = currentUserId
	if err = p.Service.Save(&data); err != nil {
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
func (p *UserController) DeleteUser(c *fiber.Ctx) error {
	var (
		data viewmodels.UserDto
		err  error
		id   uint

		currentUserId uint
	)

	if user := c.Locals("user").(*jwt.Token); user != nil {
		claims := user.Claims.(jwt.MapClaims)
		currentUserId = uint(claims["id"].(float64))
	} else {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if id, err = utils.ToUint(c.Params("id")); err != nil {
		return err
	}

	data.UpdatedBy = currentUserId
	if err = p.Service.Delete(id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)
	err = c.JSON(&data)

	return err
}
