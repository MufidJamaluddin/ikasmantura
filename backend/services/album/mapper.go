package album

import (
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
)

func toModel(data *viewmodels.AlbumDto, out *models.Album) {
	out.ID = uint(data.Id)
	out.Title = data.Title

	utils.FillCreated(data, out)
	utils.FillUpdated(data, out)
}

func toViewModel(in *models.Album, out *viewmodels.AlbumDto) {
	out.Id = int(in.ID)
	out.Title = in.Title

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}
