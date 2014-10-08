package webpay

import (
	"net/url"
	"strings"

	"github.com/bitly/go-simplejson"
)

type Event struct {
	webpayclient *WebPayClient
	path         string
}

func NewEvent(cli *WebPayClient) Event {
	ret := Event{
		webpayclient: cli,
		path:         "events",
	}

	return ret
}

// Retrieve retrieves Event information from WebPay.
func (c Event) Retrieve(chid string) (*simplejson.Json, error) {
	path := strings.Join([]string{c.path, chid}, "/")

	return c.webpayclient.Get(path, url.Values{})
}

// All returnes customer list filtered by params.
func (c Event) All(args map[string]int) (*simplejson.Json, error) {
	path := getAllPathWithQuery(c.path, args)
	return c.webpayclient.Get(path, url.Values{})
}
