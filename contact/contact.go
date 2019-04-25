package contact

import (
	"github.com/autom8ter/api"
	"github.com/autom8ter/engine/driver"
	"google.golang.org/grpc"
)

type Contact struct {
	*SMSer
	*Caller
	*Emailer
	driver.PluginFunc
}

func NewConatact() *Contact {
	c := &Contact{
		SMSer:   &SMSer{},
		Caller:  &Caller{},
		Emailer: &Emailer{},
	}
	c.PluginFunc = func(s *grpc.Server) {
		api.RegisterContactServiceServer(s, c)
	}
	return c
}
