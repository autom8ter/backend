package backend_test

import (
	"context"
	"fmt"
	"github.com/autom8ter/api"
	"github.com/autom8ter/api/common"
	"google.golang.org/grpc"
	"os"
	"testing"
)

func Test(t *testing.T) {
	conn, err := grpc.DialContext(context.TODO(), "localhost:3000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err.Error())
	}
	client := api.NewClientSet(conn)
	resp, err := client.Utility.Echo(context.TODO(), &common.String{
		Text: "Hello, " + os.Getenv("USER"),
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("%s\n", resp)

}
