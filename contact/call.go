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

func (c *Caller) SendCallBlast(b *api.CallBlast, stream api.ContactService_SendCallBlastServer) error {
	for _, t := range b.To {
		resp, ex, err := c.c.Twilio.CallWithApplicationCallbacks(b.From, t, b.App)
		api.Util.Entry().Debugln(string(api.Util.MarshalYAML(ex)))
		if err != nil {
			return status.Errorf(codes.Internal, api.Util.WrapErr(err, string(api.Util.MarshalJSON(ex))).Error())
		}
		if err := stream.Send(api.AsBytes(resp)); err != nil {
			return err
		}
	}
	return nil
}

func NewCaller(c *config.Config) *Caller {
	return &Caller{
		c: clientset.NewClientSet(c),
	}
}

func (c *Caller) SendCall(ctx context.Context, m *api.Call) (*api.Bytes, error) {
	resp, ex, err := c.c.Twilio.CallWithApplicationCallbacks(m.From, m.To, m.App)
	api.Util.Entry().Debugln(string(api.Util.MarshalYAML(ex)))
	if err != nil {
		return nil, status.Errorf(codes.Internal, api.Util.WrapErr(err, string(api.Util.MarshalJSON(ex))).Error())
	}
	return api.AsBytes(resp), nil
}
