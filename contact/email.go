package contact

import (
	"context"
	"github.com/autom8ter/api"
	"github.com/autom8ter/backend/clientset"
	"github.com/autom8ter/backend/config"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Emailer struct {
	c *clientset.ClientSet
}

func NewEmailer(c *config.Config) *Emailer {
	return &Emailer{
		c: clientset.NewClientSet(c),
	}
}

func (e *Emailer) SendEmail(ctx context.Context, email *api.EmailRequest) (*api.Message, error) {
	from := mail.NewEmail(email.FromName, email.FromEmail)
	to := mail.NewEmail(email.Email.Name, email.Email.Address)
	message := mail.NewSingleEmail(from, email.Email.Subject, to, email.Email.Plain, email.Email.Html)
	resp, err := e.c.Sendgrid.Send(message)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &api.Message{
		Value: resp.Body,
	}, nil
}
