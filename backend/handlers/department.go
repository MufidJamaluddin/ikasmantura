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

type DepartmentHandler struct {
	Service services.DepartmentService
}

// SearchDepartment godoc
// @Tags User Management
// @Summary Search Department
// @Description Search Department
// @Accept  json
// @Produce  json
// @Param q query dto.DepartmentParam true "Pagination Options"
// @Success 200 {object} []dto.DepartmentDto
// @Failure 400 {object} string
// @Router /api/v1/departments [get]
func (p *DepartmentHandler) SearchDepartment(c *fiber.Ctx) error {
	var (
		data     dto.DepartmentParam
		err      error
		total    uint
		callback func(departmentDto *dto.DepartmentDto)
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
	callback = func(dt *dto.DepartmentDto) {
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

// GetOneDepartment godoc
// @Tags User Management
// @Summary Get one data by id
// @Description Get data by id
// @Param id path int true "Department ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.DepartmentDto
// @Failure 400 {object} string
// @Router /api/v1/departments/{id} [get]
func (p *DepartmentHandler) GetOneDepartment(c *fiber.Ctx) error {
	var (
		data dto.DepartmentDto
		err  error
		id   uint
	)

	if id, err = utils.ToUint(c.Params("id")); err != nil {
		return err
	}

	if err = p.Service.FindById(id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)
	err = c.JSON(&data)

	return err
}

// UpdateDepartment godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags User Management
// @Summary Update department
// @Description Update department
// @Accept  json
// @Produce  json
// @Param id path int true "Department ID"
// @Param q body dto.DepartmentDto true "New Department Data"
// @Success 202 {object} dto.DepartmentDto
// @Failure 400 {object} string
// @Router /api/v1/departments/{id} [put]
func (p *DepartmentHandler) UpdateDepartment(c *fiber.Ctx) error {
	var (
		data dto.DepartmentDto
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

// SaveDepartment godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags User Management
// @Summary Save department
// @Description Update department
// @Accept  json
// @Produce  json
// @Param q body dto.DepartmentDto true "New Department Data"
// @Success 202 {object} dto.DepartmentDto
// @Failure 400 {object} string
// @Router /api/v1/departments [post]
func (p *DepartmentHandler) SaveDepartment(c *fiber.Ctx) error {
	var (
		data dto.DepartmentDto
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

// DeleteDepartment godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags User Management
// @Summary Delete one department  by id
// @Description Delete one department  by id
// @Param id path int true "Department ID"
// @Accept  json
// @Produce  json
// @Success 202 {object} dto.DepartmentDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/department/{id} [delete]
func (p *DepartmentHandler) DeleteDepartment(c *fiber.Ctx) error {
	var (
		data dto.DepartmentDto
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
