package main

import (
	"backend/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	var err error

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{

	})

	m.InitSchema(func(tx *gorm.DB) error {
		return tx.AutoMigrate(
			&models.User{},
			&models.UserHistory{},
			&models.UserAddress{},
			&models.UserAddressHistory{},

			&models.TempUser{},
			&models.TempUserAddress{},

			&models.About{},
			&models.AboutHistory{},

			&models.Album{},
			&models.AboutHistory{},
			&models.AlbumPhoto{},

			&models.Article{},
			&models.ArticleTopic{},

			&models.Classroom{},
			&models.ClassroomHistory{},

			&models.UserClassroom{},

			&models.TempUserClassroom{},

			&models.Department{},
			&models.DepartmentHistory{},

			&models.Event{},
			&models.UserEvent{})
	})

	err = m.Migrate()

	log.Println(err)
}
