package controllers

import (
	"backend/services"
	"backend/utils"
	"backend/viewmodels"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// @author Mufid Jamaluddin
type AboutController struct {
	Service services.AboutService
}

// GetAbout godoc
// @Tags Web Info
// @Summary Get Web About
// @Description About of IKA SMAN Situraja Website
// @Param  id path int true "About ID (default 1)"
// @Accept json
// @Produce json
// @Success 200	{object} viewmodels.AboutDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/about/{id} [get]
func (p *AboutController) GetAbout(c *fiber.Ctx) error {
	var (
		data viewmodels.AboutDto
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

// UpdateAbout godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Web Info
// @Summary Update Web About
// @Description Update About of IKA SMAN Situraja Website
// @Param  id path int true "About ID (default 1)"
// @Accept json
// @Produce json
// @Success 200	{object} viewmodels.AboutDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/about/{id} [get]
func (p *AboutController) UpdateAbout(c *fiber.Ctx) error {
	var (
		data viewmodels.AboutDto
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
