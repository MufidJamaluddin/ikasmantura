package article

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
	"log"
)

func toModel(data *viewmodels.ArticleDto, out *models.Article) {
	var (
		uid utils.UUID
		err error
	)

	if uid, err = utils.FromBase64UUID(data.Id); err != nil {
		log.Println(err.Error())
	}

	out.ID = uid
	out.Title = data.Title
	out.Body = data.Body
	out.Image = data.Image
	out.Thumbnail = data.Thumbnail
	out.CreatedBy = uint(data.UserId)
	out.UpdatedBy = uint(data.UserId)
	out.ArticleTopicId = uint(data.TopicId)

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func toViewModel(in *models.Article, out *viewmodels.ArticleDto) {
	out.Id = utils.ToBase64UUID(in.ID)
	out.Title = in.Title
	out.Body = in.Body
	out.Image = in.Image
	out.Thumbnail = in.Thumbnail
	out.UserId = int(in.CreatedBy)
	out.CreatedAt = in.CreatedAt
	out.TopicId = int(in.ArticleTopicId)

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}
