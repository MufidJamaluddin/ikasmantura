package about

import (
	"backend/models"
	"backend/repository"
	"backend/viewmodels"
	"gorm.io/gorm"
)

func FindById(db *gorm.DB, id uint, out *viewmodels.AboutDto) error {
	var (
		err   error
		model models.About
	)

	if err = repository.FindById(db, id, &model); err == nil {
		toViewModel(&model, out)
	}

	return err
}

func Update(db *gorm.DB, id uint, out *viewmodels.AboutDto) error {
	var (
		err   error
		model models.About
	)

	out.Id = int(id)
	toModel(out, &model)

	err = db.Transaction(func(tx *gorm.DB) error {
		return repository.Update(tx, &model)
	})

	return err
}
