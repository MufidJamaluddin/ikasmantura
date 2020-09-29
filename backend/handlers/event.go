package handlers

import (
	"backend/dto"
	"backend/services"
	"backend/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
	Service services.EventService
}

// SearchEvent godoc
// @Tags Event Management
// @Summary Search Event
// @Description Search Event
// @Accept  json
// @Produce  json
// @Param q query dto.EventParam true "Pagination Options"
// @Success 200 {object} []dto.EventDto
// @Failure 400 {object} string
// @Router /api/v1/events [get]
func (p *EventHandler) SearchEvent(c *fiber.Ctx) error {
	var (
		data     dto.EventParam
		err      error
		total    uint
		callback func(eventDto *dto.EventDto)
		counter  uint
	)

	if err = c.QueryParser(&data); err != nil {
		return err
	}

	if user := c.Locals("user"); user != nil {
		if tokens := user.(*jwt.Token); tokens != nil {
			claims := tokens.Claims.(jwt.MapClaims)
			data.CurrentUserId = uint(claims["id"].(float64))
		}
	}

	if total, err = p.Service.GetTotal(&data); err != nil {
		return err
	}

	c.Response().Header.Add("X-Total-Count", fmt.Sprintf("%v", total))

	counter = 0
	callback = func(dt *dto.EventDto) {
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

// GetOneEvent godoc
// @Tags Event Management
// @Summary Get one data by id
// @Description Get data by id
// @Param id path int true "Event ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.EventDto
// @Failure 400 {object} string
// @Router /api/v1/events/{id} [get]
func (p *EventHandler) GetOneEvent(c *fiber.Ctx) error {
	var (
		data dto.EventDto
		err  error
		id   uint
		jsonData []byte
	)

	if id, err = utils.ToUint(c.Params("id")); err != nil {
		return err
	}

	if err = p.Service.FindById(id, &data); err != nil {
		return err
	}

	c.Status(fiber.StatusOK)
	if jsonData, err = json.Marshal(&data); err == nil {
		err = c.Send(jsonData)
	}

	return err
}

// UpdateEvent godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Event Management
// @Summary Update event
// @Description Update event
// @Accept  json
// @Produce  json
// @Param id path int true "Event ID"
// @Param q body dto.EventDto true "New Event Data"
// @Success 202 {object} dto.EventDto
// @Failure 400 {object} string
// @Router /api/v1/events/{id} [put]
func (p *EventHandler) UpdateEvent(c *fiber.Ctx) error {
	var (
		data dto.EventDto
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

// SaveEvent godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Event Management
// @Summary Save event
// @Description Save event
// @Accept  json
// @Produce  json
// @Param q body dto.EventDto true "New Event Data"
// @Success 202 {object} dto.EventDto
// @Failure 400 {object} string
// @Router /api/v1/events [post]
func (p *EventHandler) SaveEvent(c *fiber.Ctx) error {
	var (
		data dto.EventDto
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

// DeleteEvent godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Event Management
// @Summary Delete one event  by id
// @Description Delete one event  by id
// @Param id path int true "Event ID"
// @Accept  json
// @Produce  json
// @Success 202 {object} dto.EventDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/event/{id} [delete]
func (p *EventHandler) DeleteEvent(c *fiber.Ctx) error {
	var (
		data dto.EventDto
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

// RegisterEvent godoc
// @Security BasicAuth
// @Security ApiKeyAuth
// @Tags Event Management
// @Summary Register event
// @Description Register event
// @Accept  json
// @Produce  json
// @Param id path int true "Event ID"
// @Success 202 {object} string
// @Failure 400 {object} string
// @Router /api/v1/event_register/{id} [post]
func (p *EventHandler) RegisterEvent(c *fiber.Ctx) error {
	var (
		err error
		id  uint

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

	if err = p.Service.RegisterEvent(id, currentUserId); err == nil {
		c.Status(fiber.StatusAccepted)
		err = c.SendString("{\"status\":\"OK\"}")
	}
	return err
}

func (p *EventHandler) DownloadEventTicket(c *fiber.Ctx) error {
	var (
		err error
		id  uint

		currentUserId uint

		userEventData dto.UserEventDetailDto
		pdfGen *wkhtmltopdf.PDFGenerator
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

	userEventData.UserId = currentUserId
	userEventData.EventId = id
	if err = p.Service.GetUserEvent(&userEventData); err != nil {
		return err
	}

	htmlBuf := new(bytes.Buffer)

	err = utils.HtmlTemplates.ExecuteTemplate(
		htmlBuf, "event_ticket.html", &userEventData)

	if err != nil {
		return err
	}

	if pdfGen, err = wkhtmltopdf.NewPDFGenerator(); err == nil {

		pdfGen.AddPage(wkhtmltopdf.NewPageReader(htmlBuf))
		pdfGen.Orientation.Set(wkhtmltopdf.OrientationPortrait)
		pdfGen.PageSize.Set(wkhtmltopdf.PageSizeA4)
		pdfGen.Dpi.Set(300)

		c.Response().Reset()

		c.Set("Content-Type", "application/pdf")
		c.Set("Content-Disposition", "attachment;filename=ticket.pdf")

		pdfGen.SetOutput(c.Response().BodyWriter())
	}

	return err
}
