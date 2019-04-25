package clients

import (
	"context"
	"github.com/autom8ter/gcloud"
	"github.com/autom8ter/objectify"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sfreiberg/gotwilio"
	"google.golang.org/api/option"
	"log"
	"os"
)

func init() {
	if _, err := os.Stat(GCP_CREDENTIALS); err != nil {
		log.Fatalln("expected GCP credentials.json in $PWD")
	}
	GCP = gcloud.NewGCP(option.WithCredentialsFile(GCP_CREDENTIALS))
	Context = context.WithValue(context.TODO(), "env", os.Environ())
	Twilio = gotwilio.NewTwilioClient(TWILIO_ACCOUNT, TWILIO_KEY)
	SendGrid = sendgrid.NewSendClient(SENDGRID_KEY)
}

var GCP_CREDENTIALS = "credentials.json"
var SENDGRID_KEY = os.Getenv("SENDGRID_KEY")
var SENDGRID_EMAIL = os.Getenv("SENDGRID_EMAIL")
var SENDGRID_NAME = os.Getenv("SENDGRID_NAME")
var TWILIO_ACCOUNT = os.Getenv("TWILIO_ACCOUNT")
var TWILIO_KEY = os.Getenv("TWILIO_KEY")
var TWILIO_CALL_APP = os.Getenv("TWILIO_CALL_APP")
var TWILIO_CALL_NUMBER = os.Getenv("TWILIO_CALL_NUMBER")
var TWILIO_MESSAGING_SERVICE = os.Getenv("TWILIO_MESSAGING_SERVICE")
var TWILIO_APPLICATION = os.Getenv("TWILIO_APPLICATION")
var TWILIO_SMS_CALLBACK = os.Getenv("TWILIO_SMS_CALLBACK")

var SendGrid *sendgrid.Client
var Twilio *gotwilio.Twilio
var Util = objectify.Default()
var GCP *gcloud.GCP
var Context context.Context
