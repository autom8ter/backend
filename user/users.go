package user

import (
	"context"
	"github.com/autom8ter/api"
	"github.com/autom8ter/api/common"
	"github.com/autom8ter/engine/driver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	driver.PluginFunc
}

func NewUser() *User {
	u := &User{}
	u.PluginFunc = func(s *grpc.Server) {
		api.RegisterUserServiceServer(s, u)
	}
	return u
}

func (User) QueryUsers(q *api.TokenQuery, stream api.UserService_QueryUsersServer) error {
	return status.Error(codes.Unimplemented, "api is not yet implemented")
}

func (User) CreateUser(ctx context.Context, bits *common.Bytes) (*api.User, error) {
	return nil, status.Error(codes.Unimplemented, "api is not yet implemented")
}

func (User) GetUser(ctx context.Context, id *common.Identifier) (*api.User, error) {
	return nil, status.Error(codes.Unimplemented, "api is not yet implemented")
}

func (User) DeleteUser(ctx context.Context, id *common.Identifier) (*api.User, error) {
	return nil, status.Error(codes.Unimplemented, "api is not yet implemented")
}

func (User) UpdateUser(ctx context.Context, id *api.IDBody) (*api.User, error) {
	return nil, status.Error(codes.Unimplemented, "api is not yet implemented")
}

func (User) UserRoles(id *common.Identifier, stream api.UserService_UserRolesServer) error {
	return status.Error(codes.Unimplemented, "api is not yet implemented")
}
