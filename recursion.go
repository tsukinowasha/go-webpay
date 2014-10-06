package webpay

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
)

type Recursion struct {
	webpayclient *WebPayClient
	path         string
}

func NewRecursion(cli *WebPayClient) Recursion {
	ret := Recursion{
		webpayclient: cli,
		path:         "recursions",
	}

	return ret
}

func (c Recursion) Create(amount float64, currency, customer, period, description string) (*simplejson.Json, error) {
	params := url.Values{}
	params.Add("amount", strconv.FormatFloat(amount, 'f', -1, 64))
	params.Add("currency", currency)
	params.Add("customer", customer)
	params.Add("period", period)
	params.Add("description", description)

	return c.webpayclient.Post(c.path, params)
}

func (c Recursion) Retrieve(recid string) (*simplejson.Json, error) {
	path := strings.Join([]string{c.path, recid}, "/")

	return c.webpayclient.Get(path, url.Values{})
}
