package contact

import (
	"context"
	"github.com/autom8ter/api"
	"github.com/autom8ter/backend/clients"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SMSer struct{}

func (s *SMSer) GetSMS(ctx context.Context, r *api.Identifier) (*api.SMSStatus, error) {
	panic("implement me")
}

func (s *SMSer) SendSMS(ctx context.Context, m *api.SMS) (*api.Identifier, error) {
	resp, ex, err := clients.Twilio.SendSMSWithCopilot(clients.TWILIO_MESSAGING_SERVICE, m.To, m.Message.Value, "", "")
	clients.Util.Entry().Debugln(string(clients.Util.MarshalYAML(ex)))
	if err != nil {
		return nil, status.Errorf(codes.Internal, clients.Util.WrapErr(err, string(clients.Util.MarshalJSON(ex))).Error())
	}
	return &api.Identifier{
		Id: resp.Url,
	}, nil
}
