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
type AlbumController struct {
	Service services.AlbumService
}

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
func (p *AlbumController) SearchAlbum(c *fiber.Ctx) error {
	var (
		data     viewmodels.AlbumParam
		err      error
		total    uint
		callback func(albumDto *viewmodels.AlbumDto)
		counter  uint
	)

	if err = c.QueryParser(&data); err != nil {
		return err
	}

	if total, err = p.Service.GetTotal(&data); err != nil {
		return err
	}

	c.Response().Header.Add("X-Total-Count", fmt.Sprintf("%v", total))

	counter = data.Start
	callback = func(dt *viewmodels.AlbumDto) {
		var (
			response []byte
			e        error
		)
		if dt == nil {
			response = []byte("{}")
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
	if counter < total {
		_, _ = c.Write([]byte("{}"))
	}
	_, err = c.Write(utils.ToBytes("]"))

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
func (p *AlbumController) GetOneAlbum(c *fiber.Ctx) error {
	var (
		data viewmodels.AlbumDto
		id   uint
		err  error
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
func (p *AlbumController) UpdateAlbum(c *fiber.Ctx) error {
	var (
		data viewmodels.AlbumDto
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
func (p *AlbumController) SaveAlbum(c *fiber.Ctx) error {
	var (
		data viewmodels.AlbumDto
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

	data.UpdatedBy = currentUserId
	data.CreatedBy = currentUserId

	if err = p.Service.Save(&data); err != nil {
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
func (p *AlbumController) DeleteAlbum(c *fiber.Ctx) error {
	var (
		data viewmodels.AlbumDto
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
