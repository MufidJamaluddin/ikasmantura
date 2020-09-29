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

type AlbumPhotoHandler struct {
	Service services.AlbumPhotoService
}

// SearchAlbumPhoto godoc
// @Tags Album
// @Summary Search photo data
// @Description Get photo data with pagination
// @Accept  json
// @Produce  json
// @Param q query dto.AlbumPhotoParam true "Pagination Options"
// @Success 200 {object} []dto.AlbumPhotoDto
// @Failure 400 {object} string
// @Router /api/v1/photos [get]
func (p *AlbumPhotoHandler) SearchAlbumPhoto(c *fiber.Ctx) error {
	var (
		data     dto.AlbumPhotoParam
		err      error
		total    uint
		callback func(photoDto *dto.AlbumPhotoDto)
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
	callback = func(dt *dto.AlbumPhotoDto) {
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

// GetOneAlbumPhoto godoc
// @Tags Album
// @Summary Get one album photo data by id
// @Description Get album photo data by id
// @Param id path int true "Album Photo ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.AlbumPhotoDto
// @Failure 400 {object} string
// @Router /api/v1/photos/{id} [get]
func (p *AlbumPhotoHandler) GetOneAlbumPhoto(c *fiber.Ctx) error {
	var (
		data dto.AlbumPhotoDto
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

// UpdateAlbumPhoto godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Album
// @Summary Update photo
// @Description Update photo
// @Accept  json
// @Produce  json
// @Param id path int true "AlbumPhoto ID"
// @Param q body dto.AlbumPhotoDto true "New AlbumPhoto Data"
// @Success 202 {object} dto.AlbumPhotoDto
// @Failure 400 {object} string
// @Router /api/v1/photos/{id} [put]
func (p *AlbumPhotoHandler) UpdateAlbumPhoto(c *fiber.Ctx) error {
	var (
		data dto.AlbumPhotoDto
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

// SaveAlbumPhoto godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Album
// @Summary Save photo
// @Description Save photo
// @Accept  json
// @Produce  json
// @Param q body dto.AlbumPhotoDto true "New AlbumPhoto Data"
// @Success 202 {object} dto.AlbumPhotoDto
// @Failure 400 {object} string
// @Router /api/v1/photos [put]
func (p *AlbumPhotoHandler) SaveAlbumPhoto(c *fiber.Ctx) error {
	var (
		data dto.AlbumPhotoDto
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

// DeleteAlbumPhoto godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Album
// @Summary Delete one album photo by id
// @Description Delete one album photo by id
// @Param id path int true "Album Photo ID"
// @Accept  json
// @Produce  json
// @Success 202 {object} dto.AlbumPhotoDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/photos/{id} [delete]
func (p *AlbumPhotoHandler) DeleteAlbumPhoto(c *fiber.Ctx) error {
	var (
		data dto.AlbumPhotoDto
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
