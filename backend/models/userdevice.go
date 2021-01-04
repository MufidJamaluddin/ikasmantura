package models

import (
	"database/sql/driver"
	"gorm.io/gorm/utils"
)

type deviceType string

const (
	mobile = "mobile"
	tablet = "tablet"
	desktop = "desktop"
	bot = "bot"
)

func (d *deviceType) Scan(value interface{}) error {
	*d = deviceType(value.([]byte))
	return nil
}

func (d deviceType) Value() (driver.Value, error) {
	return utils.ToString(d), nil
}
