package album

import (
	albumService "backend/services/album"
	"backend/utils"
	"backend/viewmodels"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @author Mufid Jamaluddin

// SearchAlbum godoc
// @Tags Album
// @Summary Search album data
// @Description Get album data with pagination
// @Accept  json
// @Produce  json
// @Param q query viewmodels.AlbumParam true "Pagination Options"
// @Success 200 {object} []viewmodels.AlbumDto
// @Failure 400 {object} string
// @Router /api/v1/albums [get]
func SearchAlbum(c *fiber.Ctx) error {
	var (
		data      viewmodels.AlbumParam
		err       error
		total     uint
		callback  func(albumDto *viewmodels.AlbumDto)
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

	if total, err = albumService.GetTotal(db, &data); err != nil {
		return err
	}

	c.Response().Header.Add("X-Total-Count", fmt.Sprintf("%v", total))

	// RESPONSE ARRAY JSON DATA
	// HEMAT MEMORY, NGGAK PERLU ALOKASI ARRAY, KIRIM AJA KE CLIENT SECARA MENGALIR
	isStarted = false
	callback = func(dt *viewmodels.AlbumDto) {
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
	err = albumService.Find(db, &data, callback)
	_, err = c.Write(utils.ToBytes("]"))
	// END RESPONSE ARRAY JSON DATA

	return err
}

// GetOneAlbum godoc
// @Tags Album
// @Summary Get one album data by id
// @Description Get album data by id
// @Param id path int true "Album ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} viewmodels.AlbumDto
// @Failure 400 {object} string
// @Router /api/v1/albums/{id} [get]
func GetOneAlbum(c *fiber.Ctx) error {
	var (
		data viewmodels.AlbumDto
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

	if err = albumService.FindById(db, id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusOK)

	err = c.JSON(&data)
	return err
}

// UpdateAlbum godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Album
// @Summary Update album
// @Description Update album
// @Accept  json
// @Produce  json
// @Param id path int true "Album ID"
// @Param q body viewmodels.AlbumDto true "New Album Data"
// @Success 202 {object} viewmodels.AlbumDto
// @Failure 400 {object} string
// @Router /api/v1/albums/{id} [put]
func UpdateAlbum(c *fiber.Ctx) error {
	var (
		data viewmodels.AlbumDto
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

	if err = albumService.Update(db, id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)

	err = c.JSON(&data)
	return err
}

// SaveAlbum godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Album
// @Summary Save album
// @Description Save album
// @Accept  json
// @Produce  json
// @Param q body viewmodels.AlbumDto true "New Album Data"
// @Success 202 {object} viewmodels.AlbumDto
// @Failure 400 {object} string
// @Router /api/v1/albums [post]
func SaveAlbum(c *fiber.Ctx) error {
	var (
		data viewmodels.AlbumDto
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

	if err = albumService.Save(db, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)
	err = c.JSON(&data)

	return err
}

// DeleteAlbum godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Album
// @Summary Delete one album by id
// @Description Delete one album by id
// @Param id path int true "Album ID"
// @Accept  json
// @Produce  json
// @Success 202 {object} viewmodels.AlbumDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/album/{id} [delete]
func DeleteAlbum(c *fiber.Ctx) error {
	var (
		data     viewmodels.AlbumDto
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
	if err = albumService.Delete(db, id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)

	err = c.JSON(&data)
	return err
}
