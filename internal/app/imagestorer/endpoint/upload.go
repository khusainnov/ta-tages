package endpoint

import (
	"context"
	"fmt"

	tapi "github.com/khusainnov/tag/pkg/tages-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (e *Endpoint) UploadImage(ctx context.Context, req *tapi.UploadImageRequest) (*emptypb.Empty, error) {
	if len(req.Image) == 0 {
		return nil, status.Error(codes.InvalidArgument, "request don't provide any data")
	}

	_, err := e.srv.UploadImage(req.Image)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("cannot upload the image, %w", err).Error())
	}

	return new(emptypb.Empty), ctx.Err()
}
