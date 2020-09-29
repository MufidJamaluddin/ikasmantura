package impl

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
	"backend/viewmodels"
	"database/sql"
	"gorm.io/gorm"
)

var userSearchFields []string

func init() {
	userSearchFields = []string{
		"name",
	}
}

type UserServiceImpl struct {
	DB *gorm.DB
}

func (p *UserServiceImpl) toModel(data *viewmodels.UserDto, out *models.User) {
	out.ID = data.Id
	out.Name = data.Name
	out.Username = data.Username
	out.IsAdmin = data.IsAdmin
	out.Password = data.Password
	out.Address.ID = data.Address.ID
	out.Address.Street = data.Address.Street
	out.Address.Suite = data.Address.Suite
	out.Address.City = data.Address.City
	out.Address.Zipcode = data.Address.Zipcode

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func (p *UserServiceImpl) toData(in *models.User, out *viewmodels.UserDto) {
	out.Id = in.ID
	out.Name = in.Name
	out.Username = in.Username
	out.IsAdmin = in.IsAdmin
	out.Password = in.Password
	out.Address.ID = in.Address.ID
	out.Address.Street = in.Address.Street
	out.Address.Suite = in.Address.Suite
	out.Address.City = in.Address.City
	out.Address.Zipcode = in.Address.Zipcode

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}

func (p *UserServiceImpl) GetTotal(search *viewmodels.UserParam) (uint, error) {
	var (
		err   error
		model models.User
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

func (p *UserServiceImpl) Find(search *viewmodels.UserParam, callback func(*viewmodels.UserDto)) error {
	var (
		err   error
		model models.User
		tx    *gorm.DB
		rows  *sql.Rows
	)

	tx = p.DB.Model(&model)
	search.Filter(tx, userSearchFields)

	if rows, err = tx.Rows(); err != nil {
		return err
	}

	for rows.Next() {
		err = tx.ScanRows(rows, &model)

		p.toData(&model, &search.UserDto)
		callback(&search.UserDto)
	}

	return err
}

func (p *UserServiceImpl) FindById(id uint, out *viewmodels.UserDto) error {
	var (
		err   error
		model models.User
	)

	if err = repository.FindById(p.DB, id, &model); err == nil {
		p.toData(&model, out)
	}
	return err
}

func (p *UserServiceImpl) Update(id uint, out *viewmodels.UserDto) error {
	var (
		err   error
		model models.User
	)

	out.Id = id

	p.toModel(out, &model)
	err = repository.Update(p.DB, &model)
	return err
}

func (p *UserServiceImpl) Save(out *viewmodels.UserDto) error {
	var (
		err   error
		model models.User
	)

	p.toModel(out, &model)
	if err = repository.Save(p.DB, &model); err == nil {
		out.Id = model.ID
	}
	return err
}

func (p *UserServiceImpl) Delete(id uint, out *viewmodels.UserDto) error {
	var (
		err   error
		model models.User
		hist  models.UserHistory
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
