package albumphoto

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
)

func toModel(data *viewmodels.AlbumPhotoDto, out *models.AlbumPhoto) {
	out.ID = data.Id
	out.Title = data.Title
	out.Image = data.Image
	out.Thumbnail = data.Thumbnail

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func toViewModel(in *models.AlbumPhoto, out *viewmodels.AlbumPhotoDto) {
	out.Id = in.ID
	out.Title = in.Title
	out.Image = in.Image
	out.Thumbnail = in.Thumbnail

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}
