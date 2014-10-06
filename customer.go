package webpay

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/bitly/go-simplejson"
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

// Create creates new Charge.
func (c Customer) Create(card Card) (*simplejson.Json, error) {
	params := card.AddParams(url.Values{})

	return c.webpayclient.Post(c.path, params)
}

// All returnes customer list filtered by params.
func (c Customer) All(count int, offset int, created map[string]int) (*simplejson.Json, error) {
	query := url.Values{}
	if count > 0 {
		query.Add("count", strconv.Itoa(count))
	}
	if offset > 0 {
		query.Add("offset", strconv.Itoa(offset))
	}
	/*
	    // TODO: How to set dict
		if len(created) > 0 {
			for c, t := range created {
				query.Add(c, t)
			}
		}
	*/

	if offset <= 0 {
		offset = 0
	}
	path := fmt.Sprintf("%s?%s", c.path, query.Encode())

	return c.webpayclient.Get(path, url.Values{})
}
