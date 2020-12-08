package impl

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
	"backend/viewmodels"
	"database/sql"
	"gorm.io/gorm"
)

var articleTopicSearchFields []string

func init() {
	articleTopicSearchFields = []string{
		"name",
	}
}

type ArticleTopicServiceImpl struct {
	DB *gorm.DB
}

func (p *ArticleTopicServiceImpl) toModel(data *viewmodels.ArticleTopicDto, out *models.ArticleTopic) {
	out.ID = data.Id
	out.Name = data.Name
	out.Icon = data.Icon
	out.Description = data.Description

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func (p *ArticleTopicServiceImpl) toData(in *models.ArticleTopic, out *viewmodels.ArticleTopicDto) {
	out.Id = in.ID
	out.Name = in.Name
	out.Icon = in.Icon
	out.Description = in.Description

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}

func (p *ArticleTopicServiceImpl) GetTotal(search *viewmodels.ArticleTopicParam) (uint, error) {
	var (
		err   error
		model models.ArticleTopic
		tx    *gorm.DB
		total int64
	)

	tx = p.DB.Model(&model)

	search.Filter(tx, articleTopicSearchFields)

	if err = tx.Count(&total).Error; err != nil {
		return 0, err
	}

	return uint(total), err
}

func (p *ArticleTopicServiceImpl) Find(search *viewmodels.ArticleTopicParam, callback func(*viewmodels.ArticleTopicDto)) error {
	var (
		err   error
		model models.ArticleTopic
		tx    *gorm.DB
		rows  *sql.Rows
	)

	tx = p.DB.Model(&model).Select([]string{
		"id", "userId", "title", "SUBSTRING(body, 1, 20) as body", "image",
		"createdBy", "createdAt", "updatedBy", "updatedAt",
	})

	search.Filter(tx, articleSearchFields)

	if rows, err = tx.Rows(); err != nil {
		return err
	}

	for rows.Next() {
		err = tx.ScanRows(rows, &model)

		p.toData(&model, &search.ArticleTopicDto)
		callback(&search.ArticleTopicDto)
	}

	return err
}

func (p *ArticleTopicServiceImpl) FindById(id uint, out *viewmodels.ArticleTopicDto) error {
	var (
		err   error
		model models.ArticleTopic
	)

	if err = repository.FindById(p.DB, id, &model); err == nil {
		p.toData(&model, out)
	}
	return err
}

func (p *ArticleTopicServiceImpl) Update(id uint, out *viewmodels.ArticleTopicDto) error {
	var (
		err   error
		model models.ArticleTopic
	)

	out.Id = id

	p.toModel(out, &model)
	err = repository.Update(p.DB, &model)
	return err
}

func (p *ArticleTopicServiceImpl) Save(out *viewmodels.ArticleTopicDto) error {
	var (
		err   error
		model models.ArticleTopic
	)

	p.toModel(out, &model)
	if err = repository.Save(p.DB, &model); err == nil {
		out.Id = model.ID
	}
	return err
}

func (p *ArticleTopicServiceImpl) Delete(id uint, out *viewmodels.ArticleTopicDto) error {
	var (
		err   error
		model models.ArticleTopic
		hist  models.ArticleTopicHistory
	)

	out.Id = id

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
