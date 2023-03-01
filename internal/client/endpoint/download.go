package endpoint

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"sync/atomic"

	tapi "github.com/khusainnov/tag/pkg/tages-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (e *Endpoint) Download(ctx context.Context, id string) error {
	if downloadCounter > maxDownload {
		return status.Error(codes.Aborted, "limit exceeded")
	}
	atomic.AddUint32(&downloadCounter, 1)

	req := &tapi.DownloadImageRequest{
		Id: id,
	}

	resp, err := e.client.DownloadImage(ctx, req)
	if err != nil {
		return fmt.Errorf("cannot download image, %w", err)
	}

	image := base64.StdEncoding.EncodeToString(resp.Image)
	fmt.Fprintf(os.Stdout, "%s\n", image)

	downloadCounter--

	return nil
}
