package impl

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
	"backend/viewmodels"
	"database/sql"
	"gorm.io/gorm"
)

var albumPhotoSearchFields []string

func init() {
	albumPhotoSearchFields = []string{
		"title", "album_id",
	}
}

type AlbumPhotoServiceImpl struct {
	DB *gorm.DB
}

func (p *AlbumPhotoServiceImpl) toModel(data *viewmodels.AlbumPhotoDto, out *models.AlbumPhoto) {
	out.ID = data.Id
	out.Title = data.Title
	out.Image = data.Image
	out.Thumbnail = data.Thumbnail

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func (p *AlbumPhotoServiceImpl) toData(in *models.AlbumPhoto, out *viewmodels.AlbumPhotoDto) {
	out.Id = in.ID
	out.Title = in.Title
	out.Image = in.Image
	out.Thumbnail = in.Thumbnail

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}

func (p *AlbumPhotoServiceImpl) GetTotal(search *viewmodels.AlbumPhotoParam) (uint, error) {
	var (
		err   error
		model models.AlbumPhoto
		tx    *gorm.DB
		total int64
	)

	tx = p.DB.Model(&model)

	search.Filter(tx, albumPhotoSearchFields)

	if err = tx.Count(&total).Error; err != nil {
		return 0, err
	}

	return uint(total), err
}

func (p *AlbumPhotoServiceImpl) Find(search *viewmodels.AlbumPhotoParam, callback func(*viewmodels.AlbumPhotoDto)) error {
	var (
		err   error
		model models.AlbumPhoto
		tx    *gorm.DB
		rows  *sql.Rows
	)

	tx = p.DB.Model(&model)
	search.Filter(tx, albumPhotoSearchFields)

	if rows, err = tx.Rows(); err != nil {
		return err
	}

	for rows.Next() {
		err = tx.ScanRows(rows, &model)

		p.toData(&model, &search.AlbumPhotoDto)
		callback(&search.AlbumPhotoDto)
	}

	return err
}

func (p *AlbumPhotoServiceImpl) FindById(id uint, out *viewmodels.AlbumPhotoDto) error {
	var (
		err   error
		model models.AlbumPhoto
	)

	if err = repository.FindById(p.DB, id, &model); err == nil {
		p.toData(&model, out)
	}
	return err
}

func (p *AlbumPhotoServiceImpl) Update(id uint, out *viewmodels.AlbumPhotoDto) error {
	var (
		err   error
		model models.AlbumPhoto
	)

	out.Id = id

	p.toModel(out, &model)
	err = repository.Update(p.DB, &model)
	return err
}

func (p *AlbumPhotoServiceImpl) Save(out *viewmodels.AlbumPhotoDto) error {
	var (
		err   error
		model models.AlbumPhoto
	)

	p.toModel(out, &model)
	if err = repository.Save(p.DB, &model); err == nil {
		out.Id = model.ID
	}
	return err
}

func (p *AlbumPhotoServiceImpl) Delete(id uint, out *viewmodels.AlbumPhotoDto) error {
	var (
		err   error
		model models.AlbumPhoto
	)

	model.ID = id

	if err = repository.Delete(p.DB, &model); err == nil {
		p.toData(&model, out)
	}
	return err
}
