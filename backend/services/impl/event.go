package impl

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
	"backend/viewmodels"
	"database/sql"
	"gorm.io/gorm"
)

var eventSearchFields []string

func init() {
	eventSearchFields = []string{
		"organizer", "title",
	}
}

type EventServiceImpl struct {
	DB *gorm.DB
}

func (p *EventServiceImpl) toModel(data *viewmodels.EventDto, out *models.Event) {
	out.ID = data.Id
	out.Title = data.Title
	out.Description = data.Description
	out.Start = data.Start
	out.End = data.End
	out.Image = data.Image

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func (p *EventServiceImpl) toData(in *models.Event, out *viewmodels.EventDto, isCurrentUserSearch bool) {
	out.Id = in.ID
	out.Title = in.Title
	out.Description = in.Description
	out.Start = in.Start
	out.End = in.End
	out.Image = in.Image

	if isCurrentUserSearch {
		total := len(in.Participants)
		out.IsMyEvent = total > 0
		if out.IsMyEvent {
			item := in.Participants[total-1]
			out.CurrentUserRegisData.ID = item.ID
			out.CurrentUserRegisData.UserId = item.UserId
			out.CurrentUserRegisData.EventId = item.EventId
			utils.FillCreated(item, out.CurrentUserRegisData)
		}
	}

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}

func (p *EventServiceImpl) searchFilter(tx *gorm.DB, search *viewmodels.EventParam) {
	search.Filter(tx, eventSearchFields)

	if search.StartFrom != nil {
		tx.Where("start >= ?", search.StartFrom)
	}

	if search.EndTo != nil {
		tx.Where("end <= ?", search.EndTo)
	}
}

func (p *EventServiceImpl) GetTotal(search *viewmodels.EventParam) (uint, error) {
	var (
		err   error
		total int64
		model models.Event
		tx    *gorm.DB
	)

	if search.IsMyEvent && search.CurrentUserId != 0 {
		tx = p.DB.Model(&model).Joins("JOIN user_events ON userId = ?", search.CurrentUserId)
	} else {
		tx = p.DB.Model(&model)
	}
	p.searchFilter(tx, search)

	tx.Count(&total)

	return uint(total), err
}

func (p *EventServiceImpl) Find(search *viewmodels.EventParam, callback func(*viewmodels.EventDto)) error {
	var (
		err                 error
		model               models.Event
		tx                  *gorm.DB
		rows                *sql.Rows
		isCurrentUserSearch bool
	)

	isCurrentUserSearch = search.CurrentUserId != 0

	if isCurrentUserSearch {
		if search.IsMyEvent {
			tx = p.DB.Model(&model).
				Joins("JOIN user_events ON userId = ?", search.CurrentUserId)
		} else {
			tx = p.DB.Model(&model).
				Preload("user_events", "user_id = ?", search.CurrentUserId)
		}
	} else {
		tx = p.DB.Model(&model)
	}

	p.searchFilter(tx, search)

	if rows, err = tx.Rows(); err != nil {
		return err
	}

	for rows.Next() {
		err = tx.ScanRows(rows, &model)

		p.toData(&model, &search.EventDto, isCurrentUserSearch)
		callback(&search.EventDto)
	}

	return err
}

func (p *EventServiceImpl) FindById(id uint, out *viewmodels.EventDto) error {
	var (
		err     error
		err2    error
		model   models.Event
		creator models.User
	)

	if err = repository.FindById(p.DB, id, &model); err == nil {
		p.toData(&model, out, false)
	}

	if err2 = repository.FindById(p.DB, model.CreatedBy, &creator); err2 == nil {
		out.CreatedByName = creator.Name
	}

	return err
}

func (p *EventServiceImpl) Update(id uint, out *viewmodels.EventDto) error {
	var (
		err   error
		model models.Event
	)

	out.Id = id

	p.toModel(out, &model)
	err = repository.Update(p.DB, &model)
	return err
}

func (p *EventServiceImpl) Save(out *viewmodels.EventDto) error {
	var (
		err   error
		model models.Event
	)

	p.toModel(out, &model)
	if err = repository.Save(p.DB, &model); err == nil {
		out.Id = model.ID
	}
	return err
}

func (p *EventServiceImpl) Delete(id uint, out *viewmodels.EventDto) error {
	var (
		err   error
		model models.Event
		hist  models.EventHistory
	)

	model.ID = id

	if err = repository.Delete(p.DB, &model); err == nil {
		p.toData(&model, out, false)

		p.DB.Model(&hist).
			Where("ID = ?", id).
			Order("version desc").
			Last(&hist)

		if hist.Action == "delete" {
			hist.UpdatedBy = model.UpdatedBy
			p.DB.Save(&hist)
		}
	}
	return err
}

func (p *EventServiceImpl) RegisterEvent(eventId uint, userId uint) error {
	var (
		err        error
		eventModel models.Event
		userModel  models.User
		userEvent  models.UserEvent
	)

	p.DB.First(&userModel, userId)
	p.DB.First(&eventModel, eventId)

	if eventModel.ID != 0 && userModel.ID != 0 {

		userEvent.UserId = userId
		userEvent.EventId = eventId

		err = p.DB.Create(&userEvent).Error
	}

	return err
}

func (p *EventServiceImpl) GetUserEvent(data *viewmodels.UserEventDetailDto) error {
	var (
		err       error
		userEvent models.UserEvent
	)

	p.DB.Where("user_id = ? AND event_id = ?", data.UserId, data.EventId).
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
