package main

import (
	"backend/models"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.UserHistory{},
		&models.UserAddress{},
		&models.UserAddressHistory{},

		&models.UserLogin{},

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

	log.Println(err)
}
