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
type ArticleTopicController struct {
	Service services.ArticleTopicService
}

// SearchArticle godoc
// @Tags Article
// @Summary Search article data
// @Description Get article data with pagination
// @Accept  json
// @Produce  json
// @Param q query viewmodels.ArticleTopicParam true "Pagination Options"
// @Success 200 {object} []viewmodels.ArticleTopicDto
// @Failure 400 {object} string
// @Router /api/v1/articles [get]
func (p *ArticleTopicController) SearchArticle(c *fiber.Ctx) error {
	var (
		data     viewmodels.ArticleTopicParam
		err      error
		total    uint
		callback func(topicDto *viewmodels.ArticleTopicDto)
		counter  uint
	)

	if err = c.QueryParser(&data); err != nil {
		return err
	}

	if total, err = p.Service.GetTotal(&data); err != nil {
		return err
	}

	c.Response().Header.Add("X-Total-Count", fmt.Sprintf("%v", total))

	counter = 0
	callback = func(dt *viewmodels.ArticleTopicDto) {
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

	return err
}

// GetOneArticle godoc
// @Tags Article
// @Summary Get one data by id
// @Description Get data by id
// @Param id path int true "Article ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} viewmodels.ArticleTopicDto
// @Failure 400 {object} string
// @Router /api/v1/articles/{id} [get]
func (p *ArticleTopicController) GetOneArticle(c *fiber.Ctx) error {
	var (
		data viewmodels.ArticleTopicDto
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
// @Param q body viewmodels.ArticleTopicDto true "New Article Data"
// @Success 202 {object} viewmodels.ArticleTopicDto
// @Failure 400 {object} string
// @Router /api/v1/articles/{id} [put]
func (p *ArticleTopicController) UpdateArticle(c *fiber.Ctx) error {
	var (
		data viewmodels.ArticleTopicDto
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
// @Param q body viewmodels.ArticleTopicDto true "New Article Data"
// @Success 201 {object} viewmodels.ArticleTopicDto
// @Failure 400 {object} string
// @Router /api/v1/articles/{id} [put]
func (p *ArticleTopicController) SaveArticle(c *fiber.Ctx) error {
	var (
		data viewmodels.ArticleTopicDto
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

// DeleteArticle godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Article
// @Summary Delete one article by id
// @Description Delete one article by id
// @Param id path int true "Article ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} viewmodels.ArticleTopicDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/article/{id} [delete]
func (p *ArticleTopicController) DeleteArticle(c *fiber.Ctx) error {
	var (
		data viewmodels.ArticleTopicDto
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
