package utility

import (
	"context"
	"github.com/autom8ter/api"
	"github.com/autom8ter/backend/clients"
)

func NewMarshaler() *Marshaler {
	m := &Marshaler{}

	return m
}

type Marshaler struct{}

func (m *Marshaler) MarshalJSON(ctx context.Context, bytes *api.Bytes) (*api.Bytes, error) {
	return &api.Bytes{
		Bits: clients.Util.MarshalJSON(bytes.Bits),
	}, nil
}

func (m *Marshaler) MarshalYAML(ctx context.Context, bytes *api.Bytes) (*api.Bytes, error) {
	return &api.Bytes{
		Bits: clients.Util.MarshalYAML(bytes.Bits),
	}, nil
}

func (m *Marshaler) MarshalXML(ctx context.Context, bytes *api.Bytes) (*api.Bytes, error) {
	return &api.Bytes{
		Bits: clients.Util.MarshalXML(bytes.Bits),
	}, nil
}
