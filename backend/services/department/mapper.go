package department

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
	"gorm.io/gorm"
)

func toModel(db *gorm.DB, data *viewmodels.DepartmentDto, out *models.Department) {
	out.ID = uint(data.Id)
	out.Name = data.Name
	out.Type = uint8(data.Type)
	out.UserId = uint(data.UserId)

	db.Find(&out.User, data.UserId)

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func toViewModel(in *models.Department, out *viewmodels.DepartmentDto) {
	out.Id = int(in.ID)
	out.Name = in.Name
	out.UserId = int(in.UserId)
	out.Type = int(in.Type)
	out.UserFullname = in.User.Name

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}
