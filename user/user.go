package user

import (
	"context"
	"github.com/autom8ter/api"
	"github.com/autom8ter/backend/cache"
	"github.com/autom8ter/engine/driver"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewUser() *User {
	u := &User{}
	u.PluginFunc = func(s *grpc.Server) {
		api.RegisterUserServiceServer(s, u)
	}
	return u
}

func (u *User) createStripeCustomer(ctx context.Context, info api.UserInfo) (*api.Identifier, error) {
	c, err := customer.New(&stripe.CustomerParams{
		Description: stripe.String(string(api.Util.MarshalJSON(info.AppMetadata))),
		Email:       stripe.String(info.Email),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	cache.Working.AddStripeCustomer(c)
	return &api.Identifier{
		Id: c.ID,
	}, nil
}

type User struct {
	driver.PluginFunc
}

func (u *User) GetUser(ctx context.Context, message *api.Identifier) (*api.UserInfo, error) {
	panic("implement me")
}

func (u *User) UpdateUser(ctx context.Context, info *api.UserInfo) (*api.Identifier, error) {
	panic("implement me")
}

func (u *User) CreateUser(ctx context.Context, info *api.UserInfo) (*api.Identifier, error) {
	panic("implement me")
}

func (u *User) DeleteUser(ctx context.Context, identifier *api.Identifier) (*api.Identifier, error) {
	panic("implement me")
}
