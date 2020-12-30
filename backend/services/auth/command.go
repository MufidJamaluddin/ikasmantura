package auth

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
)

func toViewModel(in *models.User, out *viewmodels.UserDto) {
	out.Id = in.ID
	out.Name = in.Name
	out.Username = in.Username
	out.Email = in.Email
	out.Password = in.Password
	out.Role = in.Role
	out.Address.ID = in.Address.ID
	out.Address.Street = in.Address.Street
	out.Address.Suite = in.Address.Suite
	out.Address.City = in.Address.City
	out.Address.Zipcode = in.Address.Zipcode

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}

func toViewModelFromTemp(in *models.TempUser, out *viewmodels.UserDto) {
	out.Id = in.ID
	out.Name = in.Name
	out.Username = in.Username
	out.Email = in.Email
	out.Password = in.Password
	out.Role = ""
	out.Address.ID = in.Address.ID
	out.Address.Street = in.Address.Street
	out.Address.Suite = in.Address.Suite
	out.Address.City = in.Address.City
	out.Address.Zipcode = in.Address.Zipcode

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}
