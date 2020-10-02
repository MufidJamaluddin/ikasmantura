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
type ArticleController struct {
	Service services.ArticleService
}

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
func (p *ArticleController) SearchArticle(c *fiber.Ctx) error {
	var (
		data     viewmodels.ArticleParam
		err      error
		total    uint
		callback func(articleDto *viewmodels.ArticleDto)
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
	counter = data.Start
	callback = func(dt *viewmodels.ArticleDto) {
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
	// END RESPONSE ARRAY JSON DATA

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
func (p *ArticleController) GetOneArticle(c *fiber.Ctx) error {
	var (
		data viewmodels.ArticleDto
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
func (p *ArticleController) UpdateArticle(c *fiber.Ctx) error {
	var (
		data viewmodels.ArticleDto
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
// @Router /api/v1/articles/{id} [put]
func (p *ArticleController) SaveArticle(c *fiber.Ctx) error {
	var (
		data viewmodels.ArticleDto
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
func (p *ArticleController) DeleteArticle(c *fiber.Ctx) error {
	var (
		data viewmodels.ArticleDto
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
