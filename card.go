package webpay

import (
	"net/url"
	"strconv"
)

type Card struct {
	number    string
	exp_month int
	exp_year  int
	cvc       string
	name      string
}

func NewCard() Card {
	ret := Card{}
	return ret
}

func (c Card) AddParams(params url.Values) url.Values {
	params.Add("card[number]", c.number)
	params.Add("card[exp_month]", strconv.Itoa(c.exp_month))
	params.Add("card[exp_year]", strconv.Itoa(c.exp_year))
	params.Add("card[cvc]", c.cvc)
	params.Add("card[name]", c.name)

	return params
}
