package utility

import (
	"github.com/autom8ter/api"
	"github.com/autom8ter/engine/driver"
	"google.golang.org/grpc"
)

type Utility struct {
	*Echoer
	*Marshaler
	*Renderer
	driver.PluginFunc
}

func NewUtility() *Utility {
	u := &Utility{
		Echoer:    NewEchoer(),
		Marshaler: NewMarshaler(),
	}
	u.PluginFunc = func(s *grpc.Server) {
		api.RegisterUtilityServiceServer(s, u)
	}
	return u
}
