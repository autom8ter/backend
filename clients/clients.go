package clients

import (
	"context"
	"github.com/autom8ter/gcloud"
	"github.com/autom8ter/objectify"
	"os"
)

func init() {
	GCP = gcloud.NewGCP()
	Context = context.WithValue(context.TODO(), "env", os.Environ())

}

var Util = objectify.Default()
var GCP *gcloud.GCP
var Context context.Context
