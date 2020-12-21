package repository

import (
	"gorm.io/gorm"
)

func FindById(db *gorm.DB, id uint, out interface{}) error {
	return db.First(out, id).Error
}

func Save(db *gorm.DB, model interface{}) error {
	db.Create(model)
	return db.Save(model).Error
}

func Update(db *gorm.DB, model interface{}) error {
	return db.Updates(model).Error
}

func Delete(db *gorm.DB, model interface{}) error {
	return db.Delete(model).Error
}
