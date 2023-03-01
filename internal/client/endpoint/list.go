package endpoint

import (
	"context"
	"fmt"
	"os"
	"sync/atomic"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

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
