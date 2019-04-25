package utility

import (
	"context"
	"github.com/autom8ter/api"
)

func NewMarshaler() *Marshaler {
	m := &Marshaler{}

	return m
}

type Marshaler struct{}

func (m *Marshaler) MarshalJSON(ctx context.Context, bytes *api.Bytes) (*api.Bytes, error) {
	return &api.Bytes{
		Bits: api.Util.MarshalJSON(bytes.Bits),
	}, nil
}

func (m *Marshaler) MarshalYAML(ctx context.Context, bytes *api.Bytes) (*api.Bytes, error) {
	return &api.Bytes{
		Bits: api.Util.MarshalYAML(bytes.Bits),
	}, nil
}

func (m *Marshaler) MarshalXML(ctx context.Context, bytes *api.Bytes) (*api.Bytes, error) {
	return &api.Bytes{
		Bits: api.Util.MarshalXML(bytes.Bits),
	}, nil
}
