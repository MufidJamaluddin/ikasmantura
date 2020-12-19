package event

import (
	eventService "backend/services/event"
	"backend/utils"
	"backend/viewmodels"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

// @author Mufid Jamaluddin

// SearchEvent godoc
// @Tags Event Management
// @Summary Search Event
// @Description Search Event
// @Accept  json
// @Produce  json
// @Param q query viewmodels.EventParam true "Pagination Options"
// @Success 200 {object} []viewmodels.EventDto
// @Failure 400 {object} string
// @Router /api/v1/events [get]
func SearchEvent(c *fiber.Ctx) error {
	var (
		data      viewmodels.EventParam
		err       error
		total     uint
		callback  func(eventDto *viewmodels.EventDto)
		isStarted bool
		db        *gorm.DB
		ok        bool
		counter   uint

		authData *viewmodels.AuthorizationModel
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err = c.QueryParser(&data); err != nil {
		return err
	}

	if authData, ok = viewmodels.GetAuthorizationData(c); ok {
		data.CurrentUserId = authData.ID
	}

	if total, err = eventService.GetTotal(db, &data); err != nil {
		return err
	}

	c.Response().Header.Add("X-Total-Count", fmt.Sprintf("%v", total))

	// RESPONSE ARRAY JSON DATA
	// HEMAT MEMORY, NGGAK PERLU ALOKASI ARRAY, KIRIM AJA KE CLIENT SECARA MENGALIR
	isStarted = false
	counter = data.GetParams.Start
	callback = func(dt *viewmodels.EventDto) {
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
	err = eventService.Find(db, &data, callback)
	_, err = c.Write(utils.ToBytes("]"))
	// END RESPONSE ARRAY JSON DATA

	if data.GetParams.Start < counter {
		c.Response().Header.Add("Content-Range",
			fmt.Sprintf("items %v-%v/%v", data.GetParams.Start, counter, total))

		if total == counter {
			c.Response().Header.SetStatusCode(fiber.StatusOK)
		} else {
			c.Response().Header.SetStatusCode(fiber.StatusPartialContent)
		}
	}

	return err
}

// GetOneEvent godoc
// @Tags Event Management
// @Summary Get one data by id
// @Description Get data by id
// @Param id path int true "Event ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} viewmodels.EventDto
// @Failure 400 {object} string
// @Router /api/v1/events/{id} [get]
func GetOneEvent(c *fiber.Ctx) error {
	var (
		data     viewmodels.EventDto
		err      error
		id       uint
		jsonData []byte
		db       *gorm.DB
		ok       bool

		authData *viewmodels.AuthorizationModel
	)

	if db, ok = c.Locals("db").(*gorm.DB); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if id, err = utils.ToUint(c.Params("id")); err != nil {
		return err
	}

	if authData, ok = viewmodels.GetAuthorizationData(c); ok {
		data.CurrentUserId = authData.ID
	}

	if err = eventService.FindById(db, id, &data); err != nil {
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
// @Param q body viewmodels.EventDto true "New Event Data"
// @Success 202 {object} viewmodels.EventDto
// @Failure 400 {object} string
// @Router /api/v1/events/{id} [put]
func UpdateEvent(c *fiber.Ctx) error {
	var (
		data      viewmodels.EventDto
		image     string
		fileName  string
		imageFile *multipart.FileHeader
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

	if err = c.BodyParser(&data); err != nil {
		return err
	}

	data.UpdatedBy = authData.ID
	if err = eventService.Update(db, id, &data); err != nil {
		return err
	}

	imageFile, err = c.FormFile("image")
	fileName = strconv.Itoa(int(data.Id))
	if err == nil {
		if image, err = utils.UploadImageJPG(c, imageFile, fileName); err == nil {
			data.Image = image
			if image, err = utils.UploadImageThumbJPG(imageFile, fileName); err == nil {
				data.Thumbnail = image
			}
		}
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
// @Param q body viewmodels.EventDto true "New Event Data"
// @Success 202 {object} viewmodels.EventDto
// @Failure 400 {object} string
// @Router /api/v1/events [post]
func SaveEvent(c *fiber.Ctx) error {
	var (
		data      viewmodels.EventDto
		fileName  string
		image     string
		imageFile *multipart.FileHeader
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

	data.UpdatedBy = authData.ID
	data.CreatedBy = authData.ID
	if err = eventService.Save(db, &data); err != nil {
		return err
	}

	imageFile, err = c.FormFile("image")
	fileName = strconv.Itoa(int(data.Id))
	if err == nil {
		if image, err = utils.UploadImageJPG(c, imageFile, fileName); err == nil {
			data.Image = image
			if image, err = utils.UploadImageThumbJPG(imageFile, fileName); err == nil {
				data.Thumbnail = image
			}
		}
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
// @Success 202 {object} viewmodels.EventDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/event/{id} [delete]
func DeleteEvent(c *fiber.Ctx) error {
	var (
		data  viewmodels.EventDto
		image string
		err   error
		id    uint
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

	if id, err = utils.ToUint(c.Params("id")); err != nil {
		return err
	}

	data.UpdatedBy = authData.ID
	if err = eventService.Delete(db, id, &data); err != nil {
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
func RegisterEvent(c *fiber.Ctx) error {
	var (
		err error
		id  uint
		db  *gorm.DB
		ok  bool

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

	if err = eventService.RegisterEvent(db, id, authData.ID); err == nil {
		c.Status(fiber.StatusAccepted)
		err = c.SendString("{\"status\":\"OK\"}")
	}
	return err
}

func DownloadEventTicket(c *fiber.Ctx) error {
	var (
		err error
		id  uint

		authData *viewmodels.AuthorizationModel

		userEventData viewmodels.UserEventDetailDto
		pdfGen        *wkhtmltopdf.PDFGenerator
		htmlBuf       bytes.Buffer

		db *gorm.DB
		ok bool
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

	userEventData.UserId = authData.ID
	userEventData.EventId = id
	if err = eventService.GetUserEvent(db, &userEventData); err != nil {
		return err
	}

	err = utils.HtmlTemplates.ExecuteTemplate(
		&htmlBuf, "event_ticket.html", &userEventData)

	if err != nil {
		return err
	}

	if pdfGen, err = wkhtmltopdf.NewPDFGenerator(); err == nil {

		pdfGen.AddPage(wkhtmltopdf.NewPageReader(&htmlBuf))
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
