package user

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
)

func toModel(data *viewmodels.UserDto, out *models.User) {
	var classrooms []models.UserClassroom

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
	out.Address.State = data.Address.State

	for _, item := range data.Classrooms {
		classrooms = append(classrooms, models.UserClassroom{
			ClassroomId: item.Id,
			UserId:      data.Id,
		})
	}

	out.Classrooms = classrooms

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func toViewModel(in *models.User, out *viewmodels.UserDto) {
	var classrooms []viewmodels.ClassroomDto

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
	out.Address.State = in.Address.State

	for _, item := range in.Classrooms {
		classrooms = append(classrooms, viewmodels.ClassroomDto{
			Id:    item.ClassroomId,
			Major: item.Classroom.Major,
			Level: item.Classroom.Level,
			Seq:   item.Classroom.Seq,
		})
	}

	out.Classrooms = classrooms

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}
