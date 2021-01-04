package article

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
	uuid "github.com/satori/go.uuid"
	"log"
)

func toModel(data *viewmodels.ArticleDto, out *models.Article) {
	var (
		uid uuid.UUID
		err error
	)

	if uid, err = uuid.FromString(data.Id); err != nil {
		log.Println(err.Error())
	}

	out.ID = models.UUID(uid)
	out.Title = data.Title
	out.Body = data.Body
	out.Image = data.Image
	out.Thumbnail = data.Thumbnail
	out.CreatedBy = data.UserId
	out.UpdatedBy = data.UserId
	out.ArticleTopicId = data.TopicId

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func toViewModel(in *models.Article, out *viewmodels.ArticleDto) {
	out.Id = in.ID.Guid().String()
	out.Title = in.Title
	out.Body = in.Body
	out.Image = in.Image
	out.Thumbnail = in.Thumbnail
	out.UserId = in.CreatedBy
	out.CreatedAt = in.CreatedAt
	out.TopicId = in.ArticleTopicId

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}
