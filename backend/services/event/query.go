package event

import (
	"backend/models"
	"backend/repository"
	"backend/viewmodels"
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

var eventSearchFields []string

func init() {
	eventSearchFields = []string{
		"organizer", "title",
	}
}

func searchFilter(tx *gorm.DB, search *viewmodels.EventParam) {
	var title string

	search.Filter(tx, eventSearchFields)

	if search.StartFrom != nil {
		tx.Where("start >= ?", search.StartFrom)
	}

	if search.EndTo != nil {
		tx.Where("end <= ?", search.EndTo)
	}

	title = strings.Trim(search.Title, " ")
	if title != "" {
		tx.Where("title LIKE ?", fmt.Sprintf("%s%", title))
	}
}

func GetTotal(db *gorm.DB, search *viewmodels.EventParam) (uint, error) {
	var (
		err   error
		total int64
		model models.Event
		tx    *gorm.DB
	)

	if search.IsMyEvent && search.CurrentUserId != 0 {
		tx = db.Model(&model).Joins("JOIN user_events ON userId = ?", search.CurrentUserId)
	} else {
		tx = db.Model(&model)
	}

	searchFilter(tx, search)

	tx.Count(&total)

	return uint(total), err
}

func Find(db *gorm.DB, search *viewmodels.EventParam, callback func(*viewmodels.EventDto)) error {
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
			tx = db.Model(&model).
				Joins("JOIN user_events ON userId = ?", search.CurrentUserId)
		} else {
			tx = db.Model(&model).
				Preload("user_events", "user_id = ?", search.CurrentUserId)
		}
	} else {
		tx = db.Model(&model)
	}

	searchFilter(tx, search)

	if rows, err = tx.Rows(); err != nil {
		return err
	}

	for rows.Next() {
		err = tx.ScanRows(rows, &model)

		toViewModel(&model, &search.EventDto, isCurrentUserSearch)
		callback(&search.EventDto)
	}

	return err
}

func FindById(db *gorm.DB, id uint, out *viewmodels.EventDto) error {
	var (
		err     error
		err2    error
		model   models.Event
		creator models.User
	)

	if err = repository.FindById(db, id, &model); err == nil {
		toViewModel(&model, out, false)
	}

	if err2 = repository.FindById(db, model.CreatedBy, &creator); err2 == nil {
		out.CreatedByName = creator.Name
	}

	return err
}
