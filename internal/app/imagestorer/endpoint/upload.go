package endpoint

import (
	"fmt"

	tapi "github.com/khusainnov/tag/pkg/tages-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (e *Endpoint) UploadImage(stream tapi.ImageService_UploadImageServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return status.Error(codes.InvalidArgument, "cannot get request")
		}

		fmt.Println(req.GetImage()[:50])

		resp, err := e.srv.UploadImage(req.GetImage())
		if err != nil {
			return status.Error(codes.Internal, fmt.Errorf("cannot upload the image, %w", err).Error())
		}

		if err = stream.SendAndClose(resp); err != nil {
			return status.Error(codes.Internal, fmt.Errorf("cannot send response data, %w", err).Error())
		}

		return nil
	}
}
