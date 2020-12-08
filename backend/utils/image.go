package utils

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/liujiawm/graphics-go/graphics"
	"image"
	"image/jpeg"
	"mime/multipart"
	"os"
	"unicode/utf8"
)

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func UploadImageJPG(c *fiber.Ctx, imageFile *multipart.FileHeader) (string, error) {
	var (
		err       error
		imagePath string
	)

	if imageFile.Header["Content-Type"][0] != "image/jpeg" {
		return "", errors.New("image isn't valid")
	}

	imagePath = fmt.Sprintf(
		"./%s/%s/%s", os.Getenv("ASSET_PATH"), "img", imageFile.Filename)

	if err = c.SaveFile(imageFile, imagePath); err != nil {
		return trimFirstRune(imagePath), err
	}

	return "", nil
}

func UploadImageThumbJPG(imageFile *multipart.FileHeader) (string, error) {
	var (
		err                 error
		thumbPathStr        string
		imagePath           multipart.File
		originalSourceImage image.Image
		thumbSourceImage    *image.RGBA
	)

	if imageFile.Header["Content-Type"][0] != "image/jpeg" {
		return "", errors.New("image isn't valid")
	}

	thumbPathStr = fmt.Sprintf(
		"./%s/%s/%s", os.Getenv("ASSET_PATH"), "thumb", imageFile.Filename)

	if imagePath, err = imageFile.Open(); imagePath == nil {
		return "", errors.New("no file")
	}
	defer imagePath.Close()

	if err != nil {
		return "", err
	}

	originalSourceImage, _, _ = image.Decode(imagePath)

	thumbSourceImage = image.NewRGBA(image.Rect(0, 0, 80, 80))

	if err = graphics.Thumbnail(thumbSourceImage, originalSourceImage); err != nil {
		return "", err
	}

	newImage, _ := os.Create(thumbPathStr)
	defer newImage.Close()

	err = jpeg.Encode(newImage, thumbSourceImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
	return trimFirstRune(thumbPathStr), err
}
