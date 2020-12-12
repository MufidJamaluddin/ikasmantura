package about

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
)

func toModel(data *viewmodels.AboutDto, out *models.About) {
	out.ID = data.Id
	out.Title = data.Title
	out.Description = data.Description
	out.Mission = data.Mission
	out.Vision = data.Vision
	out.Facebook = data.Facebook
	out.Twitter = data.Twitter
	out.Instagram = data.Instagram

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func toViewModel(in *models.About, out *viewmodels.AboutDto) {
	out.Id = in.ID
	out.Description = in.Description
	out.Mission = in.Mission
	out.Vision = in.Vision

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}
