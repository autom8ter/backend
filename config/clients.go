package config

import (
	"github.com/autom8ter/api"
	"github.com/pkg/errors"
	"os"
)

type Config struct {
	GCPServiceCredentialsPath string `validate:"required"`
	SendgridKey               string `validate:"required"`
	TwilioAccount             string `validate:"required"`
	TwilioKey                 string `validate:"required"`
	StripeKey                 string `validate:"required"`
	Auth0Domain               string `validate:"required"`
	Auth0ClientID             string `validate:"required"`
	Auth9ClientSecret         string `validate:"required"`
}

func (c *Config) Validate() error {
	if _, err := os.Stat(c.GCPServiceCredentialsPath); err != nil {
		return errors.Wrap(err, "gcp service account credentials not found. path: "+c.GCPServiceCredentialsPath)
	}
	return api.Util.Validate(c)
}

func FromEnv(gcpCredsPath string) *Config {
	return &Config{
		GCPServiceCredentialsPath: gcpCredsPath,
		SendgridKey:               os.Getenv("SENDGRID_KEY"),
		TwilioAccount:             os.Getenv("TWILIO_ACCOUNT"),
		TwilioKey:                 os.Getenv("TWILIO_KEY"),
		StripeKey:                 os.Getenv("STRIPE_KEY"),
		Auth0Domain:               os.Getenv("AUTH0_DOMAIN"),
		Auth0ClientID:             os.Getenv("AUTH0_CLIENT_ID"),
		Auth9ClientSecret:         os.Getenv("AUTH0_CLIENT_SECRET"),
	}
}
