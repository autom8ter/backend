package contact

import (
	"context"
	"fmt"
	"github.com/autom8ter/api"
	"github.com/autom8ter/api/common"
	"github.com/autom8ter/backend/clientset"
	"github.com/autom8ter/backend/config"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SMSer struct {
	c *clientset.ClientSet
}

func (s *SMSer) SendSMSBlast(blast *api.SMSBlast, stream api.ContactService_SendSMSBlastServer) error {
	for _, t := range blast.To.Strings {
		resp, ex, err := s.c.Twilio.SendSMSWithCopilot(blast.Service.Text, t.Text, blast.Message.Text, blast.Callback.Text, blast.App.Text)
		if err != nil {
			return status.Errorf(codes.Internal, errors.Wrap(err, fmt.Sprintf("%v", ex)).Error())
		}
		if err := stream.Send(&api.SMSResponse{
			Id: &common.Identifier{
				Id: common.ToString(resp.Sid),
			},
			To:       common.ToString(resp.To),
			From:     common.ToString(resp.From),
			MediaUrl: common.ToString(resp.MediaUrl),
			Body:     common.ToString(resp.Body),
			Status:   common.ToString(resp.Status),
			Annotations: common.ToStringMap(map[string]string{
				"date_created": resp.DateCreated,
				"date_updated": resp.DateUpdate,
			}),
		}); err != nil {
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

func (s *SMSer) GetSMS(ctx context.Context, r *common.Identifier) (*api.SMSResponse, error) {
	resp, ex, err := s.c.Twilio.GetSMS(r.Id.Text)
	if err != nil {
		return nil, status.Errorf(codes.Internal, errors.Wrap(err, fmt.Sprintf("%v", ex)).Error())
	}
	return &api.SMSResponse{
		Id: &common.Identifier{
			Id: common.ToString(resp.Sid),
		},
		To:       common.ToString(resp.To),
		From:     common.ToString(resp.From),
		MediaUrl: common.ToString(resp.MediaUrl),
		Body:     common.ToString(resp.Body),
		Status:   common.ToString(resp.Status),
		Annotations: common.ToStringMap(map[string]string{
			"date_created": resp.DateCreated,
			"date_updated": resp.DateUpdate,
		}),
	}, nil

}

func (s *SMSer) SendSMS(ctx context.Context, m *api.SMS) (*api.SMSResponse, error) {
	resp, ex, err := s.c.Twilio.SendSMSWithCopilot(m.Service.Text, m.To.Text, m.Message.Text, m.Callback.Text, m.App.Text)
	if err != nil {
		return nil, status.Errorf(codes.Internal, errors.Wrap(err, fmt.Sprintf("%v", ex)).Error())
	}
	return &api.SMSResponse{
		Id: &common.Identifier{
			Id: common.ToString(resp.Sid),
		},
		To:       common.ToString(resp.To),
		From:     common.ToString(resp.From),
		MediaUrl: common.ToString(resp.MediaUrl),
		Body:     common.ToString(resp.Body),
		Status:   common.ToString(resp.Status),
		Annotations: common.ToStringMap(map[string]string{
			"date_created": resp.DateCreated,
			"date_updated": resp.DateUpdate,
		}),
	}, nil
}
