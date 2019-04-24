package echo

import (
	"fmt"
	"github.com/autom8ter/api"
	"context"
	"github.com/autom8ter/engine/driver"
	"google.golang.org/grpc"
)

type Echoer struct {
	driver.PluginFunc
}

func NewEchoer() *Echoer {
	e := &Echoer{}
	e.PluginFunc = func(s *grpc.Server) {
		api.RegisterEchoServiceServer(s, e)
	}
	return e
}

func (b *Echoer) Echo(ctx context.Context, e *api.EchoMessage) (*api.EchoMessage, error) {
	return &api.EchoMessage{
		Value:                fmt.Sprintf("echoed: %s", e.Value),
	}, nil

}

