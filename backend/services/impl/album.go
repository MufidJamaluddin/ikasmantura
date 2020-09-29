package impl

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
	"backend/viewmodels"
	"database/sql"
	"gorm.io/gorm"
)

var albumSearchFields []string

func init() {
	albumSearchFields = []string{
		"title",
	}
}

type AlbumServiceImpl struct {
	DB *gorm.DB
}

func (p *AlbumServiceImpl) toModel(data *viewmodels.AlbumDto, out *models.Album) {
	out.ID = data.Id
	out.Title = data.Title

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func (p *AlbumServiceImpl) toData(in *models.Album, out *viewmodels.AlbumDto) {
	out.Id = in.ID
	out.Title = in.Title

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}

func (p *AlbumServiceImpl) GetTotal(search *viewmodels.AlbumParam) (uint, error) {
	var (
		err   error
		model models.Album
		tx    *gorm.DB
		total int64
	)

	tx = p.DB.Model(&model)

	search.Filter(tx, articleSearchFields)

	if err = tx.Count(&total).Error; err != nil {
		return 0, err
	}

	return uint(total), err
}

func (p *AlbumServiceImpl) Find(search *viewmodels.AlbumParam, callback func(*viewmodels.AlbumDto)) error {
	var (
		err   error
		model models.Album
		tx    *gorm.DB
		rows  *sql.Rows
	)

	tx = p.DB.Model(&model)
	search.Filter(tx, albumSearchFields)

	if rows, err = tx.Rows(); err != nil {
		return err
	}

	for rows.Next() {
		err = tx.ScanRows(rows, &model)

		p.toData(&model, &search.AlbumDto)
		callback(&search.AlbumDto)
	}

	return err
}

func (p *AlbumServiceImpl) FindById(id uint, out *viewmodels.AlbumDto) error {
	var (
		err   error
		model models.Album
	)

	if err = repository.FindById(p.DB, id, &model); err == nil {
		p.toData(&model, out)
	}
	return err
}

func (p *AlbumServiceImpl) Update(id uint, data *viewmodels.AlbumDto) error {
	var (
		err   error
		model models.Album
	)

	data.Id = id

	p.toModel(data, &model)
	err = repository.Update(p.DB, &model)
	return err
}

func (p *AlbumServiceImpl) Save(out *viewmodels.AlbumDto) error {
	var (
		err   error
		model models.Album
	)

	p.toModel(out, &model)
	if err = repository.Save(p.DB, &model); err == nil {
		out.Id = model.ID
	}
	return err
}

func (p *AlbumServiceImpl) Delete(id uint, out *viewmodels.AlbumDto) error {
	var (
		err   error
		model models.Album
		hist  models.AlbumHistory
	)

	model.ID = id

	if err = repository.Delete(p.DB, &model); err == nil {
		p.toData(&model, out)

		p.DB.Model(&hist).
			Where("ID = ?", id).
			Order("version desc").
			Last(&hist)

		if hist.Action == "delete" {
			hist.UpdatedBy = model.UpdatedBy
			p.DB.Save(&hist)
		}
	}
	return err
}
