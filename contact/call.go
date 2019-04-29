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

type Caller struct {
	c *clientset.ClientSet
}

func (c *Caller) SendCallBlast(b *api.CallBlast, stream api.ContactService_SendCallBlastServer) error {
	for _, t := range b.To.Strings {
		resp, ex, err := c.c.Twilio.CallWithApplicationCallbacks(b.From.Text, t.Text, b.App.Text)
		if err != nil {
			return status.Errorf(codes.Internal, errors.Wrap(err, fmt.Sprintf("%v", ex)).Error())
		}
		if err := stream.Send(&api.CallResponse{
			Id:            &common.Identifier{Id: common.ToString(resp.Sid)},
			To:            common.ToString(resp.To),
			From:          common.ToString(resp.From),
			Status:        common.ToString(resp.Status),
			AnsweredBy:    common.ToString(resp.AnsweredBy),
			ForwardedFrom: common.ToString(resp.ForwardedFrom),
			CallerName:    common.ToString(resp.CallerName),
			Annotations: common.ToStringMap(map[string]string{
				"annotation":       resp.Annotation,
				"date_created":     resp.DateCreated,
				"date_updated":     resp.DateUpdated,
				"end_time":         resp.EndTime,
				"phone_number_sid": resp.PhoneNumberSid,
			}),
		}); err != nil {
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

func (c *Caller) SendCall(ctx context.Context, m *api.Call) (*api.CallResponse, error) {
	resp, ex, err := c.c.Twilio.CallWithApplicationCallbacks(m.From.Text, m.To.Text, m.App.Text)
	if err != nil {
		return nil, status.Errorf(codes.Internal, errors.Wrap(err, fmt.Sprintf("%v", ex)).Error())
	}
	return &api.CallResponse{
		Id:            &common.Identifier{Id: common.ToString(resp.Sid)},
		To:            common.ToString(resp.To),
		From:          common.ToString(resp.From),
		Status:        common.ToString(resp.Status),
		AnsweredBy:    common.ToString(resp.AnsweredBy),
		ForwardedFrom: common.ToString(resp.ForwardedFrom),
		CallerName:    common.ToString(resp.CallerName),
		Annotations: common.ToStringMap(map[string]string{
			"annotation":       resp.Annotation,
			"date_created":     resp.DateCreated,
			"date_updated":     resp.DateUpdated,
			"end_time":         resp.EndTime,
			"phone_number_sid": resp.PhoneNumberSid,
		}),
	}, nil
}
