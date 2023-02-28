package adapters

import (
	"github.com/khusainnov/tag/internal/model"
	tapi "github.com/khusainnov/tag/pkg/tages-api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ImageToPb(data *model.Image) *tapi.Image {
	return &tapi.Image{
		Name:         data.Name,
		CreatedDate:  timestamppb.New(data.CreatedAt),
		ModifiedDate: timestamppb.New(data.EditedAt),
	}
}

func ListImageToPb(list []*model.Image) *tapi.ListImagesResponse {
	var resp = make([]*tapi.Image, len(list))

	for i, v := range list {
		resp[i] = ImageToPb(v)
	}

	return &tapi.ListImagesResponse{Images: resp}
}
