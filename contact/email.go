package contact

import (
	"context"
	"github.com/autom8ter/api"
	"github.com/autom8ter/api/common"
	"github.com/autom8ter/backend/clientset"
	"github.com/autom8ter/backend/config"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Emailer struct {
	c *clientset.ClientSet
}

func (e *Emailer) SendEmailBlast(email *api.EmailBlastRequest, stream api.ContactService_SendEmailBlastServer) error {
	for k, v := range email.Blast.NameAddress.StringMap {
		from := mail.NewEmail(email.FromName.Text, email.FromEmail.Text)
		to := mail.NewEmail(k, v.Text)
		message := mail.NewSingleEmail(from, email.Blast.Subject.Text, to, email.Blast.Plain.Text, email.Blast.Html.Text)
		resp, err := e.c.Sendgrid.Send(message)
		if err != nil {
			return err
		}
		if err := stream.Send(&common.String{
			Text: resp.Body,
		}); err != nil {
			return common.ToError(common.ToError(err, resp.Body), string(resp.StatusCode))
		}
	}
	return nil
}

func NewEmailer(c *config.Config) *Emailer {
	return &Emailer{
		c: clientset.NewClientSet(c),
	}
}

func (e *Emailer) SendEmail(ctx context.Context, email *api.EmailRequest) (*common.String, error) {
	from := mail.NewEmail(email.FromName.Text, email.FromEmail.Text)
	to := mail.NewEmail(email.Email.Name.Text, email.Email.Address.Text)
	message := mail.NewSingleEmail(from, email.Email.Subject.Text, to, email.Email.Plain.Text, email.Email.Html.Text)
	resp, err := e.c.Sendgrid.Send(message)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &common.String{
		Text: resp.Body,
	}, nil
}
