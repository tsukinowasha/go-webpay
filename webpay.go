package webpay

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
)

type WebPayClient struct {
	header     map[string]string
	options    map[string]string
	auth_token string
	Charge     Charge
	Customer   Customer
	Recursion  Recursion
	Token      Token
}

func NewWebPayClient(auth_token string) WebPayClient {
	cli := WebPayClient{
		auth_token: auth_token,
		header: map[string]string{
			"Content-Type":    "application/x-www-form-urlencoded",
			"Accept":          "application/json",
			"User-Agent":      "webpay/2.1.1 golang",
			"Accept-Language": "en",
			"Authorization":   "Bearer " + auth_token,
		},
		options: map[string]string{
			"api_base": "https://api.webpay.jp/v1",
		},
	}

	cli.Charge = NewCharge(&cli)
	cli.Customer = NewCustomer(&cli)
	cli.Recursion = NewRecursion(&cli)
	cli.Token = NewToken(&cli)

	return cli
}

func (cli WebPayClient) SetAcceptLanguage(lang string) {
	cli.header["Accept-Language"] = lang
}

func (cli WebPayClient) Post(path string, params url.Values) (*simplejson.Json, error) {
	return cli.Request("POST", path, params)
}
func (cli WebPayClient) Get(path string, params url.Values) (*simplejson.Json, error) {
	return cli.Request("GET", path, params)
}

func (cli WebPayClient) Request(method, path string, params url.Values) (*simplejson.Json, error) {
	client := &http.Client{}

	u := cli.options["api_base"] + "/" + path
	req, _ := http.NewRequest(
		method,
		u,
		strings.NewReader(params.Encode()),
	)
	for k, v := range cli.header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	js, err := simplejson.NewJson(body)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// GetId returns 'id' of the JSON.
func GetId(json *simplejson.Json) (string, error) {
	return json.Get("id").String()
}
