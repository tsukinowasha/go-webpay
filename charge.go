package webpay

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
)

type Charge struct {
	webpayclient *WebPayClient
	path         string
}

func NewCharge(cli *WebPayClient) Charge {
	ret := Charge{
		webpayclient: cli,
		path:         "charges",
	}

	return ret
}

// CreateByCustomerId creates Charge.
func (c Charge) Create(amount float64, currency string, card Card) (*simplejson.Json, error) {
	params := url.Values{}
	params.Add("amount", strconv.FormatFloat(amount, 'f', -1, 64))
	params.Add("currency", currency)

	params = card.AddParams(params)

	return c.webpayclient.Post(c.path, params)
}

// CreateByCustomer creates Charge from specified Customer.
func (c Charge) CreateByCustomer(amount float64, currency string, customer string) (*simplejson.Json, error) {
	params := url.Values{}
	params.Add("amount", strconv.FormatFloat(amount, 'f', -1, 64))
	params.Add("currency", currency)
	params.Add("customer", customer)

	return c.webpayclient.Post(c.path, params)
}

// Retrieve retrieves Charge information from WebPay.
func (c Charge) Retrieve(chid string) (*simplejson.Json, error) {
	path := strings.Join([]string{c.path, chid}, "/")

	return c.webpayclient.Get(path, url.Values{})
}

// All returnes customer list filtered by params.
func (c Charge) All(args map[string]int) (*simplejson.Json, error) {
	path := getAllPathWithQuery(c.path, args)
	return c.webpayclient.Get(path, url.Values{})
}
