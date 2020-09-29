package impl

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
	"backend/viewmodels"
	"gorm.io/gorm"
)

type AboutServiceImpl struct {
	DB *gorm.DB
}

func (p *AboutServiceImpl) toModel(data *viewmodels.AboutDto, out *models.About) {
	out.ID = data.Id
	out.Description = data.Description
	out.Mission = data.Mission
	out.Vision = data.Vision

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func (p *AboutServiceImpl) toData(in *models.About, out *viewmodels.AboutDto) {
	out.Id = in.ID
	out.Description = in.Description
	out.Mission = in.Mission
	out.Vision = in.Vision

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}

func (p *AboutServiceImpl) FindById(id uint, out *viewmodels.AboutDto) error {
	var (
		err   error
		model models.About
	)

	if err = repository.FindById(p.DB, id, &model); err == nil {
		p.toData(&model, out)
	}

	return err
}

func (p *AboutServiceImpl) Update(id uint, out *viewmodels.AboutDto) error {
	var (
		err   error
		model models.About
	)

	out.Id = id
	p.toModel(out, &model)
	err = repository.Update(p.DB, &model)
	return err
}
