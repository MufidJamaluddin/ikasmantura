package album

import (
	"backend/models"
	"backend/repository"
	"backend/viewmodels"
	"gorm.io/gorm"
)

func Update(db *gorm.DB, id uint, data *viewmodels.AlbumDto) error {
	var (
		err   error
		model models.Album
	)

	data.Id = int(id)

	toModel(data, &model)

	err = db.Transaction(func(tx *gorm.DB) error {
		return repository.Update(tx, &model)
	})

	return err
}

func Save(db *gorm.DB, out *viewmodels.AlbumDto) error {
	var (
		err   error
		model models.Album
	)

	toModel(out, &model)
	err = db.Transaction(func(tx *gorm.DB) error {
		var errT error
		if errT = repository.Save(tx, &model); err == nil {
			out.Id = int(model.ID)
		}
		return errT
	})
	return err
}

func Delete(db *gorm.DB, id uint, out *viewmodels.AlbumDto) error {
	var (
		err   error
		model models.Album
	)

	model.ID = id

	err = db.Transaction(func(tx *gorm.DB) error {
		var errT error
		if errT = repository.Delete(tx, &model); err == nil {
			toViewModel(&model, out)
		}
		return errT
	})

	return err
}
