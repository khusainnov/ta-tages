package endpoint

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/khusainnov/tag/internal/client/internal"
	tapi "github.com/khusainnov/tag/pkg/tages-api"
	"go.uber.org/zap"
	"golang.org/x/sync/semaphore"
	"google.golang.org/protobuf/types/known/emptypb"
)

// limits for client connections
const (
	maxUpload = 10
	maxList   = 100
)

type Endpoint struct {
	sync.WaitGroup
	client tapi.ImageServiceClient
}

func NewEndpoint(client tapi.ImageServiceClient) *Endpoint {
	return &Endpoint{client: client}
}

func (e *Endpoint) Upload(ctx context.Context, log *zap.Logger) error {
	var count int

	cc, err := e.client.UploadImage(ctx)
	if err != nil {
		return fmt.Errorf("cannot upload the image, %w", err)
	}

	semaphor := semaphore.NewWeighted(int64(maxUpload))

	images := internal.Images
	for count < maxUpload {
		for i := 0; i < len(images)-1; i++ {
			if err = semaphor.Acquire(ctx, 1); err != nil {
				log.Error("cannot acquire semaphore", zap.Error(err))
				continue
			}

			e.Add(1)
			go func(m int) {
				// release the semaphore when goroutine exits
				defer func() {
					semaphor.Release(1)
					e.Done()
				}()

				req := &tapi.UploadImageRequest{
					Image: images[m],
				}

				if err = cc.Send(req); err != nil {
					log.Error("cannot send upload request", zap.Error(err))
					return
				}

				count++
			}(i)

		}
	}
	e.Wait()

	return nil
}

func (e *Endpoint) List(ctx context.Context, log *zap.Logger) error {
	var count int

	cc, err := e.client.ListImages(ctx)
	if err != nil {
		return fmt.Errorf("cannot get list of images, %w", err)
	}

	semaphor := semaphore.NewWeighted(int64(maxList))

	for count < maxList {
		if err = semaphor.Acquire(ctx, 1); err != nil {
			log.Error("cannot acquire semaphore", zap.Error(err))
			continue
		}

		e.Add(1)
		go func() {
			// release the semaphore when goroutine exits
			defer func() {
				semaphor.Release(1)
				e.Done()
			}()

			req := &emptypb.Empty{}
			if err = cc.Send(req); err != nil {
				log.Error("cannot send list request", zap.Error(err))
				return
			}

			count++
		}()

	}

	rsp := &tapi.ListImagesResponse{}
	err = cc.RecvMsg(rsp)
	if err == io.EOF {
		return err
	}
	if err != nil {
		log.Error("cannot receive response list", zap.Error(err))
		return err
	}

	for _, v := range rsp.Images {
		fmt.Fprintf(os.Stdout, "%v\n", v)
	}

	return nil
}
