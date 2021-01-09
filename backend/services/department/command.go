package department

import (
	"backend/models"
	"backend/repository"
	"backend/viewmodels"
	"gorm.io/gorm"
)

func Update(db *gorm.DB, id uint, out *viewmodels.DepartmentDto) error {
	var (
		err   error
		model models.Department
	)

	out.Id = int(id)

	toModel(db, out, &model)
	err = repository.Update(db, &model)
	return err
}

func Save(db *gorm.DB, out *viewmodels.DepartmentDto) error {
	var (
		err   error
		model models.Department
	)

	toModel(db, out, &model)
	if err = repository.Save(db, &model); err == nil {
		out.Id = int(model.ID)
	}
	return err
}

func Delete(db *gorm.DB, id uint, out *viewmodels.DepartmentDto) error {
	var (
		err   error
		model models.Department
	)

	model.ID = id

	if err = repository.Delete(db, &model); err == nil {
		toViewModel(&model, out)
	}
	return err
}
