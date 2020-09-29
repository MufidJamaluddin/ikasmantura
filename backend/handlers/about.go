package handlers

import (
	"backend/dto"
	"backend/services"
	"backend/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type AboutHandler struct {
	Service services.AboutService
}

// GetAbout godoc
// @Tags Web Info
// @Summary Get Web About
// @Description About of IKA SMAN Situraja Website
// @Param  id path int true "About ID (default 1)"
// @Accept json
// @Produce json
// @Success 200	{object} dto.AboutDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/about/{id} [get]
func (p *AboutHandler) GetAbout(c *fiber.Ctx) error {
	var (
		data dto.AboutDto
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
// @Success 200	{object} dto.AboutDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/about/{id} [get]
func (p *AboutHandler) UpdateAbout(c *fiber.Ctx) error {
	var (
		data dto.AboutDto
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
