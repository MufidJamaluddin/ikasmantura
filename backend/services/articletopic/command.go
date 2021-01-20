package articletopic

import (
	"backend/models"
	"backend/repository"
	"backend/viewmodels"
	"gorm.io/gorm"
)

func Update(db *gorm.DB, id uint, out *viewmodels.ArticleTopicDto) error {
	var (
		err   error
		model models.ArticleTopic
	)

	out.Id = id

	toModel(out, &model)
	err = repository.Update(db, &model)
	return err
}

func Save(db *gorm.DB, out *viewmodels.ArticleTopicDto) error {
	var (
		err   error
		model models.ArticleTopic
	)

	toModel(out, &model)
	if err = repository.Save(db, &model); err == nil {
		out.Id = model.ID
	}
	return err
}

func Delete(db *gorm.DB, id uint, out *viewmodels.ArticleTopicDto) error {
	var (
		err   error
		model models.ArticleTopic
	)

	out.Id = id

	if err = db.Model(&model).Where("id = ?", id).Delete(&model).Error; err == nil {
		toViewModel(&model, out)
	}
	return err
}
