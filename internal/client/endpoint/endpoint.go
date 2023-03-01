package endpoint

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"

	tapi "github.com/khusainnov/tag/pkg/tages-api"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// limits for client connections
const (
	maxUpload = 10
	maxList   = 100
)

var (
	uploadCounter uint32 = 0
	listCounter   uint32 = 0
)

type Endpoint struct {
	L      *zap.Logger
	mu     *sync.Mutex
	wg     sync.WaitGroup
	client tapi.ImageServiceClient
}

func NewEndpoint(client tapi.ImageServiceClient, log *zap.Logger) *Endpoint {
	return &Endpoint{
		L:      log,
		client: client,
	}
}

func (e *Endpoint) Upload(ctx context.Context, image []byte) error {
	atomic.AddUint32(&uploadCounter, 1)

	if uploadCounter > maxUpload {
		return status.Error(codes.Aborted, "limit reached")
	}

	cc, err := e.client.UploadImage(ctx)
	if err != nil {
		return fmt.Errorf("cannot upload the image, %w", err)
	}

	req := &tapi.UploadImageRequest{
		Image: image,
	}

	if err = cc.Send(req); err != nil {
		e.L.Error("cannot send upload request", zap.Error(err))
		return fmt.Errorf("cannot send upload request, %w", err)
	}

	uploadCounter--

	return nil
}

func (e *Endpoint) List(ctx context.Context) error {
	atomic.AddUint32(&listCounter, 1)

	cc, err := e.client.ListImages(ctx)
	if err != nil {
		return fmt.Errorf("cannot get list of images, %w", err)
	}

	rsp := &tapi.ListImagesResponse{}

	req := &emptypb.Empty{}
	if err = cc.Send(req); err != nil {
		e.L.Error("cannot send list request", zap.Error(err))
		return err
	}

	err = cc.RecvMsg(rsp)
	if err == io.EOF {
		return err
	}
	if err != nil {
		e.L.Error("cannot receive response list", zap.Error(err))
		return err
	}

	for _, v := range rsp.Images {
		fmt.Fprintf(os.Stdout, "%v\n", v)
	}

	listCounter--

	return nil
}
