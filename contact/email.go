package contact

import (
	"context"
	"github.com/autom8ter/api"
	"github.com/autom8ter/backend/clients"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Emailer struct{}

func (e *Emailer) SendEmail(ctx context.Context, email *api.Email) (*api.Message, error) {
	from := mail.NewEmail(clients.SENDGRID_NAME, clients.SENDGRID_EMAIL)
	to := mail.NewEmail(email.Name, email.Address)
	message := mail.NewSingleEmail(from, email.Subject, to, email.Plain, email.Html)
	resp, err := clients.SendGrid.Send(message)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &api.Message{
		Value: resp.Body,
	}, nil
}
