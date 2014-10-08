package webpay

import (
	"net/url"
	"strings"

	"github.com/bitly/go-simplejson"
)

type Token struct {
	webpayclient *WebPayClient
	path         string
}

func NewToken(cli *WebPayClient) Token {
	ret := Token{
		webpayclient: cli,
		path:         "tokens",
	}

	return ret
}

func (c Token) Create(uuid string, card Card) (*simplejson.Json, error) {
	params := url.Values{}
	if uuid != "" {
		params.Add("uuid", uuid)
	}
	params = card.AddParams(params)

	return c.webpayclient.Post(c.path, params)
}

func (c Token) Retrieve(chid string) (*simplejson.Json, error) {
	path := strings.Join([]string{c.path, chid}, "/")

	return c.webpayclient.Get(path, url.Values{})
}
