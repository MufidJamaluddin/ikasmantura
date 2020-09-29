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

type AlbumHandler struct {
	Service services.AlbumService
}

// SearchAlbum godoc
// @Tags Album
// @Summary Search album data
// @Description Get album data with pagination
// @Accept  json
// @Produce  json
// @Param q query dto.AlbumParam true "Pagination Options"
// @Success 200 {object} []dto.AlbumDto
// @Failure 400 {object} string
// @Router /api/v1/albums [get]
func (p *AlbumHandler) SearchAlbum(c *fiber.Ctx) error {
	var (
		data     dto.AlbumParam
		err      error
		total    uint
		callback func(albumDto *dto.AlbumDto)
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
	callback = func(dt *dto.AlbumDto) {
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
// @Success 200 {object} dto.AlbumDto
// @Failure 400 {object} string
// @Router /api/v1/albums/{id} [get]
func (p *AlbumHandler) GetOneAlbum(c *fiber.Ctx) error {
	var (
		data dto.AlbumDto
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
// @Param q body dto.AlbumDto true "New Album Data"
// @Success 202 {object} dto.AlbumDto
// @Failure 400 {object} string
// @Router /api/v1/albums/{id} [put]
func (p *AlbumHandler) UpdateAlbum(c *fiber.Ctx) error {
	var (
		data dto.AlbumDto
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
// @Param q body dto.AlbumDto true "New Album Data"
// @Success 202 {object} dto.AlbumDto
// @Failure 400 {object} string
// @Router /api/v1/albums [post]
func (p *AlbumHandler) SaveAlbum(c *fiber.Ctx) error {
	var (
		data dto.AlbumDto
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
// @Success 202 {object} dto.AlbumDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/album/{id} [delete]
func (p *AlbumHandler) DeleteAlbum(c *fiber.Ctx) error {
	var (
		data dto.AlbumDto
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
