package article

import (
	"backend/models"
	"backend/repository"
	"backend/viewmodels"
	"gorm.io/gorm"
	"log"
)

func Update(db *gorm.DB, id string, out *viewmodels.ArticleDto) error {
	var (
		model models.Article
		tx    *gorm.DB
	)

	out.Id = id

	toModel(out, &model)
	tx = db.Model(&model).Where("id = ?", model.ID).Updates(&model)
	if tx.Error != nil {
		if tx.Statement != nil {
			log.Println(tx.Statement.SQL.String())
		}
		log.Println(tx.Error.Error())
	}
	return tx.Error
}

func Save(db *gorm.DB, out *viewmodels.ArticleDto) error {
	var (
		err   error
		model models.Article
	)

	toModel(out, &model)
	if err = repository.Save(db, &model); err == nil {
		toViewModel(&model, out)
	}
	return err
}

func Delete(db *gorm.DB, id string, out *viewmodels.ArticleDto) error {
	var (
		err   error
		model models.Article
	)

	out.Id = id

	toModel(out, &model)
	if err = repository.Delete(db, &model); err == nil {
		toViewModel(&model, out)
	}
	return err
}
