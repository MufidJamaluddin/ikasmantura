package article

import (
	articleService "backend/services/article"
	"backend/utils"
	"backend/viewmodels"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"mime/multipart"
	"os"
	"strings"
)

// @author Mufid Jamaluddin

// SearchArticle godoc
// @Tags Article
// @Summary Search article data
// @Description Get article data with pagination
// @Accept  json
// @Produce  json
// @Param q query viewmodels.ArticleParam true "Pagination Options"
// @Success 200 {object} []viewmodels.ArticleDto
// @Failure 400 {object} string
// @Router /api/v1/articles [get]
func SearchArticle(c *fiber.Ctx) error {
	var (
		data      viewmodels.ArticleParam
		err       error
		total     uint
		callback  func(articleDto *viewmodels.ArticleDto)
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

	if total, err = articleService.GetTotal(db, &data); err != nil {
		return err
	}

	c.Response().Header.Add("X-Total-Count", fmt.Sprintf("%v", total))

	// RESPONSE ARRAY JSON DATA
	// HEMAT MEMORY, NGGAK PERLU ALOKASI ARRAY, KIRIM AJA KE CLIENT SECARA MENGALIR
	isStarted = false
	counter = data.Start
	callback = func(dt *viewmodels.ArticleDto) {
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
	err = articleService.Find(db, &data, callback)
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

// GetOneArticle godoc
// @Tags Article
// @Summary Get one data by id
// @Description Get data by id
// @Param id path int true "Article ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} viewmodels.ArticleDto
// @Failure 400 {object} string
// @Router /api/v1/articles/{id} [get]
func GetOneArticle(c *fiber.Ctx) error {
	var (
		data viewmodels.ArticleDto
		err  error
		id   string
		db   *gorm.DB
		ok   bool
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if id = c.Params("id"); id == "" {
		return errors.New("field ID wajib diisi")
	}

	if err = articleService.FindById(db, id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)

	err = c.JSON(&data)
	return err
}

// UpdateArticle godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Article
// @Summary Update article
// @Description Update article
// @Accept  json
// @Produce  json
// @Param id path int true "Article ID"
// @Param q body viewmodels.ArticleDto true "New Article Data"
// @Success 202 {object} viewmodels.ArticleDto
// @Failure 400 {object} string
// @Router /api/v1/articles/{id} [put]
func UpdateArticle(c *fiber.Ctx) error {
	var (
		data      viewmodels.ArticleDto
		image     string
		thumbnail string
		imageFile *multipart.FileHeader
		fileName  string
		err       error
		id        string
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

	if id = c.Params("id"); id == "" {
		return err
	}

	if err = articleService.FindById(db, id, &data); err != nil {
		_ = c.SendStatus(fiber.StatusNotFound)
		return err
	}

	if data.CreatedBy != authData.ID && authData.Role != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	fileName = id
	imageFile, err = c.FormFile("image")
	if err == nil {
		if data.Image != "" {
			_ = os.Remove(fmt.Sprintf("/%s", data.Image))
		}
		if data.Thumbnail != "" {
			_ = os.Remove(fmt.Sprintf("/%s", data.Thumbnail))
		}

		if image, err = utils.UploadImageJPG(c, imageFile, fileName); err == nil {
			thumbnail, err = utils.UploadImageThumbJPG(imageFile, fileName)
		}
	}

	if err != nil {
		log.Print(err.Error())
	}

	if err = c.BodyParser(&data); err != nil {
		return err
	}

	data.Image = image
	data.Thumbnail = thumbnail
	data.UserId = int(authData.ID)

	if err = articleService.Update(db, id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)

	err = c.JSON(&data)
	return err
}

// SaveArticle godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Article
// @Summary Save article
// @Description Save article
// @Accept  json
// @Produce  json
// @Param id path int true "Article ID"
// @Param q body viewmodels.ArticleDto true "New Article Data"
// @Success 201 {object} viewmodels.ArticleDto
// @Failure 400 {object} string
// @Router /api/v1/articles/{id} [post]
func SaveArticle(c *fiber.Ctx) error {
	var (
		data      viewmodels.ArticleDto
		imageFile *multipart.FileHeader
		image     string
		fileName  string
		err       error
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

	if err = c.BodyParser(&data); err != nil {
		return err
	}

	data.CreatedBy = authData.ID
	data.UpdatedBy = authData.ID

	if err = articleService.Save(db, &data); err != nil {
		return err
	}

	fileName = data.Id
	imageFile, err = c.FormFile("image")
	if err == nil {
		if image, err = utils.UploadImageJPG(c, imageFile, fileName); err == nil {
			data.Image = image
			if image, err = utils.UploadImageThumbJPG(imageFile, fileName); err == nil {
				data.Thumbnail = image
			}
		}
	}

	if data.Image != "" || data.Thumbnail != "" {
		if err = articleService.Update(db, data.Id, &data); err != nil {
			return err
		}
	}

	c.Status(fiber.StatusAccepted)
	err = c.JSON(&data)
	return err
}

// DeleteArticle godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Article
// @Summary Delete one article by id
// @Description Delete one article by id
// @Param id path int true "Article ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} viewmodels.ArticleDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/article/{id} [delete]
func DeleteArticle(c *fiber.Ctx) error {
	var (
		data  viewmodels.ArticleDto
		image string
		err   error
		id    string
		db    *gorm.DB
		ok    bool

		authData *viewmodels.AuthorizationModel
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if authData, ok = viewmodels.GetAuthorizationData(c); !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if id = c.Params("id"); id == "" {
		return errors.New("field ID wajib diisi")
	}

	if err = articleService.FindById(db, id, &data); err != nil {
		log.Println(err.Error())
		return c.SendStatus(fiber.StatusNotFound)
	}

	if data.CreatedBy != authData.ID && authData.Role != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	data.UpdatedBy = authData.ID
	if err = articleService.Delete(db, id, &data); err != nil {
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
