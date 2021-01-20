package albumphoto

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
)

func toModel(data *viewmodels.AlbumPhotoDto, out *models.AlbumPhoto) {
	out.ID = uint(data.Id)
	out.AlbumId = uint(data.AlbumId)
	out.Title = data.Title
	out.Image = data.Image
	out.Thumbnail = data.Thumbnail

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func toViewModel(in *models.AlbumPhoto, out *viewmodels.AlbumPhotoDto) {
	out.Id = int(in.ID)
	out.AlbumId = int(in.AlbumId)
	out.Title = in.Title
	out.Image = in.Image
	out.Thumbnail = in.Thumbnail

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}
