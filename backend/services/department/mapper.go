package department

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
	"gorm.io/gorm"
)

func toModel(db *gorm.DB, data *viewmodels.DepartmentDto, out *models.Department) {
	out.ID = data.Id
	out.Name = data.Name
	out.Type = data.Type
	out.UserId = data.UserId

	db.Find(&out.User, data.UserId)

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func toViewModel(in *models.Department, out *viewmodels.DepartmentDto) {
	out.Id = in.ID
	out.Name = in.Name
	out.UserId = in.UserId
	out.Type = in.Type
	out.UserFullname = in.User.Name

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}
