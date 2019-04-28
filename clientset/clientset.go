package clientset

import (
	"github.com/autom8ter/api/common"
	"github.com/autom8ter/backend/config"
	"github.com/autom8ter/gcloud"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sfreiberg/gotwilio"
	"github.com/stripe/stripe-go"
	"google.golang.org/api/option"
)

type ClientSet struct {
	Twilio   *gotwilio.Twilio
	GCP      *gcloud.GCP
	Sendgrid *sendgrid.Client
	// Auth0 *auth0.Client
}

func NewClientSet(c *config.Config) *ClientSet {
	if err := c.Validate(); err != nil {
		common.Util.Entry().Warnln("validating config:", err.Error())
	}
	stripe.Key = c.StripeKey
	return &ClientSet{
		Twilio:   gotwilio.NewTwilioClient(c.TwilioAccount, c.TwilioKey),
		GCP:      gcloud.NewGCP(option.WithCredentialsFile(c.GCPServiceCredentialsPath)),
		Sendgrid: sendgrid.NewSendClient(c.SendgridKey),
	}
}
