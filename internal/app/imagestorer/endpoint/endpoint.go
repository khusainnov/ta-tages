package endpoint

import (
	"github.com/khusainnov/tag/internal/app/imagestorer/service"
	tapi "github.com/khusainnov/tag/pkg/tages-api"
)

type Endpoint struct {
	tapi.UnimplementedImageServiceServer
	srv *service.Service
}

func NewEndpoint(srv *service.Service) *Endpoint {
	return &Endpoint{srv: srv}
}
