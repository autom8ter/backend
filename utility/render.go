package utility

import (
	"context"
	"github.com/autom8ter/api"
)

type Renderer struct{}

func (r *Renderer) Render(ctx context.Context, tmpl *api.RenderRequest) (*api.Bytes, error) {
	b := api.NewBytes()
	err := tmpl.Template.RenderBytes(b, tmpl.Data)
	if err != nil {
		return nil, err
	}
	return b, nil
}
