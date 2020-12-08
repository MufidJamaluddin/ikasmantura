package viewmodels

import "backend/utils"

type AlbumDto struct {
	Id     uint   `query:"-" json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	UserId uint   `query:"userId,omitempty" json:"userId,omitempty" xml:"userId,omitempty" form:"userId,omitempty"`
	Title  string `query:"title,omitempty" json:"title,omitempty" xml:"title,omitempty" form:"title,omitempty"`
	utils.Created
	utils.Updated
}

func (p *AlbumDto) GetId() uint {
	return p.Id
}

type AlbumPhotoDto struct {
	AlbumId   uint   `query:"albumId" json:"albumId,omitempty" xml:"albumId,omitempty" form:"albumId,omitempty"`
	Title     string `query:"title" json:"title,omitempty" xml:"title,omitempty" form:"title,omitempty"`
	Id        uint   `query:"-" json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	Image     string `query:"-" json:"original,omitempty" xml:"original,omitempty" form:"original,omitempty"`
	Thumbnail string `query:"-" json:"thumbnail,omitempty" xml:"thumbnail,omitempty" form:"thumbnail,omitempty"`
	utils.Created
	utils.Updated
}

func (p *AlbumPhotoDto) GetId() uint {
	return p.Id
}

type AlbumPhotoParam struct {
	utils.GetParams
	AlbumPhotoDto
}

func (p *AlbumPhotoParam) GetModel() interface{} {
	return p
}
