package user

import (
	"context"
	"fmt"
	"github.com/autom8ter/api"
	"github.com/autom8ter/engine/driver"
	"github.com/pkg/errors"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/plan"
	"github.com/stripe/stripe-go/product"
	"github.com/stripe/stripe-go/sub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"time"
)

func NewUserCache() *UserCache {
	c := &UserCache{
		Customers: make(map[string]*stripe.Customer),
		Plans:     make(map[string]*stripe.Plan),
		Products:  make(map[string]*stripe.Product),
		Charges:   make(map[string]*stripe.Charge),
	}
	c.PluginFunc = func(s *grpc.Server) {
		api.RegisterUserServiceServer(s, c)
	}
	return c
}

func (cache *UserCache) Sync() {
	customers := customer.List(nil)
	for customers.Next() {
		c := customers.Customer()
		cache.Customers[c.Email] = c
	}
	plans := plan.List(nil)
	for plans.Next() {
		p := plans.Plan()
		cache.Plans[p.Nickname] = p
	}
	products := product.List(nil)
	for products.Next() {
		p := products.Product()
		cache.Products[p.Name] = p
	}
	charges := charge.List(nil)
	for charges.Next() {
		c := charges.Charge()
		cache.Charges[c.ID] = c
	}
}

func (cache *UserCache) Loop(duration time.Duration) {
	for {
		cache.Sync()
		time.Sleep(duration)
	}
}

func (cache *UserCache) addStripeCustomer(c *stripe.Customer) {
	cache.Customers[c.Email] = c
}

func (cache *UserCache) removeStripeCustomer(email string) {
	cache.Customers[email] = nil
}

func (cache *UserCache) getStripeCustomer(email string) *stripe.Customer {
	return cache.Customers[email]
}

func (cache *UserCache) TotalStripeCustomers() int {
	return len(cache.Customers)
}

func (cache *UserCache) TotalStripePlans() int {
	return len(cache.Plans)
}

func (cache *UserCache) TotalStripeCharges() int {
	return len(cache.Charges)
}

func (cache *UserCache) TotalStripeProducts() int {
	return len(cache.Products)
}

func (cache *UserCache) subscribeStripeCustomer(ctx context.Context, email, plan, cardnum, expmonth, expyear, cvc string) (*api.Identifier, error) {
	cust := cache.Customers[email]
	// create a subscription
	s, err := sub.New(&stripe.SubscriptionParams{
		Customer: stripe.String(cust.ID),
		Plan:     stripe.String(plan),
		Card: &stripe.CardParams{
			Number:   stripe.String(cardnum),
			ExpMonth: stripe.String(expmonth),
			ExpYear:  stripe.String(expyear),
			CVC:      stripe.String(cvc),
		},
	})
	if err != nil {
		return nil, err
	}
	return &api.Identifier{
		Id: s.ID,
	}, nil
}

func (cache *UserCache) unsubscribeStripeCustomer(ctx context.Context, email, plan string) (*api.Identifier, error) {
	cust := cache.Customers[email]
	for _, s := range cust.Subscriptions.Data {
		if s.Plan.Nickname == plan {
			s, err := sub.Cancel(s.ID, nil)
			if err != nil {
				return nil, err
			}
			return &api.Identifier{
				Id: s.ID,
			}, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("plan: %s not found for customer: %s", plan, email))
}

func (cache *UserCache) createStripeCustomer(ctx context.Context, info api.UserInfo) (*api.Identifier, error) {
	c, err := customer.New(&stripe.CustomerParams{
		Description: stripe.String(string(api.Util.MarshalJSON(info.AppMetadata))),
		Email:       stripe.String(info.Email),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	cache.addStripeCustomer(c)
	return &api.Identifier{
		Id: c.ID,
	}, nil
}

type UserCache struct {
	Customers     map[string]*stripe.Customer
	Plans         map[string]*stripe.Plan
	Products      map[string]*stripe.Product
	Charges       map[string]*stripe.Charge
	Subscriptions map[string]*stripe.Subscription
	driver.PluginFunc
}

func (cache *UserCache) Subscribe(ctx context.Context, request *api.SubscribeRequest) (*api.Identifier, error) {
	return cache.subscribeStripeCustomer(ctx, request.Email, request.Plan, request.Card.Number, request.Card.ExpMonth, request.Card.ExpYear, request.Card.Cvc)
}

func (cache *UserCache) Unsubscribe(ctx context.Context, r *api.UnSubscribeRequest) (*api.Identifier, error) {
	return cache.unsubscribeStripeCustomer(ctx, r.Email, r.Plan)
}
