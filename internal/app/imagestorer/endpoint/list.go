package endpoint

import (
	"fmt"
	"io"

	tapi "github.com/khusainnov/tag/pkg/tages-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (e *Endpoint) ListImages(stream tapi.ImageService_ListImagesServer) error {
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return status.Error(codes.InvalidArgument, "empty data")
		}
		if err != nil {
			return status.Error(codes.Internal, "cannot get request")
		}
		rsp, err := e.srv.ListImage()
		if err != nil {
			return status.Error(codes.Internal, fmt.Errorf("cannot load list of images, %w", err).Error())
		}

		if err = stream.SendAndClose(rsp); err != nil {
			return status.Error(codes.Internal, fmt.Errorf("cannot send response data, %w", err).Error())
		}

		return nil
	}
}
