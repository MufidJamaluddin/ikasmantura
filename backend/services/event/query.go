package event

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
	"backend/viewmodels"
	"database/sql"
	"gorm.io/gorm"
	"strings"
)

var eventSearchFields []string

func init() {
	eventSearchFields = []string{
		"organizer", "title",
	}
}

func searchFilter(tx *gorm.DB, search *viewmodels.EventParam, withLimit bool) {
	var title string

	search.Filter(tx, eventSearchFields, withLimit)

	if search.StartFrom != nil {
		tx.Where("start >= ?", search.StartFrom)
	}

	if search.EndTo != nil {
		tx.Where("end <= ?", search.EndTo)
	}

	title = strings.Trim(search.Title, " ")
	if title != "" {
		title = utils.ToLikeSQL(title)
		tx.Where("title LIKE ?", title)
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
		tx = db.Model(&model).Joins("INNER JOIN user_events ON user_id = ?", search.CurrentUserId)
	} else {
		tx = db.Model(&model)
	}

	searchFilter(tx, search, false)

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

	tx = db.Session(&gorm.Session{FullSaveAssociations: false})
	tx = tx.Model(&model)

	searchFilter(tx, search, true)

	if isCurrentUserSearch {
		if search.IsMyEvent {
			tx.Statement.Joins = nil
			tx = tx.Joins("INNER JOIN user_events ON user_id = ?", search.CurrentUserId)
		} else {
			tx.Statement.Preloads = nil
			tx = tx.Preload("Participants", "user_id = ?", search.CurrentUserId)
		}
	}

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

func FindById(db *gorm.DB, id string, out *viewmodels.EventDto) error {
	var (
		err     error
		err2    error
		model   models.Event
		creator models.User
		uid     utils.UUID
	)

	if uid, err = utils.FromBase64UUID(id); err != nil {
		return err
	}

	if err = db.Where("id = ?", uid.OrderedValue().Bytes()).First(&model).Error; err == nil {
		toViewModel(&model, out, false)
	}

	if err2 = repository.FindById(db, model.CreatedBy, &creator); err2 == nil {
		out.CreatedByName = creator.Name
	}

	return err
}
