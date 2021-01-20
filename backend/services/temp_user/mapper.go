package temp_user

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
	"crypto/sha1"
	"fmt"
	"hash"
	"log"
	"strings"
	"unicode/utf8"
)

func toPermanentModel(data *models.TempUser, out *models.User) {
	var classrooms []models.UserClassroom

	out.Name = data.Name
	out.Username = data.Username
	out.Email = data.Email
	out.EmailValid = data.EmailValid
	out.Role = "member"
	out.Password = data.Password
	out.ForceYear = data.ForceYear
	out.Job = data.Job
	out.JobDesc = data.JobDesc
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
	var (
		classrooms []models.TempUserClassroom
		hasher     hash.Hash
		err error
	)

	out.ID = uint(data.Id)
	out.Name = data.Name
	out.Username = data.Username
	out.Email = data.Email
	out.EmailValid = data.EmailValid

	if out.ConfirmEmailToken, err = utils.FromBase64UUID(data.ConfirmEmailToken); err != nil {
		log.Println(err.Error())
	}

	out.ForceYear = data.ForceYear

	if data.Password != "" {
		hasher = sha1.New()
		hasher.Write(utils.ToBytes(data.Password))
		out.Password = fmt.Sprintf("%x", hasher.Sum(nil))
	}

	data.Job = strings.Trim(data.Job, " ")
	if utf8.RuneCountInString(data.Job) > 35 {
		data.Job = data.Job[:35]
	}
	out.Job = data.Job

	data.JobDesc = strings.Trim(data.JobDesc, " ")
	if utf8.RuneCountInString(data.JobDesc) > 85 {
		data.JobDesc = data.JobDesc[:85]
	}
	out.JobDesc = data.JobDesc

	data.Phone = strings.Trim(data.Phone, " ")
	if utf8.RuneCountInString(data.Phone) > 13 {
		data.Phone = data.Phone[:13]
	}
	out.Phone = data.Phone

	out.Address.ID = uint(data.Address.ID)
	out.Address.Street = data.Address.Street

	data.Address.Suite = strings.Trim(data.Address.Suite, " ")
	if utf8.RuneCountInString(data.Address.Suite) > 35 {
		data.Address.Suite = data.Address.Suite[:35]
	}
	out.Address.Suite = data.Address.Suite

	data.Address.City = strings.Trim(data.Address.City, " ")
	if utf8.RuneCountInString(data.Address.City) > 35 {
		out.Address.City = data.Address.City[:35]
	} else {
		out.Address.City = data.Address.City
	}

	data.Address.Zipcode = strings.Trim(data.Address.Zipcode, " ")
	if utf8.RuneCountInString(data.Address.Zipcode) > 4 {
		out.Address.Zipcode = data.Address.Zipcode[:4]
	} else {
		out.Address.Zipcode = data.Address.Zipcode
	}

	data.Address.State = strings.Trim(data.Address.State, " ")
	if utf8.RuneCountInString(data.Address.State) > 35 {
		out.Address.State = data.Address.State[:35]
	} else {
		out.Address.State = data.Address.State
	}

	for _, item := range data.Classrooms {
		classrooms = append(classrooms, models.TempUserClassroom{
			ClassroomId: uint(item),
			UserId:      uint(data.Id),
		})
	}

	out.Classrooms = classrooms

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func toViewModel(in *models.TempUser, out *viewmodels.UserDto) {
	var classrooms []int

	out.Id = int(in.ID)
	out.Name = in.Name
	out.Username = in.Username
	out.Email = in.Email
	out.EmailValid = in.EmailValid
	out.ConfirmEmailToken = utils.ToBase64UUID(in.ConfirmEmailToken)
	out.Phone = in.Phone
	out.Role = ""
	out.Password = in.Password
	out.RefreshToken = utils.ToBase64UUID(in.RefreshToken)
	out.ForceYear = in.ForceYear
	out.Job = in.Job
	out.JobDesc = in.JobDesc
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
