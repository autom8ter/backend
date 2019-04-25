package contact

import (
	"github.com/autom8ter/api"
	"github.com/autom8ter/backend/config"
	"github.com/autom8ter/engine/driver"
	"google.golang.org/grpc"
)

type Contact struct {
	*SMSer
	*Caller
	*Emailer
	driver.PluginFunc
}

func NewConatact(cfg *config.Config) *Contact {
	c := &Contact{
		SMSer:   NewSMSer(cfg),
		Caller:  NewCaller(cfg),
		Emailer: NewEmailer(cfg),
	}
	c.PluginFunc = func(s *grpc.Server) {
		api.RegisterContactServiceServer(s, c)
	}
	return c
}
