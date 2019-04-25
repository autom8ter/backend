package utility

import (
	"context"
	"github.com/autom8ter/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"html/template"
)

type Renderer struct{}

func (r *Renderer) Render(ctx context.Context, tmpl *api.Template) (*api.Bytes, error) {
	t, err := template.New("").Funcs(api.FuncMap()).Parse(tmpl.Text)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to render template: %s", err.Error())
	}
	bits := &api.Bytes{
		Bits: []byte{},
	}
	err = api.Util.RenderHTML(t, tmpl.Data, bits)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to render template: %s", err.Error())
	}
	return bits, nil
}
