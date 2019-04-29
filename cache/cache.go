package cache

import (
	"github.com/autom8ter/api"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/plan"
	"github.com/stripe/stripe-go/product"
	"github.com/stripe/stripe-go/sub"
	"time"
)

type Cache struct {
	Looping       bool
	SyncFrequency time.Duration
	Users         map[string]*api.User
	Customers     map[string]*stripe.Customer
	Plans         map[string]*stripe.Plan
	Products      map[string]*stripe.Product
	Charges       map[string]*stripe.Charge
	Subscriptions map[string]*stripe.Subscription
}

func DefaultCache() *Cache {
	return &Cache{
		Looping:       false,
		SyncFrequency: DEFAULT_SYNC_FREQUENCY,
		Users:         make(map[string]*api.User),
		Customers:     make(map[string]*stripe.Customer),
		Plans:         make(map[string]*stripe.Plan),
		Products:      make(map[string]*stripe.Product),
		Charges:       make(map[string]*stripe.Charge),
		Subscriptions: make(map[string]*stripe.Subscription),
	}
}

func (cache *Cache) Sync() {
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
	subs := sub.List(nil)
	for subs.Next() {
		s := subs.Subscription()
		cache.Subscriptions[s.ID] = s
	}
}

func (cache *Cache) Loop() {
	cache.Looping = true
	for cache.Looping {
		cache.Sync()
		time.Sleep(cache.SyncFrequency)
	}
	cache.Looping = false
}

var Working = DefaultCache()

var DEFAULT_SYNC_FREQUENCY = 1 * time.Minute

func Init(frequency time.Duration) {
	if frequency != 0 {
		Working.SyncFrequency = frequency
	}
	go Working.Loop()
}
