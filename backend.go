package main

import (
	"context"
	"github.com/autom8ter/api"
	"github.com/autom8ter/api/common"
	"github.com/autom8ter/fire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
)

var firedb *fire.Client
var addr string

func main() {
	ctx := context.TODO()
	projectID := os.Getenv("GCLOUD_PROJECT")
	if projectID == "" {
		log.Fatalf("Set Firebase project ID via GCLOUD_PROJECT env variable.")
	}

	firedb, err := fire.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Cannot create client: %v", err)
	}
	defer firedb.Fire.Close()
	var db = &api.Database{
		GetUserFunc: func(ctx context.Context, email *common.Identifier) (user *api.User, e error) {
			return nil, status.Error(codes.Unimplemented, "unimplemented")

		},
		DeleteUserFunc: func(ctx context.Context, email *common.Identifier) (empty *common.Empty, e error) {
			return nil, status.Error(codes.Unimplemented, "unimplemented")

		},
		UpdateUserFunc: func(ctx context.Context, u *api.UpdateUserRequest) (user *api.User, e error) {
			return nil, status.Error(codes.Unimplemented, "unimplemented")

		},
		CreateUserFunc: func(ctx context.Context, u *api.User) (user *api.User, e error) {
			return nil, status.Error(codes.Unimplemented, "unimplemented")
		},
		ListUsersFunc: func(e *common.Empty, stream api.DBService_ListUsersServer) error {
			return nil
		},
	}
	db.PluginFunc = func(s *grpc.Server) {
		api.RegisterDBServiceServer(s, db)
	}
	var deb = &api.Debugger{
		DebugFunc: func(ctx context.Context, s *common.String) (i *common.String, e error) {
			return s, nil
		},
	}
	deb.PluginFunc = func(s *grpc.Server) {
		api.RegisterDebugServiceServer(s, deb)
	}
	if err := api.Serve(addr, db, deb); err != nil {
		if addr == "" {
			addr = ":3000"
		}
		log.Fatalln(err.Error())
	}
}
