package utility

import (
	"github.com/autom8ter/api"
	"github.com/autom8ter/backend/config"
	"github.com/autom8ter/engine/driver"
	"google.golang.org/grpc"
)

type Utility struct {
	*Echoer
	driver.PluginFunc
}

func NewUtility(c *config.Config) *Utility {
	u := &Utility{
		Echoer: NewEchoer(c),
	}
	u.PluginFunc = func(s *grpc.Server) {
		api.RegisterUtilityServiceServer(s, u)
	}
	return u
}
