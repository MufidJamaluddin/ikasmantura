package albumphoto

import (
	"backend/models"
	"backend/repository"
	"backend/viewmodels"
	"gorm.io/gorm"
)

func Update(db *gorm.DB, id uint, out *viewmodels.AlbumPhotoDto) error {
	var (
		err   error
		model models.AlbumPhoto
	)

	out.Id = id

	toModel(out, &model)

	err = repository.Update(db, &model)
	return err
}

func Save(db *gorm.DB, out *viewmodels.AlbumPhotoDto) error {
	var (
		err   error
		model models.AlbumPhoto
	)

	toModel(out, &model)

	if err = repository.Save(db, &model); err == nil {
		out.Id = model.ID
	}
	return err
}

func Delete(db *gorm.DB, id uint, out *viewmodels.AlbumPhotoDto) error {
	var (
		err   error
		model models.AlbumPhoto
	)

	model.ID = id

	if err = repository.Delete(db, &model); err == nil {
		toViewModel(&model, out)
	}
	return err
}
