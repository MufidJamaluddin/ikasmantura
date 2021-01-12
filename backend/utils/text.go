package utils

import (
	"fmt"
	"regexp"
)

var spacePattern = regexp.MustCompile(`\s+`)

func ToLikeSQL(text string) (sql string) {
	sql = spacePattern.ReplaceAllString(text, "%")
	sql = fmt.Sprintf("%s%s%s", "%", sql, "%")
	return
}
