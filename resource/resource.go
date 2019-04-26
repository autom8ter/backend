package resource

import (
	"context"
	"github.com/autom8ter/api"
	"github.com/autom8ter/engine/driver"
	"google.golang.org/grpc"
)

type Resource struct {
	driver.PluginFunc
}

func NewResource() *Resource {
	r := &Resource{}
	r.PluginFunc = func(s *grpc.Server) {
		api.RegisterResourceServiceServer(s, r)
	}
	return r
}

func (Resource) GetResource(ctx context.Context, request *api.ResourceRequest) (*api.Bytes, error) {
	return request.Do()
}
