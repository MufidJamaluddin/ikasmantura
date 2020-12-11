package articletopic

import (
	"backend/models"
	"backend/repository"
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

func GetTotal(db *gorm.DB, search *viewmodels.ArticleTopicParam) (uint, error) {
	var (
		err   error
		model models.ArticleTopic
		tx    *gorm.DB
		total int64
	)

	tx = db.Model(&model)

	search.Filter(tx, articleTopicSearchFields)

	if err = tx.Count(&total).Error; err != nil {
		return 0, err
	}

	return uint(total), err
}

func Find(db *gorm.DB, search *viewmodels.ArticleTopicParam, callback func(*viewmodels.ArticleTopicDto)) error {
	var (
		err   error
		model models.ArticleTopic
		tx    *gorm.DB
		rows  *sql.Rows
	)

	tx = db.Model(&model)

	search.Filter(tx, articleTopicSearchFields)

	if rows, err = tx.Rows(); err != nil {
		return err
	}

	for rows.Next() {
		err = tx.ScanRows(rows, &model)

		toViewModel(&model, &search.ArticleTopicDto)
		callback(&search.ArticleTopicDto)
	}

	return err
}

func FindById(db *gorm.DB, id uint, out *viewmodels.ArticleTopicDto) error {
	var (
		err   error
		model models.ArticleTopic
	)

	if err = repository.FindById(db, id, &model); err == nil {
		toViewModel(&model, out)
	}
	return err
}
