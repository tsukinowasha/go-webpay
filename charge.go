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

func (c Charge) Create(amount float64, currency string, card Card) (*simplejson.Json, error) {
	params := url.Values{}
	params.Add("amount", strconv.FormatFloat(amount, 'f', -1, 64))
	params.Add("currency", currency)

	params = card.AddParams(params)

	return c.webpayclient.Post(c.path, params)
}

func (c Charge) Retrieve(chid string) (*simplejson.Json, error) {
	path := strings.Join([]string{c.path, chid}, "/")

	return c.webpayclient.Get(path, url.Values{})
}
