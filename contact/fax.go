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

type Faxer struct {
	c *clientset.ClientSet
}

func NewFaxer(c *config.Config) *Faxer {
	return &Faxer{
		c: clientset.NewClientSet(c),
	}
}

func (f *Faxer) SendFax(ctx context.Context, r *api.FaxRequest) (*common.Bytes, error) {
	resp, ex, err := f.c.Twilio.SendFax(r.To.Text, r.From.Text, r.MediaUrl.Text, r.Quality.Text, r.Callback.Text, r.StoreMedia.Answer)
	if err != nil {
		return nil, status.Errorf(codes.Internal, errors.Wrap(err, fmt.Sprintf("%v", ex)).Error())
	}
	return common.AsBytes(resp), nil
}
