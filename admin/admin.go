package admin

import (
	"context"
	"github.com/autom8ter/api"
	"github.com/autom8ter/api/common"
	"github.com/autom8ter/backend/cache"
	"github.com/autom8ter/engine/driver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewAdmin() *Admin {
	a := &Admin{}
	a.PluginFunc = func(s *grpc.Server) {
		api.RegisterAdminServiceServer(s, a)
	}
	return a
}

type Admin struct {
	driver.PluginFunc
}

func (a *Admin) StartCache(ctx context.Context, r *api.StartCacheRequest) (*common.Empty, error) {
	if cache.Working.Looping && r.Frequency.ParseDuration() == cache.Working.SyncFrequency {
		return nil, status.Error(codes.AlreadyExists, "cache is already running with the requested frequency")
	}
	if cache.Working.Looping && r.Frequency.ParseDuration() != cache.Working.SyncFrequency {
		cache.Working.SyncFrequency = r.Frequency.ParseDuration()
		return &common.Empty{}, nil
	}
	if !cache.Working.Looping {
		cache.Init(r.Frequency.ParseDuration())
		return &common.Empty{}, nil
	}
	return nil, nil
}

func (a *Admin) StopCache(ctx context.Context, r *common.Empty) (*common.Empty, error) {
	if !cache.Working.Looping {
		return nil, status.Error(codes.AlreadyExists, "cache is already stopped")
	}
	cache.Working.Looping = false
	return &common.Empty{}, nil
}
