package handlers

import (
	"backend/dto"
	"backend/services"
	"backend/utils"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Service services.UserService
}

// SearchUser godoc
// @Tags User Management
// @Summary Search Department
// @Description Search Department
// @Accept  json
// @Produce  json
// @Param q query dto.UserParam true "Pagination Options"
// @Success 200 {object} []dto.UserDto
// @Failure 400 {object} string
// @Router /api/v1/users [get]
func (p *UserHandler) SearchUser(c *fiber.Ctx) error {
	var (
		data     dto.UserParam
		err      error
		total    uint
		callback func(userDto *dto.UserDto)
		counter  uint
	)

	if err = c.QueryParser(&data); err != nil {
		return err
	}

	if total, err = p.Service.GetTotal(&data); err != nil {
		return err
	}

	c.Response().Header.Add("X-Total-Count", fmt.Sprintf("%v", total))

	counter = 0
	callback = func(dt *dto.UserDto) {
		var (
			response []byte
			e error
		)
		if dt == nil {
			response = []byte("{}")
		} else if response, e = json.Marshal(dt); e == nil {
			_, _ = c.Write(response)
		}
		counter++
		if counter < total {
			_, _ = c.Write([]byte(","))
		}
	}

	_, err = c.Write(utils.ToBytes("["))
	err = p.Service.Find(&data, callback)
	if counter < total {
		_, _ = c.Write([]byte("{}"))
	}
	_, err = c.Write(utils.ToBytes("]"))

	return err
}

// GetOneUser godoc
// @Tags User Management
// @Summary Get one data by id
// @Description Get data by id
// @Param id path int true "User ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.UserDto
// @Failure 400 {object} string
// @Router /api/v1/users/{id} [get]
func (p *UserHandler) GetOneUser(c *fiber.Ctx) error {
	var (
		data dto.UserDto
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
// @Param q body dto.UserDto true "New User Data"
// @Success 202 {object} dto.UserDto
// @Failure 400 {object} string
// @Router /api/v1/users/{id} [put]
func (p *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var (
		data dto.UserDto
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
// @Param q body dto.UserDto true "New User Data"
// @Success 202 {object} dto.UserDto
// @Failure 400 {object} string
// @Router /api/v1/users [post]
func (p *UserHandler) SaveUser(c *fiber.Ctx) error {
	var (
		data dto.UserDto
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
// @Success 202 {object} dto.UserDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/user/{id} [delete]
func (p *UserHandler) DeleteUser(c *fiber.Ctx) error {
	var (
		data dto.UserDto
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
