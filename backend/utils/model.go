package utils

import (
	history "github.com/vcraescu/gorm-history/v2"
	"time"
)

type Model interface {
	GetId() uint
}

type ICreated interface {
	SetCreated(userId uint, createdAt time.Time)
	GetCreatedBy() uint
	GetCreatedAt() time.Time
}

type IUpdated interface {
	SetUpdated(userId uint, updatedAt time.Time)
	GetUpdatedBy() uint
	GetUpdatedAt() time.Time
}

type History struct {
	history.Entry
	Action history.Action `gorm:"type:ENUM('create', 'update', 'delete')" gorm-history:"action"`
}
