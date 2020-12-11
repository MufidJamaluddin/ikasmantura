package article

import (
	"backend/models"
	"backend/repository"
	"backend/viewmodels"
	"gorm.io/gorm"
)

func Update(db *gorm.DB, id uint, out *viewmodels.ArticleDto) error {
	var (
		err   error
		model models.Article
	)

	out.Id = id

	toModel(out, &model)
	err = repository.Update(db, &model)
	return err
}

func Save(db *gorm.DB, out *viewmodels.ArticleDto) error {
	var (
		err   error
		model models.Article
	)

	toModel(out, &model)
	if err = repository.Save(db, &model); err == nil {
		out.Id = model.ID
	}
	return err
}

func Delete(db *gorm.DB, id uint, out *viewmodels.ArticleDto) error {
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
