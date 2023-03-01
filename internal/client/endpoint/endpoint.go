package endpoint

import (
	tapi "github.com/khusainnov/tag/pkg/tages-api"
	"go.uber.org/zap"
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
