package utils

import (
	"log"
	"net/url"
	"os"
)

func GetBasePath() *url.URL {
	var (
		basePath *url.URL
		err      error
	)
	basePath, err = url.Parse(os.Getenv("BASE_PATH_MAIN"))
	if err != nil {
		log.Println(err.Error())
	}
	return basePath
}
