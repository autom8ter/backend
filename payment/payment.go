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

func (s *Payment) Subscribe(ctx context.Context, request *api.SubscribeRequest) (*api.SubscriptionResponse, error) {
	cust := cache.Working.Customers[request.Email.Text]
	// create a subscription
	resp, err := sub.New(&stripe.SubscriptionParams{
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
	var p api.Plan
	switch {
	case resp.Plan.Nickname == "free", resp.Plan.Nickname == "Free", resp.Plan.Nickname == "FREE":
		p = api.Plan_FREE
	case resp.Plan.Nickname == "basic", resp.Plan.Nickname == "Basic", resp.Plan.Nickname == "BASIC":
		p = api.Plan_BASIC
	case resp.Plan.Nickname == "premium", resp.Plan.Nickname == "Premium", resp.Plan.Nickname == "PREMIUM":
		p = api.Plan_PREMIUM
	default:
		p = api.Plan_FREE
	}
	a := &api.SubscriptionResponse{
		Id:           &common.Identifier{Id: common.ToString(resp.ID)},
		Amount:       common.ToInt64(int(resp.Plan.Amount)),
		DaysUntilDue: common.ToInt64(int(resp.DaysUntilDue)),
		Annotations: common.ToStringMap(map[string]string{
			"live_mode": fmt.Sprintf("%v", resp.Livemode),
		}),
		Plan:   p,
		User:   cache.Working.Users[request.Email.Text],
		Status: common.ToString(string(resp.Status)),
	}
	for k, v := range resp.Metadata {
		a.Annotations.Put(k, common.ToString(v))
	}

	return a, nil
}

func (s *Payment) Unsubscribe(ctx context.Context, request *api.UnSubscribeRequest) (*api.SubscriptionResponse, error) {
	cust := cache.Working.Customers[request.Email.Text]
	for _, s := range cust.Subscriptions.Data {
		if s.Plan.Nickname == request.Plan.Normalize().Text {
			resp, err := sub.Cancel(s.ID, nil)
			if err != nil {
				return nil, err
			}
			var p api.Plan
			switch {
			case resp.Plan.Nickname == "free", resp.Plan.Nickname == "Free", resp.Plan.Nickname == "FREE":
				p = api.Plan_FREE
			case resp.Plan.Nickname == "basic", resp.Plan.Nickname == "Basic", resp.Plan.Nickname == "BASIC":
				p = api.Plan_BASIC
			case resp.Plan.Nickname == "premium", resp.Plan.Nickname == "Premium", resp.Plan.Nickname == "PREMIUM":
				p = api.Plan_PREMIUM
			default:
				p = api.Plan_FREE
			}

			a := &api.SubscriptionResponse{
				Id:           &common.Identifier{Id: common.ToString(resp.ID)},
				Amount:       common.ToInt64(int(resp.Plan.Amount)),
				DaysUntilDue: common.ToInt64(int(resp.DaysUntilDue)),
				Annotations: common.ToStringMap(map[string]string{
					"live_mode": fmt.Sprintf("%v", resp.Livemode),
				}),
				Plan:   p,
				User:   cache.Working.Users[request.Email.Text],
				Status: common.ToString(string(resp.Status)),
			}
			for k, v := range resp.Metadata {
				a.Annotations.Put(k, common.ToString(v))
			}

			return a, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("plan: %s not found for customer: %s", request.Plan, request.Email))
}
