package endpoint

import (
	"context"
	"fmt"

	tapi "github.com/khusainnov/tag/pkg/tages-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (e *Endpoint) ListImages(ctx context.Context, _ *emptypb.Empty) (*tapi.ListImagesResponse, error) {
	rsp, err := e.srv.ListImage()
	if err != nil {
		return &tapi.ListImagesResponse{}, status.Error(codes.Internal, fmt.Errorf("cannot load list of images, %w", err).Error())
	}

	return rsp, ctx.Err()

}
