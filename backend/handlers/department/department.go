package department

import (
	departmentService "backend/services/department"
	"backend/utils"
	"backend/viewmodels"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @author Mufid Jamaluddin

// SearchDepartment godoc
// @Tags User Management
// @Summary Search Department
// @Description Search Department
// @Accept  json
// @Produce  json
// @Param q query viewmodels.DepartmentParam true "Pagination Options"
// @Success 200 {object} []viewmodels.DepartmentDto
// @Failure 400 {object} string
// @Router /api/v1/departments [get]
func SearchDepartment(c *fiber.Ctx) error {
	var (
		data      viewmodels.DepartmentParam
		err       error
		total     uint
		callback  func(departmentDto *viewmodels.DepartmentDto)
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

	if total, err = departmentService.GetTotal(db, &data); err != nil {
		return err
	}

	c.Response().Header.Add("X-Total-Count", fmt.Sprintf("%v", total))

	// RESPONSE ARRAY JSON DATA
	// HEMAT MEMORY, NGGAK PERLU ALOKASI ARRAY, KIRIM AJA KE CLIENT SECARA MENGALIR
	isStarted = false
	counter = data.Start
	callback = func(dt *viewmodels.DepartmentDto) {
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
	err = departmentService.Find(db, &data, callback)
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

// GetOneDepartment godoc
// @Tags User Management
// @Summary Get one data by id
// @Description Get data by id
// @Param id path int true "Department ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} viewmodels.DepartmentDto
// @Failure 400 {object} string
// @Router /api/v1/departments/{id} [get]
func GetOneDepartment(c *fiber.Ctx) error {
	var (
		data viewmodels.DepartmentDto
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

	if err = departmentService.FindById(db, id, &data); err != nil {
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
// @Param q body viewmodels.DepartmentDto true "New Department Data"
// @Success 202 {object} viewmodels.DepartmentDto
// @Failure 400 {object} string
// @Router /api/v1/departments/{id} [put]
func UpdateDepartment(c *fiber.Ctx) error {
	var (
		data viewmodels.DepartmentDto
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

	if err = c.BodyParser(&data); err != nil {
		return err
	}

	data.UpdatedBy = authData.ID
	if err = departmentService.Update(db, id, &data); err != nil {
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
// @Param q body viewmodels.DepartmentDto true "New Department Data"
// @Success 202 {object} viewmodels.DepartmentDto
// @Failure 400 {object} string
// @Router /api/v1/departments [post]
func SaveDepartment(c *fiber.Ctx) error {
	var (
		data viewmodels.DepartmentDto
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
	if err = departmentService.Save(db, &data); err != nil {
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
// @Success 202 {object} viewmodels.DepartmentDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/department/{id} [delete]
func DeleteDepartment(c *fiber.Ctx) error {
	var (
		data viewmodels.DepartmentDto
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
	if err = departmentService.Delete(db, id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)
	err = c.JSON(&data)

	return err
}
