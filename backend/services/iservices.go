package services

import (
	"backend/viewmodels"
)

type AboutService interface {
	FindById(id uint, out *viewmodels.AboutDto) error
	Update(id uint, data *viewmodels.AboutDto) error
}

type AlbumService interface {
	GetTotal(params *viewmodels.AlbumParam) (uint, error)
	Find(params *viewmodels.AlbumParam, callback func(albumDto *viewmodels.AlbumDto)) error
	FindById(id uint, out *viewmodels.AlbumDto) error
	Save(data *viewmodels.AlbumDto) error
	Update(id uint, data *viewmodels.AlbumDto) error
	Delete(id uint, out *viewmodels.AlbumDto) error
}

type AlbumPhotoService interface {
	GetTotal(model *viewmodels.AlbumPhotoParam) (uint, error)
	Find(model *viewmodels.AlbumPhotoParam, callback func(photoDto *viewmodels.AlbumPhotoDto)) error
	FindById(id uint, out *viewmodels.AlbumPhotoDto) error
	Save(data *viewmodels.AlbumPhotoDto) error
	Update(id uint, data *viewmodels.AlbumPhotoDto) error
	Delete(id uint, out *viewmodels.AlbumPhotoDto) error
}

type ArticleService interface {
	GetTotal(model *viewmodels.ArticleParam) (uint, error)
	Find(model *viewmodels.ArticleParam, callback func(articleDto *viewmodels.ArticleDto)) error
	FindById(id uint, out *viewmodels.ArticleDto) error
	Save(data *viewmodels.ArticleDto) error
	Update(id uint, data *viewmodels.ArticleDto) error
	Delete(id uint, out *viewmodels.ArticleDto) error
}

type ArticleTopicService interface {
	GetTotal(model *viewmodels.ArticleTopicParam) (uint, error)
	Find(model *viewmodels.ArticleTopicParam, callback func(topicDto *viewmodels.ArticleTopicDto)) error
	FindById(id uint, out *viewmodels.ArticleTopicDto) error
	Save(data *viewmodels.ArticleTopicDto) error
	Update(id uint, data *viewmodels.ArticleTopicDto) error
	Delete(id uint, out *viewmodels.ArticleTopicDto) error
}

type AuthService interface {
	Login(data *viewmodels.LoginDto) error
}

type DepartmentService interface {
	GetTotal(search *viewmodels.DepartmentParam) (uint, error)
	Find(search *viewmodels.DepartmentParam, callback func(departmentDto *viewmodels.DepartmentDto)) error
	FindById(id uint, out *viewmodels.DepartmentDto) error
	Save(data *viewmodels.DepartmentDto) error
	Update(id uint, data *viewmodels.DepartmentDto) error
	Delete(id uint, out *viewmodels.DepartmentDto) error
}

type EventService interface {
	GetTotal(model *viewmodels.EventParam) (uint, error)
	Find(model *viewmodels.EventParam, callback func(eventDto *viewmodels.EventDto)) error
	FindById(id uint, out *viewmodels.EventDto) error
	Save(model *viewmodels.EventDto) error
	Update(id uint, model *viewmodels.EventDto) error
	Delete(id uint, out *viewmodels.EventDto) error

	RegisterEvent(eventId uint, userId uint) error
	GetUserEvent(data *viewmodels.UserEventDetailDto) error
}

type UserService interface {
	GetTotal(model *viewmodels.UserParam) (uint, error)
	Find(model *viewmodels.UserParam, callback func(userDto *viewmodels.UserDto)) error
	FindById(id uint, out *viewmodels.UserDto) error
	Save(model *viewmodels.UserDto) error
	Update(id uint, model *viewmodels.UserDto) error
	Delete(id uint, out *viewmodels.UserDto) error
}
