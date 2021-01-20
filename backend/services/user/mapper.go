package user

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
	"crypto/sha1"
	"fmt"
	"hash"
)

func toModel(data *viewmodels.UserDto, out *models.User) {
	var (
		classrooms []models.UserClassroom
		hasher     hash.Hash
	)

	out.ID = uint(data.Id)
	out.Name = data.Name
	out.Username = data.Username
	out.Email = data.Email
	out.EmailValid = data.EmailValid
	out.Role = data.Role

	if data.Password != "" {
		hasher = sha1.New()
		hasher.Write(utils.ToBytes(data.Password))
		out.Password = fmt.Sprintf("%x", hasher.Sum(nil))
	}

	out.ForceYear = data.ForceYear
	out.Job = data.Job
	out.JobDesc = data.JobDesc
	out.Address.ID = uint(data.Address.ID)
	out.Address.Street = data.Address.Street
	out.Address.Suite = data.Address.Suite
	out.Address.City = data.Address.City
	out.Address.Zipcode = data.Address.Zipcode
	out.Address.State = data.Address.State

	for _, item := range data.Classrooms {
		classrooms = append(classrooms, models.UserClassroom{
			ClassroomId: uint(item),
			UserId:      uint(data.Id),
		})
	}

	out.Classrooms = classrooms

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func toViewModel(in *models.User, out *viewmodels.UserDto) {
	var classrooms []int

	out.Id = int(in.ID)
	out.Name = in.Name
	out.Username = in.Username
	out.Email = in.Email
	out.EmailValid = in.EmailValid
	out.Role = in.Role
	out.Password = in.Password
	out.ForceYear = in.ForceYear
	out.Job = in.Job
	out.JobDesc = in.JobDesc
	out.RefreshToken = utils.ToBase64UUID(in.RefreshToken)
	out.Address.ID = int(in.Address.ID)
	out.Address.Street = in.Address.Street
	out.Address.Suite = in.Address.Suite
	out.Address.City = in.Address.City
	out.Address.Zipcode = in.Address.Zipcode
	out.Address.State = in.Address.State

	for _, item := range in.Classrooms {
		classrooms = append(classrooms, int(item.ClassroomId))
	}

	out.Classrooms = classrooms

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}
