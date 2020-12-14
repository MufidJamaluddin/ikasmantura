package classroom

import (
	classroomService "backend/services/classroom"
	"backend/utils"
	"backend/viewmodels"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @author Mufid Jamaluddin

// SearchClassroom godoc
// @Tags Classroom
// @Summary Search classroom data
// @Description Get classroom data with pagination
// @Accept  json
// @Produce  json
// @Param q query viewmodels.ClassroomParam true "Pagination Options"
// @Success 200 {object} []viewmodels.ClassroomDto
// @Failure 400 {object} string
// @Router /api/v1/classrooms [get]
func SearchClassroom(c *fiber.Ctx) error {
	var (
		data      viewmodels.ClassroomParam
		err       error
		total     uint
		callback  func(classroomDto *viewmodels.ClassroomDto)
		isStarted bool
		db        *gorm.DB
		ok        bool
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err = c.QueryParser(&data); err != nil {
		return err
	}

	if total, err = classroomService.GetTotal(db, &data); err != nil {
		return err
	}

	c.Response().Header.Add("X-Total-Count", fmt.Sprintf("%v", total))

	// RESPONSE ARRAY JSON DATA
	// HEMAT MEMORY, NGGAK PERLU ALOKASI ARRAY, KIRIM AJA KE CLIENT SECARA MENGALIR
	isStarted = false
	callback = func(dt *viewmodels.ClassroomDto) {
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
	}

	_, err = c.Write(utils.ToBytes("["))
	err = classroomService.Find(db, &data, callback)
	_, err = c.Write(utils.ToBytes("]"))
	// END RESPONSE ARRAY JSON DATA

	return err
}

// GetOneClassroom godoc
// @Tags Classroom
// @Summary Get one classroom data by id
// @Description Get classroom data by id
// @Param id path int true "Classroom ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} viewmodels.ClassroomDto
// @Failure 400 {object} string
// @Router /api/v1/classrooms/{id} [get]
func GetOneClassroom(c *fiber.Ctx) error {
	var (
		data viewmodels.ClassroomDto
		id   uint
		err  error
		db   *gorm.DB
		ok   bool
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if id, err = utils.ToUint(c.Params("id")); err != nil {
		return err
	}

	if err = classroomService.FindById(db, id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusOK)

	err = c.JSON(&data)
	return err
}

// UpdateClassroom godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Classroom
// @Summary Update classroom
// @Description Update classroom
// @Accept  json
// @Produce  json
// @Param id path int true "Classroom ID"
// @Param q body viewmodels.ClassroomDto true "New Classroom Data"
// @Success 202 {object} viewmodels.ClassroomDto
// @Failure 400 {object} string
// @Router /api/v1/classrooms/{id} [put]
func UpdateClassroom(c *fiber.Ctx) error {
	var (
		data viewmodels.ClassroomDto
		id   uint
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

	if id, err = utils.ToUint(c.Params("id")); err != nil {
		return err
	}

	if err = c.BodyParser(&data); err != nil {
		return err
	}

	data.UpdatedBy = authData.ID

	if err = classroomService.Update(db, id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)

	err = c.JSON(&data)
	return err
}

// SaveClassroom godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Classroom
// @Summary Save classroom
// @Description Save classroom
// @Accept  json
// @Produce  json
// @Param q body viewmodels.ClassroomDto true "New Classroom Data"
// @Success 202 {object} viewmodels.ClassroomDto
// @Failure 400 {object} string
// @Router /api/v1/classrooms [post]
func SaveClassroom(c *fiber.Ctx) error {
	var (
		data viewmodels.ClassroomDto
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

	data.UpdatedBy = authData.ID
	data.CreatedBy = authData.ID

	if err = classroomService.Save(db, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)
	err = c.JSON(&data)

	return err
}

// DeleteClassroom godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Classroom
// @Summary Delete one classroom by id
// @Description Delete one classroom by id
// @Param id path int true "Classroom ID"
// @Accept  json
// @Produce  json
// @Success 202 {object} viewmodels.ClassroomDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/classroom/{id} [delete]
func DeleteClassroom(c *fiber.Ctx) error {
	var (
		data     viewmodels.ClassroomDto
		authData *viewmodels.AuthorizationModel
		err      error
		id       uint
		db       *gorm.DB
		ok       bool
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
	if err = classroomService.Delete(db, id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)

	err = c.JSON(&data)
	return err
}
