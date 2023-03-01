package endpoint

import (
	"context"
	"fmt"
	"sync/atomic"

	tapi "github.com/khusainnov/tag/pkg/tages-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (e *Endpoint) Upload(ctx context.Context, image []byte) error {
	if uploadCounter > maxUpload {
		return status.Error(codes.Aborted, "limit exceeded")
	}

	atomic.AddUint32(&uploadCounter, 1)

	req := &tapi.UploadImageRequest{
		Image: image,
	}

	_, err := e.client.UploadImage(ctx, req)
	if err != nil {
		return fmt.Errorf("cannot upload the image, %w", err)
	}

	uploadCounter--

	return nil
}
