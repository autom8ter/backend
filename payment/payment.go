package payment

import (
	"context"
	"fmt"
	"github.com/autom8ter/api"
	"github.com/autom8ter/api/common"
	"github.com/autom8ter/backend/cache"
	"github.com/autom8ter/engine/driver"
	"github.com/pkg/errors"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/sub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Payment struct {
	driver.PluginFunc
}

func (s *Payment) SearchPhoneNumber(r *api.SearchPhoneNumberRequest, stream api.PaymentService_SearchPhoneNumberServer) error {
	return status.Error(codes.Unimplemented, "api is not yet implemented")
}

func NewPayment() *Payment {
	p := &Payment{}
	p.PluginFunc = func(s *grpc.Server) {
		api.RegisterPaymentServiceServer(s, p)
	}
	return p
}

func (*Payment) PurchasePhoneNumber(c context.Context, i *api.PhoneNumber) (*api.PhoneNumberResource, error) {
	return nil, status.Error(codes.Unimplemented, "api is not yet implemented")
}

func (s *Payment) Subscribe(ctx context.Context, request *api.SubscribeRequest) (*common.Bytes, error) {
	cust := cache.Working.Customers[request.Email.Text]
	// create a subscription
	subs, err := sub.New(&stripe.SubscriptionParams{
		Customer: common.ToString(cust.ID).Pointer(),
		Plan:     request.Plan.Normalize().Pointer(),
		Card: &stripe.CardParams{
			Number:   request.Card.Number.Pointer(),
			ExpMonth: request.Card.ExpMonth.Pointer(),
			ExpYear:  request.Card.ExpYear.Pointer(),
			CVC:      request.Card.Cvc.Pointer(),
		},
	})
	if err != nil {
		return nil, err
	}
	return common.AsBytes(subs), nil
}

func (s *Payment) Unsubscribe(ctx context.Context, request *api.UnSubscribeRequest) (*common.Bytes, error) {
	cust := cache.Working.Customers[request.Email.Text]
	for _, s := range cust.Subscriptions.Data {
		if s.Plan.Nickname == request.Plan.Normalize().Text {
			s, err := sub.Cancel(s.ID, nil)
			if err != nil {
				return nil, err
			}
			return common.AsBytes(s), nil
		}
	}
	return nil, errors.New(fmt.Sprintf("plan: %s not found for customer: %s", request.Plan, request.Email))
}
