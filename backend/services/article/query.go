package article

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
	"database/sql"
	"gorm.io/gorm"
	"log"
	"strings"
)

var articleSearchFields []string

func init() {
	articleSearchFields = []string{
		"title",
	}
}

func GetTotal(db *gorm.DB, search *viewmodels.ArticleParam) (uint, error) {
	var (
		err   error
		model models.Article
		tx    *gorm.DB
		total int64 = 0
	)

	tx = db.Model(&model)

	searchFilter(tx, search, false)

	if err = tx.Count(&total).Error; err != nil {
		return 0, err
	}

	return uint(total), err
}

func searchFilter(tx *gorm.DB, search *viewmodels.ArticleParam, withLimit bool) {
	var title string

	search.Filter(tx, articleSearchFields, withLimit)

	if search.StartFrom != nil {
		tx.Where("start >= ?", search.StartFrom)
	}

	if search.EndTo != nil {
		tx.Where("end <= ?", search.EndTo)
	}

	title = strings.Trim(search.Title, " ")
	if title != "" {
		title = utils.ToLikeSQL(title)
		tx.Where("title LIKE ?", title)
	}

	if search.TopicId != 0 {
		tx.Where("article_topic_id = ?", search.TopicId)
	}
}

func Find(db *gorm.DB, search *viewmodels.ArticleParam, callback func(*viewmodels.ArticleDto)) error {
	var (
		err   error
		model models.Article
		tx    *gorm.DB
		rows  *sql.Rows
	)

	tx = db.Model(&model)

	searchFilter(tx, search, true)

	tx.Select(
		"id, title, SUBSTRING(body, 1, 50) as body, thumbnail, image, " +
			"created_by, created_at, updated_by, updated_at",
	)

	if rows, err = tx.Rows(); err != nil {
		return err
	}

	for rows.Next() {
		err = tx.ScanRows(rows, &model)

		toViewModel(&model, &search.ArticleDto)
		callback(&search.ArticleDto)
	}

	return err
}

func FindById(db *gorm.DB, id string, out *viewmodels.ArticleDto) error {
	var (
		err     error
		model   models.Article
		user    models.User
		uid     utils.UUID
		session *gorm.DB
	)

	if uid, err = utils.FromBase64UUID(id); err != nil {
		return err
	}

	session = db.Session(&gorm.Session{SkipDefaultTransaction: false})

	if err = session.Model(&model).Table("articles").
		Where("id = ?", uid).
		First(&model).Error; err == nil {
		toViewModel(&model, out)

		if err = session.Model(&user).Table("users").
			First(&user, "id = ?", out.CreatedBy).Error; err == nil {
			out.CreatedByName = user.Name
		} else {
			log.Println(err.Error())
		}

		err = nil
	}

	return err
}
