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

func (s *SMSer) SendSMSBlast(blast *api.SMSBlast, stream api.ContactService_SendSMSBlastServer) error {
	for _, t := range blast.To {
		resp, ex, err := s.c.Twilio.SendSMSWithCopilot(blast.Service, t, blast.Message.Value, blast.Callback, blast.App)
		api.Util.Entry().Debugln(string(api.Util.MarshalYAML(ex)))
		if err != nil {
			return status.Errorf(codes.Internal, api.Util.WrapErr(err, string(api.Util.MarshalJSON(ex))).Error())
		}
		if err := stream.Send(api.AsBytes(resp)); err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}
	return nil
}

func NewSMSer(c *config.Config) *SMSer {
	return &SMSer{
		c: clientset.NewClientSet(c),
	}
}

func (s *SMSer) GetSMS(ctx context.Context, r *api.Identifier) (*api.Bytes, error) {
	resp, ex, err := s.c.Twilio.GetSMS(r.Id)
	api.Util.Entry().Debugln(string(api.Util.MarshalYAML(ex)))
	if err != nil {
		return nil, status.Errorf(codes.Internal, api.Util.WrapErr(err, string(api.Util.MarshalJSON(ex))).Error())
	}
	return api.AsBytes(resp), nil

}

func (s *SMSer) SendSMS(ctx context.Context, m *api.SMS) (*api.Bytes, error) {
	resp, ex, err := s.c.Twilio.SendSMSWithCopilot(m.Service, m.To, m.Message.Value, m.Callback, m.App)
	api.Util.Entry().Debugln(string(api.Util.MarshalYAML(ex)))
	if err != nil {
		return nil, status.Errorf(codes.Internal, api.Util.WrapErr(err, string(api.Util.MarshalJSON(ex))).Error())
	}
	return api.AsBytes(resp), nil
}
