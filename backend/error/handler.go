package error

import (
	"github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
	"log"
)

// @author Mufid Jamaluddin
func CustomErrorHandler(c *fiber.Ctx, err error) error {

	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	log.Println(err.Error())

	if ce, ok := err.(*CustomError); ok {
		if ce.IsCanSendToClient() {
			return c.SendStatus(code)
		}
	} else {
		if stack, ok := err.(*errors.Error); ok {
			log.Println(stack.ErrorStack())
		}
	}

	return c.Status(code).SendString(err.Error())
}
