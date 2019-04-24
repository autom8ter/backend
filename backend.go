package backend

import (
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/engine"
)

type Backend struct {
	Plugins []driver.PluginFunc
}

func (b *Backend) AddPlugin(fn driver.PluginFunc) {
	b.Plugins=append(b.Plugins, fn)
}


func (b *Backend) Serve(addr string, debug bool) error {
	return engine.Serve(addr, debug, b.asPlugins()...)
}

func (b *Backend) asPlugins() []driver.Plugin {
	plugs := []driver.Plugin{}
	for _, p := range b.Plugins {
		plugs = append(plugs, p)
	}
	return plugs
}
