package webpay

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
)

type WebPayClient struct {
	header    map[string]string
	options   map[string]string
	authToken string
	Charge    Charge
	Customer  Customer
	Recursion Recursion
	Token     Token
	Event     Event
	Account   Account
	mode      string
}

func NewWebPayClient(auth_token string) WebPayClient {
	cli := WebPayClient{
		authToken: auth_token,
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
	cli.Event = NewEvent(&cli)
	cli.Account = NewAccount(&cli)
	cli.mode = "real"

	return cli
}

func (cli WebPayClient) SetAcceptLanguage(lang string) {
	cli.header["Accept-Language"] = lang
}

func (cli WebPayClient) Post(path string, params url.Values) (*simplejson.Json, error) {
	return cli.Request("POST", path, params)
}
func (cli WebPayClient) Delete(path string, params url.Values) (*simplejson.Json, error) {
	return cli.Request("DELETE", path, params)
}
func (cli WebPayClient) Get(path string, params url.Values) (*simplejson.Json, error) {
	return cli.Request("GET", path, params)
}

func (cli WebPayClient) Request(method, path string, params url.Values) (*simplejson.Json, error) {
	if cli.mode != "real" {
		return cli.returnMockJson(method, path, params)
	}

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
	req.SetBasicAuth(cli.authToken, "")
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

// getAllPathWithQuery creates path with query from map args.
func getAllPathWithQuery(path string, args map[string]int) string {
	query := url.Values{}
	if count, exists := args["count"]; exists {
		query.Add("count", strconv.Itoa(count))
	}
	if offset, exists := args["offset"]; exists {
		query.Add("offset", strconv.Itoa(offset))
	}

	if created, exists := args["created"]; exists {
		query.Add("created", strconv.Itoa(created))
	}
	if gt, exists := args["gt"]; exists {
		query.Add("gt", strconv.Itoa(gt))
	}
	if gte, exists := args["gte"]; exists {
		query.Add("gte", strconv.Itoa(gte))
	}
	if lt, exists := args["lt"]; exists {
		query.Add("lt", strconv.Itoa(lt))
	}
	if lte, exists := args["lte"]; exists {
		query.Add("lte", strconv.Itoa(lte))
	}

	return fmt.Sprintf("%s?%s", path, query.Encode())
}

// returnMockJson return JSON which is read from under test_data dir for testing.
func (cli WebPayClient) returnMockJson(method, path string, params url.Values) (*simplejson.Json, error) {
	path = strings.Replace(path, "?", "_", -1)
	path = strings.Replace(path, "=", "_", -1)
	path = strings.Replace(path, "&", "_", -1)
	path = strings.Replace(path, "/", "_", -1)

	if method == "DELETE" {
		path = "delete_" + path
	}

	// 	fmt.Println(path)

	dat, err := ioutil.ReadFile("test_data/" + path)
	if err != nil {
		js, _ := simplejson.NewJson([]byte("{}"))
		return js, nil
	}

	js, err := simplejson.NewJson(dat)
	if err != nil {
		return nil, err
	}

	return js, nil
}
