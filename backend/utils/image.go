package utils

import (
	"fmt"
	"github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/liujiawm/graphics-go/graphics"
	"image"
	"image/jpeg"
	"mime/multipart"
	"os"
	"path/filepath"
	"unicode/utf8"
)

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func UploadImageJPG(c *fiber.Ctx, imageFile *multipart.FileHeader, fileName string) (string, error) {
	var (
		err       error
		imagePath string
	)

	if imageFile.Header["Content-Type"][0] != "image/jpeg" {
		return "", errors.New("image isn't valid")
	}

	fileName = fmt.Sprintf("%s%s", fileName, filepath.Ext(imageFile.Filename))

	imagePath = fmt.Sprintf(
		"./%s/%s/%s", os.Getenv("ASSET_PATH"), "img", fileName)

	if err = c.SaveFile(imageFile, imagePath); err == nil {
		return trimFirstRune(imagePath), err
	}

	return "", nil
}

func UploadImageThumbJPG(imageFile *multipart.FileHeader, fileName string) (string, error) {
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

	fileName = fmt.Sprintf("%s%s", fileName, filepath.Ext(imageFile.Filename))

	thumbPathStr = fmt.Sprintf(
		"./%s/%s/%s", os.Getenv("ASSET_PATH"), "thumb", fileName)

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
