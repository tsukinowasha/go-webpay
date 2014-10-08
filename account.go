package webpay

import (
	"net/url"

	"github.com/bitly/go-simplejson"
)

type Account struct {
	webpayclient *WebPayClient
	path         string
}

func NewAccount(cli *WebPayClient) Account {
	ret := Account{
		webpayclient: cli,
		path:         "account",
	}

	return ret
}

func (c Account) Retrieve() (*simplejson.Json, error) {
	return c.webpayclient.Get(c.path, url.Values{})
}
