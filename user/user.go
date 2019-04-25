package user

import (
	"context"
	"github.com/autom8ter/api"
	"github.com/autom8ter/backend/cache"
	"github.com/autom8ter/engine/driver"
	"github.com/pkg/errors"
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

func (u *User) createStripeCustomer(ctx context.Context, info *api.User) (*api.Identifier, error) {
	c, err := customer.New(&stripe.CustomerParams{
		Description: stripe.String(string(api.Util.MarshalJSON(info.AppMetadata))),
		Email:       stripe.String(info.Email),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	cache.Working.Customers[c.Email] = c
	return &api.Identifier{
		Id: c.ID,
	}, nil
}

type User struct {
	driver.PluginFunc
}

func (u *User) GetUser(ctx context.Context, r *api.UserByEmailRequest) (*api.User, error) {
	usr := cache.Working.Users[r.Email]
	if usr == nil {
		return nil, errors.New("user not found in cache")
	}
	return usr, nil
}

func (u *User) UpdateUser(context.Context, *api.UserRequest) (*api.Bytes, error) {
	panic("implement me")
}

func (u *User) CreateUser(ctx context.Context, r *api.UserRequest) (*api.Bytes, error) {
	panic("implement me")
}

func (u *User) DeleteUser(context.Context, *api.UserByEmailRequest) (*api.Bytes, error) {
	panic("implement me")
}

func (u *User) ListUsers(t *api.ManagementToken, stream api.UserService_ListUsersServer) error {
	users := []*api.User{}

	for _, user := range users {
		err := stream.Send(user)
		if err != nil {
			return err
		}
	}
	return nil
}
