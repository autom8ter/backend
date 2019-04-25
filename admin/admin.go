package admin

import (
	"context"
	"github.com/autom8ter/api"
	"github.com/autom8ter/backend/cache"
	"github.com/autom8ter/backend/clientset"
	"github.com/autom8ter/backend/config"
	"github.com/autom8ter/engine/driver"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

type Admin struct {
	email *mail.Email
	c     *clientset.ClientSet
	driver.PluginFunc
}

func NewAdmin(fromemail string, fromname string, c *config.Config) *Admin {
	a := &Admin{
		email: &mail.Email{
			Name:    fromname,
			Address: fromemail,
		},
		c: clientset.NewClientSet(c),
	}
	a.PluginFunc = func(s *grpc.Server) {
		api.RegisterAdminServiceServer(s, a)
	}
	return a
}

func (a *Admin) GetDashboard(ctx context.Context, secret *api.Secret) (*api.Dashboard, error) {
	if secret.Text == os.Getenv("SECRET") {
		return &api.Dashboard{
			Users: &api.UsersWidget{
				Count: int64(cache.Working.TotalUsers()),
			},
			Customers: &api.CustomersWidget{
				Count: int64(cache.Working.TotalStripeCustomers()),
			},
			Plans: &api.PlansWidget{
				Count: int64(cache.Working.TotalStripePlans()),
			},
			Subscriptions: &api.SubscriptionsWidget{
				Count: int64(cache.Working.TotalSubscriptions()),
			},
			Charges: &api.ChargesWidget{
				Count: int64(cache.Working.TotalSubscriptions()),
			},
		}, nil
	} else {
		return nil, status.Errorf(codes.Unauthenticated, "provided secret does not match")
	}
}

func (a *Admin) EmailUser(ctx context.Context, email *api.Email) (*api.Message, error) {
	m := mail.NewSingleEmail(a.email, email.Subject, &mail.Email{
		Name:    email.Name,
		Address: email.Address,
	}, email.Plain, email.Html)
	resp, err := a.c.Sendgrid.Send(m)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &api.Message{
		Value: resp.Body,
	}, nil
}
