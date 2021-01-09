package temp_user

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
	"strings"
)

func toPermanentModel(data *models.TempUser, out *models.User) {
	var classrooms []models.UserClassroom

	out.ID = data.ID
	out.Name = data.Name
	out.Username = data.Username
	out.Role = "member"
	out.Password = data.Password
	out.ForceYear = data.ForceYear
	out.RefreshToken = data.RefreshToken
	out.Address.ID = data.Address.ID
	out.Address.Street = data.Address.Street
	out.Address.Suite = data.Address.Suite
	out.Address.City = data.Address.City
	out.Address.Zipcode = data.Address.Zipcode

	for _, item := range data.Classrooms {
		classrooms = append(classrooms, models.UserClassroom{
			ClassroomId: item.ClassroomId,
			UserId:      item.UserId,
		})
	}

	out.Classrooms = classrooms

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func toTempModel(data *viewmodels.UserDto, out *models.TempUser) {
	var classrooms []models.TempUserClassroom

	out.ID = uint(data.Id)
	out.Name = data.Name
	out.Username = data.Username
	out.Password = data.Password
	out.ForceYear = data.ForceYear

	data.Phone = strings.Trim(data.Phone, " ")
	if len(data.Phone) > 13 {
		data.Phone = data.Phone[0:13]
	}

	out.Phone = data.Phone
	out.Address.ID = uint(data.Address.ID)
	out.Address.Street = data.Address.Street

	data.Address.Suite = strings.Trim(data.Address.Suite, " ")
	if len(data.Address.Suite) > 35 {
		out.Address.Suite = data.Address.Suite[0:35]
	} else {
		out.Address.Suite = data.Address.Suite
	}

	data.Address.City = strings.Trim(data.Address.City, " ")
	if len(data.Address.City) > 35 {
		out.Address.City = data.Address.City[0:35]
	} else {
		out.Address.City = data.Address.City
	}

	data.Address.Zipcode = strings.Trim(data.Address.Zipcode, " ")
	if len(data.Address.Zipcode) > 4 {
		out.Address.Zipcode = data.Address.Zipcode[0:4]
	} else {
		out.Address.Zipcode = data.Address.Zipcode
	}

	data.Address.State = strings.Trim(data.Address.State, " ")
	if len(data.Address.State) > 35 {
		out.Address.State = data.Address.State[0:35]
	} else {
		out.Address.State = data.Address.State
	}

	for _, item := range data.Classrooms {
		classrooms = append(classrooms, models.TempUserClassroom{
			ClassroomId: uint(item.Id),
			UserId:      uint(data.Id),
		})
	}

	out.Classrooms = classrooms

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func toViewModel(in *models.TempUser, out *viewmodels.UserDto) {
	var classrooms []viewmodels.ClassroomDto

	out.Id = int(in.ID)
	out.Name = in.Name
	out.Username = in.Username
	out.Email = in.Email
	out.Phone = in.Phone
	out.Role = ""
	out.Password = in.Password
	out.RefreshToken = utils.ToBase64UUID(in.RefreshToken)
	out.ForceYear = in.ForceYear
	out.Address.ID = int(in.Address.ID)
	out.Address.Street = in.Address.Street
	out.Address.Suite = in.Address.Suite
	out.Address.City = in.Address.City
	out.Address.Zipcode = in.Address.Zipcode
	out.Address.State = in.Address.State

	for _, item := range in.Classrooms {
		classrooms = append(classrooms, viewmodels.ClassroomDto{
			Id:    int(item.ClassroomId),
			Major: item.Classroom.Major,
			Level: item.Classroom.Level,
			Seq:   int(item.Classroom.Seq),
		})
	}

	out.Classrooms = classrooms

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}
