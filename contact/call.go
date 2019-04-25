package contact

import (
	"context"
	"github.com/autom8ter/api"
	"github.com/autom8ter/backend/clientset"
	"github.com/autom8ter/backend/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Caller struct {
	c *clientset.ClientSet
}

func NewCaller(c *config.Config) *Caller {
	return &Caller{
		c: clientset.NewClientSet(c),
	}
}

func (c *Caller) SendCall(ctx context.Context, m *api.Call) (*api.Identifier, error) {
	resp, ex, err := c.c.Twilio.CallWithApplicationCallbacks(m.From, m.To, m.App)
	api.Util.Entry().Debugln(string(api.Util.MarshalYAML(ex)))
	if err != nil {
		return nil, status.Errorf(codes.Internal, api.Util.WrapErr(err, string(api.Util.MarshalJSON(ex))).Error())
	}
	return &api.Identifier{
		Id: resp.Uri,
	}, nil
}
