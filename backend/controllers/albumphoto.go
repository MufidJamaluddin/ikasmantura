package controllers

import (
	"backend/services"
	"backend/utils"
	"backend/viewmodels"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"os"
	"strings"
)

// @author Mufid Jamaluddin
type AlbumPhotoController struct {
	Service services.AlbumPhotoService
}

// SearchAlbumPhoto godoc
// @Tags Album
// @Summary Search photo data
// @Description Get photo data with pagination
// @Accept  json
// @Produce  json
// @Param q query viewmodels.AlbumPhotoParam true "Pagination Options"
// @Success 200 {object} []viewmodels.AlbumPhotoDto
// @Failure 400 {object} string
// @Router /api/v1/photos [get]
func (p *AlbumPhotoController) SearchAlbumPhoto(c *fiber.Ctx) error {
	var (
		data     viewmodels.AlbumPhotoParam
		err      error
		total    uint
		callback func(photoDto *viewmodels.AlbumPhotoDto)
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
	counter = 0
	callback = func(dt *viewmodels.AlbumPhotoDto) {
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
	// END RESPONSE ARRAY JSON DATA

	return err
}

// GetOneAlbumPhoto godoc
// @Tags Album
// @Summary Get one album photo data by id
// @Description Get album photo data by id
// @Param id path int true "Album Photo ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} viewmodels.AlbumPhotoDto
// @Failure 400 {object} string
// @Router /api/v1/photos/{id} [get]
func (p *AlbumPhotoController) GetOneAlbumPhoto(c *fiber.Ctx) error {
	var (
		data viewmodels.AlbumPhotoDto
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
// @Param q body viewmodels.AlbumPhotoDto true "New AlbumPhoto Data"
// @Success 202 {object} viewmodels.AlbumPhotoDto
// @Failure 400 {object} string
// @Router /api/v1/photos/{id} [put]
func (p *AlbumPhotoController) UpdateAlbumPhoto(c *fiber.Ctx) error {
	var (
		data      viewmodels.AlbumPhotoDto
		imageFile *multipart.FileHeader
		image     string
		err       error
		id        uint

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

	data.CreatedBy = currentUserId
	data.UpdatedBy = currentUserId

	imageFile, err = c.FormFile("image")
	if err == nil {
		if image, err = utils.UploadImageJPG(c, imageFile); err == nil {
			data.Image = image
			if image, err = utils.UploadImageThumbJPG(imageFile); err == nil {
				data.Thumbnail = image
			}
		}
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
// @Param q body viewmodels.AlbumPhotoDto true "New AlbumPhoto Data"
// @Success 202 {object} viewmodels.AlbumPhotoDto
// @Failure 400 {object} string
// @Router /api/v1/photos [put]
func (p *AlbumPhotoController) SaveAlbumPhoto(c *fiber.Ctx) error {
	var (
		data      viewmodels.AlbumPhotoDto
		imageFile *multipart.FileHeader
		image     string
		err       error

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

	imageFile, err = c.FormFile("image")
	if err == nil {
		if image, err = utils.UploadImageJPG(c, imageFile); err == nil {
			data.Image = image
			if image, err = utils.UploadImageThumbJPG(imageFile); err == nil {
				data.Thumbnail = image
			}
		}
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
// @Success 202 {object} viewmodels.AlbumPhotoDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/photos/{id} [delete]
func (p *AlbumPhotoController) DeleteAlbumPhoto(c *fiber.Ctx) error {
	var (
		data  viewmodels.AlbumPhotoDto
		image string
		err   error
		id    uint

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

	if image = strings.Trim(data.Image, " "); image != "" {
		_ = os.Remove(fmt.Sprintf("/%s", image))
	}

	if image = strings.Trim(data.Thumbnail, " "); image != "" {
		_ = os.Remove(fmt.Sprintf("/%s", image))
	}

	c.Status(fiber.StatusAccepted)
	err = c.JSON(&data)

	return err
}
