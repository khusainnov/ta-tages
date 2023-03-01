package endpoint

import (
	"context"

	tapi "github.com/khusainnov/tag/pkg/tages-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (e *Endpoint) DownloadImage(ctx context.Context, req *tapi.DownloadImageRequest) (*tapi.DownloadImageResponse, error) {
	if req.Id == "" {
		return &tapi.DownloadImageResponse{}, status.Error(codes.InvalidArgument, "empty image name")
	}

	resp, err := e.srv.DownloadImage(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return resp, ctx.Err()
}
