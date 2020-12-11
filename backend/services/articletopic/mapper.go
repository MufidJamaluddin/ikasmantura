package articletopic

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
)

func toModel(data *viewmodels.ArticleTopicDto, out *models.ArticleTopic) {
	out.ID = data.Id
	out.Name = data.Name
	out.Icon = data.Icon
	out.Description = data.Description

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func toViewModel(in *models.ArticleTopic, out *viewmodels.ArticleTopicDto) {
	out.Id = in.ID
	out.Name = in.Name
	out.Icon = in.Icon
	out.Description = in.Description

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}
