package classroom

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
)

func toModel(data *viewmodels.ClassroomDto, out *models.Classroom) {
	out.ID = data.Id
	out.Level = data.Level
	out.Major = data.Major
	out.Seq = data.Seq

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func toViewModel(in *models.Classroom, out *viewmodels.ClassroomDto) {
	out.Id = in.ID
	out.Level = in.Level
	out.Major = in.Major
	out.Seq = in.Seq

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}
