package user

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
)

func toModel(data *viewmodels.UserDto, out *models.User) {
	out.ID = data.Id
	out.Name = data.Name
	out.Username = data.Username
	out.IsAdmin = data.IsAdmin
	out.Password = data.Password
	out.ForceYear = data.ForceYear
	out.Address.ID = data.Address.ID
	out.Address.Street = data.Address.Street
	out.Address.Suite = data.Address.Suite
	out.Address.City = data.Address.City
	out.Address.Zipcode = data.Address.Zipcode

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func toViewModel(in *models.User, out *viewmodels.UserDto) {
	out.Id = in.ID
	out.Name = in.Name
	out.Username = in.Username
	out.IsAdmin = in.IsAdmin
	out.Password = in.Password
	out.ForceYear = in.ForceYear
	out.Address.ID = in.Address.ID
	out.Address.Street = in.Address.Street
	out.Address.Suite = in.Address.Suite
	out.Address.City = in.Address.City
	out.Address.Zipcode = in.Address.Zipcode

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}
