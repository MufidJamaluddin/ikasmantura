package event

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
	"log"
)

func toModel(data *viewmodels.EventDto, out *models.Event) {
	var (
		uid utils.UUID
		err error
	)

	if uid, err = utils.FromBase64UUID(data.Id); err != nil {
		log.Println(err.Error())
	}

	out.ID = uid
	out.Title = data.Title
	out.Description = data.Description
	out.Start = data.Start
	out.End = data.End
	out.Image = data.Image
	out.Thumbnail = data.Thumbnail

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func toViewModel(in *models.Event, out *viewmodels.EventDto, isCurrentUserSearch bool) {
	out.Id = utils.ToBase64UUID(in.ID)
	out.Title = in.Title
	out.Description = in.Description
	out.Start = in.Start
	out.End = in.End
	out.Image = in.Image
	out.Thumbnail = in.Thumbnail

	if isCurrentUserSearch {
		total := len(in.Participants)
		out.IsMyEvent = total > 0
		if out.IsMyEvent {
			item := in.Participants[total-1]
			out.CurrentUserRegisData.ID = int(item.ID)
			out.CurrentUserRegisData.UserId = int(item.UserId)
			out.CurrentUserRegisData.EventId = utils.ToBase64UUID(item.EventId)
			utils.FillCreated(item, out.CurrentUserRegisData)
		}
	}

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}
