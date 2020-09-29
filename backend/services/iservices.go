package services

import (
	"backend/dto"
)

type AboutService interface {
	FindById(id uint, out *dto.AboutDto) error
	Update(id uint, data *dto.AboutDto) error
}

type AlbumService interface {
	GetTotal(params *dto.AlbumParam) (uint, error)
	Find(params *dto.AlbumParam, callback func(albumDto *dto.AlbumDto)) error
	FindById(id uint, out *dto.AlbumDto) error
	Save(data *dto.AlbumDto) error
	Update(id uint, data *dto.AlbumDto) error
	Delete(id uint, out *dto.AlbumDto) error
}

type AlbumPhotoService interface {
	GetTotal(model *dto.AlbumPhotoParam) (uint, error)
	Find(model *dto.AlbumPhotoParam, callback func(photoDto *dto.AlbumPhotoDto)) error
	FindById(id uint, out *dto.AlbumPhotoDto) error
	Save(data *dto.AlbumPhotoDto) error
	Update(id uint, data *dto.AlbumPhotoDto) error
	Delete(id uint, out *dto.AlbumPhotoDto) error
}

type ArticleService interface {
	GetTotal(model *dto.ArticleParam) (uint, error)
	Find(model *dto.ArticleParam, callback func(articleDto *dto.ArticleDto)) error
	FindById(id uint, out *dto.ArticleDto) error
	Save(data *dto.ArticleDto) error
	Update(id uint, data *dto.ArticleDto) error
	Delete(id uint, out *dto.ArticleDto) error
}

type ArticleTopicService interface {
	GetTotal(model *dto.ArticleTopicParam) (uint, error)
	Find(model *dto.ArticleTopicParam, callback func(topicDto *dto.ArticleTopicDto)) error
	FindById(id uint, out *dto.ArticleTopicDto) error
	Save(data *dto.ArticleTopicDto) error
	Update(id uint, data *dto.ArticleTopicDto) error
	Delete(id uint, out *dto.ArticleTopicDto) error
}

type AuthService interface {
	Login(data *dto.LoginDto) error
}

type DepartmentService interface {
	GetTotal(search *dto.DepartmentParam) (uint, error)
	Find(search *dto.DepartmentParam, callback func(departmentDto *dto.DepartmentDto)) error
	FindById(id uint, out *dto.DepartmentDto) error
	Save(data *dto.DepartmentDto) error
	Update(id uint, data *dto.DepartmentDto) error
	Delete(id uint, out *dto.DepartmentDto) error
}

type EventService interface {
	GetTotal(model *dto.EventParam) (uint, error)
	Find(model *dto.EventParam, callback func(eventDto *dto.EventDto)) error
	FindById(id uint, out *dto.EventDto) error
	Save(model *dto.EventDto) error
	Update(id uint, model *dto.EventDto) error
	Delete(id uint, out *dto.EventDto) error

	RegisterEvent(eventId uint, userId uint) error
	GetUserEvent(data *dto.UserEventDetailDto) error
}

type UserService interface {
	GetTotal(model *dto.UserParam) (uint, error)
	Find(model *dto.UserParam, callback func(userDto *dto.UserDto)) error
	FindById(id uint, out *dto.UserDto) error
	Save(model *dto.UserDto) error
	Update(id uint, model *dto.UserDto) error
	Delete(id uint, out *dto.UserDto) error
}
