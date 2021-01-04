package event

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
	"backend/viewmodels"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func Update(db *gorm.DB, id string, out *viewmodels.EventDto) error {
	var (
		err   error
		model models.Event
	)

	out.Id = id

	toModel(out, &model)
	err = repository.Update(db, &model)
	return err
}

func Save(db *gorm.DB, out *viewmodels.EventDto) error {
	var (
		err   error
		model models.Event
	)

	toModel(out, &model)
	if err = repository.Save(db, &model); err == nil {
		out.Id = model.ID.Guid().String()
	}
	return err
}

func Delete(db *gorm.DB, id string, out *viewmodels.EventDto) error {
	var (
		err   error
		model models.Event
		uid   utils.UUID
	)

	if uid, err = utils.FromBase64UUID(id); err != nil {
		return err
	}

	model.ID = uid

	if err = repository.Delete(db, &model); err == nil {
		toViewModel(&model, out, false)
	}
	return err
}

func RegisterEvent(db *gorm.DB, eventId uint, userId uint) error {
	var (
		err        error
		eventModel models.Event
		userModel  models.User
		userEvent  models.UserEvent
	)

	db.First(&userModel, userId)
	db.First(&eventModel, eventId)

	if eventModel.ID.Guid() != uuid.Nil && userModel.ID != 0 {

		userEvent.UserId = userId
		userEvent.EventId = eventId

		err = db.Model(&userEvent).Create(&userEvent).Error
	}

	return err
}

func GetUserEvent(db *gorm.DB, data *viewmodels.UserEventDetailDto) error {
	var (
		err       error
		userEvent models.UserEvent
	)

	db.Where("user_events.user_id = ? AND user_events.event_id = ?", data.UserId, data.EventId).
		Joins("event").
		Joins("user").
		First(&userEvent)

	data.Description = userEvent.Event.Description
	data.EventName = userEvent.Event.Title
	data.Organizer = userEvent.Event.Organizer
	data.StartStr = userEvent.Event.Start.Format("Monday, 2 January 2006, 15:04")
	data.EndStr = userEvent.Event.End.Format("Monday, 2 January 2006, 15:04")
	data.UserFullName = userEvent.User.Name
	data.UserEmail = userEvent.User.Email

	return err
}
