package backend

import (
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/driver"
)

type Backend struct {
	plugins []driver.PluginFunc
}

func NewBackend(plugs ...driver.PluginFunc) *Backend {
	return &Backend{
		plugs,
	}
}

func (b *Backend) Serve(addr string, debug bool) error {
	return engine.Serve(addr, debug, b.asPlugins()...)
}

func (b *Backend) asPlugins() []driver.Plugin {
	plugs := []driver.Plugin{}
	for _, v := range b.plugins {
		plugs = append(plugs, v)
	}
	return plugs
}
