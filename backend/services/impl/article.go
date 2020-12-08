package impl

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
	"backend/viewmodels"
	"database/sql"
	"gorm.io/gorm"
)

var articleSearchFields []string

func init() {
	articleSearchFields = []string{
		"title", "article_topic_id",
	}
}

type ArticleServiceImpl struct {
	DB *gorm.DB
}

func (p *ArticleServiceImpl) toModel(data *viewmodels.ArticleDto, out *models.Article) {
	out.ID = data.Id
	out.Title = data.Title
	out.Body = data.Body
	out.Image = data.Image
	out.Thumbnail = data.Thumbnail
	out.CreatedBy = data.UserId
	out.UpdatedBy = data.UserId
	out.ArticleTopicId = data.TopicId

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func (p *ArticleServiceImpl) toData(in *models.Article, out *viewmodels.ArticleDto) {
	out.Id = in.ID
	out.Title = in.Title
	out.Body = in.Body
	out.Image = in.Image
	out.Thumbnail = in.Thumbnail
	out.UserId = in.CreatedBy
	out.CreatedAt = in.CreatedAt
	out.TopicId = in.ArticleTopicId

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}

func (p *ArticleServiceImpl) GetTotal(search *viewmodels.ArticleParam) (uint, error) {
	var (
		err   error
		model models.Article
		tx    *gorm.DB
		total int64
	)

	tx = p.DB.Model(&model)

	p.searchFilter(tx, search)

	if err = tx.Count(&total).Error; err != nil {
		return 0, err
	}

	return uint(total), err
}

func (p *ArticleServiceImpl) searchFilter(tx *gorm.DB, search *viewmodels.ArticleParam) {
	search.Filter(tx, articleSearchFields)

	if search.StartFrom != nil {
		tx.Where("start >= ?", search.StartFrom)
	}

	if search.EndTo != nil {
		tx.Where("end <= ?", search.EndTo)
	}
}

func (p *ArticleServiceImpl) Find(search *viewmodels.ArticleParam, callback func(*viewmodels.ArticleDto)) error {
	var (
		err   error
		model models.Article
		tx    *gorm.DB
		rows  *sql.Rows
	)

	tx = p.DB.Model(&model).Select([]string{
		"id", "title", "SUBSTRING(body, 1, 20) as body", "thumbnail", "image",
		"created_by", "created_at", "updated_by", "updated_at",
	})

	p.searchFilter(tx, search)

	if rows, err = tx.Rows(); err != nil {
		return err
	}

	for rows.Next() {
		err = tx.ScanRows(rows, &model)

		p.toData(&model, &search.ArticleDto)
		callback(&search.ArticleDto)
	}

	return err
}

func (p *ArticleServiceImpl) FindById(id uint, out *viewmodels.ArticleDto) error {
	var (
		err   error
		model models.Article
		user  models.User
	)

	if err = repository.FindById(p.DB, id, &model); err == nil {
		p.toData(&model, out)

		p.DB.Select("name").First(&user, model.CreatedBy)
		out.CreatedByName = user.Name
	}
	return err
}

func (p *ArticleServiceImpl) Update(id uint, out *viewmodels.ArticleDto) error {
	var (
		err   error
		model models.Article
	)

	out.Id = id

	p.toModel(out, &model)
	err = repository.Update(p.DB, &model)
	return err
}

func (p *ArticleServiceImpl) Save(out *viewmodels.ArticleDto) error {
	var (
		err   error
		model models.Article
	)

	p.toModel(out, &model)
	if err = repository.Save(p.DB, &model); err == nil {
		out.Id = model.ID
	}
	return err
}

func (p *ArticleServiceImpl) Delete(id uint, out *viewmodels.ArticleDto) error {
	var (
		err   error
		model models.Article
	)

	out.Id = id

	p.toModel(out, &model)
	if err = repository.Delete(p.DB, &model); err == nil {
		p.toData(&model, out)
	}
	return err
}
