package contact

import (
	"context"
	"github.com/autom8ter/api"
	"github.com/autom8ter/backend/clientset"
	"github.com/autom8ter/backend/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SMSer struct {
	c *clientset.ClientSet
}

func NewSMSer(c *config.Config) *SMSer {
	return &SMSer{
		c: clientset.NewClientSet(c),
	}
}

func (s *SMSer) GetSMS(ctx context.Context, r *api.Identifier) (*api.SMSStatus, error) {
	resp, ex, err := s.c.Twilio.GetSMS(r.Id)
	api.Util.Entry().Debugln(string(api.Util.MarshalYAML(ex)))
	if err != nil {
		return nil, status.Errorf(codes.Internal, api.Util.WrapErr(err, string(api.Util.MarshalJSON(ex))).Error())
	}
	return &api.SMSStatus{
		Id: &api.Identifier{
			Id: resp.Sid,
		},
		Sms: &api.SMS{
			To: resp.To,
			Message: &api.Message{
				Value: resp.Body,
			},
			MediaURL: resp.MediaUrl,
		},
		Status: resp.Status,
		Uri:    resp.Url,
	}, nil

}

func (s *SMSer) SendSMS(ctx context.Context, m *api.SMS) (*api.Identifier, error) {
	resp, ex, err := s.c.Twilio.SendSMSWithCopilot(m.Service, m.To, m.Message.Value, m.Callback, m.App)
	api.Util.Entry().Debugln(string(api.Util.MarshalYAML(ex)))
	if err != nil {
		return nil, status.Errorf(codes.Internal, api.Util.WrapErr(err, string(api.Util.MarshalJSON(ex))).Error())
	}
	return &api.Identifier{
		Id: resp.Url,
	}, nil
}
