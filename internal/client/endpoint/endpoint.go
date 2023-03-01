package endpoint

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"sync/atomic"

	tapi "github.com/khusainnov/tag/pkg/tages-api"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// limits for client connections
const (
	maxUpload   = 10
	maxList     = 100
	maxDownload = 10
)

var (
	uploadCounter   uint32 = 0
	listCounter     uint32 = 0
	downloadCounter uint32 = 0
)

type Endpoint struct {
	L      *zap.Logger
	client tapi.ImageServiceClient
}

func NewEndpoint(client tapi.ImageServiceClient, log *zap.Logger) *Endpoint {
	return &Endpoint{
		L:      log,
		client: client,
	}
}

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

func (e *Endpoint) List(ctx context.Context) error {
	if listCounter > maxList {
		return status.Error(codes.Aborted, "limit exceeded")
	}

	atomic.AddUint32(&listCounter, 1)

	req := &emptypb.Empty{}

	resp, err := e.client.ListImages(ctx, req)
	if err != nil {
		return fmt.Errorf("cannot get list of images, %w", err)
	}

	for _, v := range resp.Images {
		fmt.Fprintf(os.Stdout, "%v\n", v)
	}

	listCounter--

	return nil
}

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
