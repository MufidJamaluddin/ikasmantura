package about

import (
	aboutService "backend/services/about"
	"backend/utils"
	"backend/viewmodels"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @author Mufid Jamaluddin

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
func GetAbout(c *fiber.Ctx) error {
	var (
		data viewmodels.AboutDto
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

	if err = aboutService.FindById(db, id, &data); err != nil {
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
// @Router /api/v1/about/{id} [put]
func UpdateAbout(c *fiber.Ctx) error {
	var (
		data viewmodels.AboutDto
		id   uint
		err  error
		db   *gorm.DB
		ok   bool

		authData *viewmodels.AuthorizationModel
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if authData, ok = viewmodels.GetAuthorizationData(c); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if id, err = utils.ToUint(c.Params("id")); err != nil {
		return err
	}

	if err = c.BodyParser(&data); err != nil {
		return err
	}

	data.UpdatedBy = authData.ID

	if err = aboutService.Update(db, id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusAccepted)

	err = c.JSON(&data)
	return err
}
