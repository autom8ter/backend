package cache

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/plan"
	"github.com/stripe/stripe-go/product"
	"time"
)

type Cache struct {
	Users         map[string]interface{}
	Customers     map[string]*stripe.Customer
	Plans         map[string]*stripe.Plan
	Products      map[string]*stripe.Product
	Charges       map[string]*stripe.Charge
	Subscriptions map[string]*stripe.Subscription
}

func NewCache() *Cache {
	return &Cache{
		Users:     make(map[string]interface{}),
		Customers: make(map[string]*stripe.Customer),
		Plans:     make(map[string]*stripe.Plan),
		Products:  make(map[string]*stripe.Product),
		Charges:   make(map[string]*stripe.Charge),
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
}

func (cache *Cache) Loop(duration time.Duration) {
	for {
		cache.Sync()
		time.Sleep(duration)
	}
}

func (cache *Cache) AddStripeCustomer(c *stripe.Customer) {
	cache.Customers[c.Email] = c
}

func (cache *Cache) RemoveStripeCustomer(email string) {
	cache.Customers[email] = nil
}

func (cache *Cache) GetStripeCustomer(email string) *stripe.Customer {
	return cache.Customers[email]
}

func (cache *Cache) TotalStripeCustomers() int {
	return len(cache.Customers)
}

func (cache *Cache) TotalStripePlans() int {
	return len(cache.Plans)
}

func (cache *Cache) TotalStripeCharges() int {
	return len(cache.Charges)
}

func (cache *Cache) TotalStripeProducts() int {
	return len(cache.Products)
}

func (cache *Cache) TotalUsers() int {
	return len(cache.Users)
}

func (cache *Cache) TotalSubscriptions() int {
	return len(cache.Subscriptions)
}

var Working = NewCache()

var SYNC_FREQUENCY = 1 * time.Minute

func Init() {
	go Working.Loop(SYNC_FREQUENCY)
}
