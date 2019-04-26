package payment

import (
	"context"
	"fmt"
	"github.com/autom8ter/api"
	"github.com/autom8ter/backend/cache"
	"github.com/autom8ter/engine/driver"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/sub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Payment struct {
	driver.PluginFunc
}

func NewPayment() *Payment {
	p := &Payment{}
	p.PluginFunc = func(s *grpc.Server) {
		api.RegisterPaymentServiceServer(s, p)
	}
	return p
}

func (*Payment) PurchasePhoneNumber(c context.Context, i *api.PhoneNumber) (*api.PhoneNumberResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "service not yet available")
}

func (s *Payment) Subscribe(ctx context.Context, request *api.SubscribeRequest) (*api.Bytes, error) {
	cust := cache.Working.Customers[request.Email]
	// create a subscription
	subs, err := sub.New(&stripe.SubscriptionParams{
		Customer: stripe.String(cust.ID),
		Plan:     stripe.String(request.Plan),
		Card: &stripe.CardParams{
			Number:   stripe.String(request.Card.Number),
			ExpMonth: stripe.String(request.Card.ExpMonth),
			ExpYear:  stripe.String(request.Card.ExpYear),
			CVC:      stripe.String(request.Card.Cvc),
		},
	})
	if err != nil {
		return nil, err
	}
	return api.AsBytes(subs), nil
}

func (s *Payment) Unsubscribe(ctx context.Context, request *api.UnSubscribeRequest) (*api.Bytes, error) {
	cust := cache.Working.Customers[request.Email]
	for _, s := range cust.Subscriptions.Data {
		if s.Plan.Nickname == request.Plan {
			s, err := sub.Cancel(s.ID, nil)
			if err != nil {
				return nil, err
			}
			return api.AsBytes(s), nil
		}
	}
	return nil, api.Util.NewError(fmt.Sprintf("plan: %s not found for customer: %s", request.Plan, request.Email))
}
