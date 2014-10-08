package webpay

// If you want to test with real WebPay, change this to "real".
var TestMode = "mock"

var TestCard = Card{
	number:    "4242-4242-4242-4242",
	exp_month: 11,
	exp_year:  2014,
	cvc:       "123",
	name:      "KEI KUBO",
}

var TestAuthToken = "test_secret_eHn4TTgsGguBcW764a2KA8Yd"
var TestCharge = "ch_3u32gEgiy2xjfEf"
var TestCustomer = "cus_45d3MV5phaXJ4uv"
var TestRecursion = "rec_8hla3E10n2T84Ne"
var TestToken = "tok_3ybc93ckR01qeKx"
var TestEvent = "evt_cFRc4s9FD7BE1Ox"

func NewWebPayClientForTesting(mode string, auth_token string) WebPayClient {
	if mode == "real" {
		return NewWebPayClient(auth_token)
	} else {
		return NewWebPayClientForTest(auth_token)
	}
}

func NewWebPayClientForTest(auth_token string) WebPayClient {
	cli := WebPayClient{
		authToken: auth_token,
		header:    map[string]string{},
		options: map[string]string{
			"api_base": "https://api.webpay.jp/v1",
		},
		mode: "mock",
	}

	cli.Charge = NewCharge(&cli)
	cli.Customer = NewCustomer(&cli)
	cli.Recursion = NewRecursion(&cli)
	cli.Token = NewToken(&cli)
	cli.Event = NewEvent(&cli)
	cli.Account = NewAccount(&cli)

	return cli
}
