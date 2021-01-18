package user

import (
	"backend/models"
	"backend/repository"
	"backend/viewmodels"
	"gorm.io/gorm"
)

func Update(db *gorm.DB, id uint, out *viewmodels.UserDto) error {
	var (
		err   error
		model models.User
	)

	out.Id = int(id)

	toModel(out, &model)
	err = repository.Update(db.Omit("id", "username", "password"), &model)
	return err
}

func Save(db *gorm.DB, out *viewmodels.UserDto) error {
	var (
		err   error
		model models.User
	)

	toModel(out, &model)
	if err = repository.Save(db, &model); err == nil {
		out.Id = int(model.ID)
	}
	return err
}

func Delete(db *gorm.DB, id uint, out *viewmodels.UserDto) error {
	var (
		err   error
		model models.User
	)

	model.ID = id

	if err = repository.Delete(db, &model); err == nil {
		toViewModel(&model, out)
	}
	return err
}
