package albumphoto

import (
	albumPhotoService "backend/services/albumphoto"
	"backend/utils"
	"backend/viewmodels"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

// @author Mufid Jamaluddin

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
func SearchAlbumPhoto(c *fiber.Ctx) error {
	var (
		data      viewmodels.AlbumPhotoParam
		err       error
		total     uint
		callback  func(photoDto *viewmodels.AlbumPhotoDto)
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

	if total, err = albumPhotoService.GetTotal(db, &data); err != nil {
		return err
	}

	c.Response().Header.Add("X-Total-Count", fmt.Sprintf("%v", total))

	// RESPONSE ARRAY JSON DATA
	// HEMAT MEMORY, NGGAK PERLU ALOKASI ARRAY, KIRIM AJA KE CLIENT SECARA MENGALIR
	isStarted = false
	counter = data.Start
	callback = func(dt *viewmodels.AlbumPhotoDto) {
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
	err = albumPhotoService.Find(db, &data, callback)
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
func GetOneAlbumPhoto(c *fiber.Ctx) error {
	var (
		data viewmodels.AlbumPhotoDto
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

	if err = albumPhotoService.FindById(db, id, &data); err != nil {
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
func UpdateAlbumPhoto(c *fiber.Ctx) error {
	var (
		data      viewmodels.AlbumPhotoDto
		imageFile *multipart.FileHeader
		image     string
		thumbnail string
		fileName  string
		err       error
		id        uint
		db        *gorm.DB
		ok        bool

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

	if err = albumPhotoService.FindById(db, id, &data); err != nil {
		_ = c.SendStatus(fiber.StatusNotFound)
		return err
	}

	imageFile, err = c.FormFile("image")

	if err == nil {
		if data.Image != "" {
			_ = os.Remove(fmt.Sprintf("/%s", data.Image))
		}
		if data.Thumbnail != "" {
			_ = os.Remove(fmt.Sprintf("/%s", data.Thumbnail))
		}

		fileName = strconv.Itoa(data.Id)
		if image, err = utils.UploadImageJPG(c, imageFile, fileName); err == nil {
			thumbnail, err = utils.UploadImageThumbJPG(imageFile, fileName)
		}
	} else {
		log.Println(err.Error())
	}

	if err = c.BodyParser(&data); err != nil {
		return err
	}

	data.Image = image
	data.Thumbnail = thumbnail
	data.UpdatedBy = authData.ID

	if err = albumPhotoService.Update(db, id, &data); err != nil {
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
// @Router /api/v1/photos [post]
func SaveAlbumPhoto(c *fiber.Ctx) error {
	var (
		data      viewmodels.AlbumPhotoDto
		imageFile *multipart.FileHeader
		image     string
		err       error
		db        *gorm.DB
		ok        bool

		authData *viewmodels.AuthorizationModel
		fileName string
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

	imageFile, err = c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("No images!")
	}

	if err = albumPhotoService.Save(db, &data); err != nil {
		return err
	}

	fileName = strconv.Itoa(data.Id)
	if image, err = utils.UploadImageJPG(c, imageFile, fileName); err == nil {
		data.Image = image
		if image, err = utils.UploadImageThumbJPG(imageFile, fileName); err == nil {
			data.Thumbnail = image
		}
	}

	if err != nil {
		log.Print(err.Error())
	}

	if data.Image != "" {
		if err = albumPhotoService.Update(db, uint(data.Id), &data); err != nil {
			return err
		}
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
func DeleteAlbumPhoto(c *fiber.Ctx) error {
	var (
		data  viewmodels.AlbumPhotoDto
		image string
		err   error
		id    uint
		db    *gorm.DB
		ok    bool
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if id, err = utils.ToUint(c.Params("id")); err != nil {
		return err
	}

	if err = albumPhotoService.Delete(db, id, &data); err != nil {
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
