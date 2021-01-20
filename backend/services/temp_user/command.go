package temp_user

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
	"backend/viewmodels"
	"github.com/go-errors/errors"
	uuid "github.com/satori/go.uuid"
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

func ConfirmEmail(db *gorm.DB, username string, confirmEmailToken utils.UUID) (err error)  {
	var (
		user models.TempUser
		tx *gorm.DB
	)

	tx = db.Model(&user)
	tx = tx.Omit("id", "username", "password")

	if err = tx.First(&user, "username = ?", username).Error; err != nil {
		return err
	}

	if user.ConfirmEmailToken != confirmEmailToken {
		err = errors.Errorf("email confirmation for username %s is invalid!", username)
	}

	user.EmailValid = true

	err = tx.Save(&user).Error
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
	model.ConfirmEmailToken = utils.UUID(uuid.NewV4())

	if err = db.Create(&model).Error; err == nil {
		out.Id = int(model.ID)
		out.ConfirmEmailToken = utils.ToBase64UUID(model.ConfirmEmailToken)
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
