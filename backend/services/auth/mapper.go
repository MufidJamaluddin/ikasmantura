package auth

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
)

func toViewModel(in *models.User, out *viewmodels.UserDto) {
	out.Id = int(in.ID)
	out.Name = in.Name
	out.Username = in.Username
	out.Email = in.Email
	out.Password = in.Password
	out.Role = in.Role
	out.RefreshToken = utils.ToBase64UUID(in.RefreshToken)
	out.Address.ID = int(in.Address.ID)
	out.Address.Street = in.Address.Street
	out.Address.Suite = in.Address.Suite
	out.Address.City = in.Address.City
	out.Address.Zipcode = in.Address.Zipcode

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}

func tempUserToUserModel(data *models.TempUser, out *models.User) {
	out.ID = 0
	out.Name = data.Name
	out.Username = data.Username
	out.Role = ""
	out.Password = data.Password
	out.ForceYear = data.ForceYear
	out.Job = data.Job
	out.JobDesc = data.JobDesc
	out.RefreshToken = data.RefreshToken

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}
