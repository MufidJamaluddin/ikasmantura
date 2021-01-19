package temp_user

import (
	"backend/models"
	"backend/repository"
	"backend/viewmodels"
	"gorm.io/gorm"
)

func Verify(db *gorm.DB, id uint, out *viewmodels.UserDto) error {
	var (
		err            error
		tempModel      models.TempUser
		permanentModel models.User
		session        *gorm.DB
	)

	out.Id = int(id)

	toTempModel(out, &tempModel)

	session = db.Session(&gorm.Session{SkipDefaultTransaction: false, FullSaveAssociations: true})

	err = session.Transaction(func(tx *gorm.DB) error {
		var errorTransact error

		if errorTransact = tx.Model(&tempModel).
			Where("id = ?", int(id)).
			Updates(&tempModel).Error; errorTransact != nil {
			return errorTransact
		}

		if errorTransact = tx.Model(&tempModel).
			Preload("Address").
			Preload("Classrooms").
			First(&tempModel, "id = ?", int(id)).Error; errorTransact != nil {
			return errorTransact
		}

		toPermanentModel(&tempModel, &permanentModel)

		if errorTransact = tx.Model(&permanentModel).
			Create(&permanentModel).Error; errorTransact != nil {
			return errorTransact
		}

		errorTransact = tx.Model(&tempModel).Delete(&tempModel).Error

		return errorTransact
	})

	return err
}

func Update(db *gorm.DB, id uint, out *viewmodels.UserDto) error {
	var (
		err   error
		model models.TempUser
	)

	out.Id = int(id)

	toTempModel(out, &model)
	err = repository.Update(db.Omit("id", "username", "password"), &model)
	return err
}

func Save(db *gorm.DB, out *viewmodels.UserDto) error {
	var (
		err   error
		model models.TempUser
	)

	toTempModel(out, &model)
	if err = db.Create(&model).Error; err == nil {
		out.Id = int(model.ID)
	}
	return err
}

func Delete(db *gorm.DB, id uint, out *viewmodels.UserDto) error {
	var (
		err   error
		model models.TempUser
	)

	model.ID = id

	if err = repository.Delete(db, &model); err == nil {
		toViewModel(&model, out)
	}
	return err
}
