package impl

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
	"backend/viewmodels"
	"database/sql"
	"gorm.io/gorm"
)

var departmentSearchFields []string

func init() {
	departmentSearchFields = []string{
		"name",
	}
}

type DepartmentServiceImpl struct {
	DB *gorm.DB
}

func (p *DepartmentServiceImpl) toModel(data *viewmodels.DepartmentDto, out *models.Department) {
	out.ID = data.Id
	out.Name = data.Name
	out.Type = data.Type
	out.UserId = data.UserId

	p.DB.Find(&out.User, data.UserId)

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func (p *DepartmentServiceImpl) toData(in *models.Department, out *viewmodels.DepartmentDto) {
	out.Id = in.ID
	out.Name = in.Name
	out.UserId = in.UserId
	out.Type = in.Type
	out.UserFullname = in.User.Name

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}

func (p *DepartmentServiceImpl) GetTotal(search *viewmodels.DepartmentParam) (uint, error) {
	var (
		err   error
		model models.Department
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

func (p *DepartmentServiceImpl) Find(search *viewmodels.DepartmentParam, callback func(*viewmodels.DepartmentDto)) error {
	var (
		err   error
		model models.Department
		tx    *gorm.DB
		rows  *sql.Rows
	)

	tx = p.DB.Model(&model).Joins("User")
	search.Filter(tx, departmentSearchFields)

	if rows, err = tx.Rows(); err != nil {
		return err
	}

	for rows.Next() {
		err = tx.ScanRows(rows, &model)

		p.toData(&model, &search.DepartmentDto)
		callback(&search.DepartmentDto)
	}

	return err
}

func (p *DepartmentServiceImpl) FindById(id uint, out *viewmodels.DepartmentDto) error {
	var (
		err   error
		model models.Department
	)

	if err = repository.FindById(p.DB, id, &model); err == nil {
		p.toData(&model, out)
	}
	return err
}

func (p *DepartmentServiceImpl) Update(id uint, out *viewmodels.DepartmentDto) error {
	var (
		err   error
		model models.Department
	)

	out.Id = id

	p.toModel(out, &model)
	err = repository.Update(p.DB, &model)
	return err
}

func (p *DepartmentServiceImpl) Save(out *viewmodels.DepartmentDto) error {
	var (
		err   error
		model models.Department
	)

	p.toModel(out, &model)
	if err = repository.Save(p.DB, &model); err == nil {
		out.Id = model.ID
	}
	return err
}

func (p *DepartmentServiceImpl) Delete(id uint, out *viewmodels.DepartmentDto) error {
	var (
		err   error
		model models.Department
		hist  models.DepartmentHistory
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
