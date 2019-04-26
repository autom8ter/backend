package contact

import (
	"github.com/autom8ter/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Searcher struct {
}

func (Searcher) SearchPhoneNumber(r *api.SearchPhoneNumberRequest, stream api.ContactService_SearchPhoneNumberServer) error {
	return status.Errorf(codes.Unimplemented, "service not yet available")
}
