package contact

import (
	"context"
	"github.com/autom8ter/api"
	"github.com/autom8ter/backend/clients"
	"github.com/sfreiberg/gotwilio"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Caller struct{}

func (*Caller) SendCall(ctx context.Context, m *api.Call) (*api.Identifier, error) {

	resp, ex, err := clients.Twilio.CallWithUrlCallbacks(clients.TWILIO_CALL_NUMBER, m.To, gotwilio.NewCallbackParameters(m.Callback))
	clients.Util.Entry().Debugln(string(clients.Util.MarshalYAML(ex)))
	if err != nil {
		return nil, status.Errorf(codes.Internal, clients.Util.WrapErr(err, string(clients.Util.MarshalJSON(ex))).Error())
	}
	return &api.Identifier{
		Id: resp.Uri,
	}, nil
}
