package webpay

import (
	"net/url"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
)

type Customer struct {
	webpayclient *WebPayClient
	path         string
}

func NewCustomer(cli *WebPayClient) Customer {
	ret := Customer{
		webpayclient: cli,
		path:         "customers",
	}

	return ret
}

// Create creates new Customer.
func (c Customer) Create(card Card) (*simplejson.Json, error) {
	params := card.AddParams(url.Values{})

	return c.webpayclient.Post(c.path, params)
}

// Delete deltes a Customer.
func (c Customer) Delete(customerId string) (*simplejson.Json, error) {
	path := strings.Join([]string{c.path, customerId}, "/")

	return c.webpayclient.Delete(path, url.Values{})
}

// All returnes customer list filtered by params.
func (c Customer) All(args map[string]int) (*simplejson.Json, error) {
	path := getAllPathWithQuery(c.path, args)
	return c.webpayclient.Get(path, url.Values{})
}

// CreateByToken creates Customer from the token.
func (c Customer) CreateByToken(token, email, description string) (*simplejson.Json, error) {
	params := url.Values{}
	params.Add("card", token)
	params.Add("email", email)
	params.Add("description", description)

	return c.webpayclient.Post(c.path, params)
}
