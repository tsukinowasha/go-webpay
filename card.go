package webpay

import (
	"net/url"
	"strconv"
)

type Card struct {
	Number    string
	Exp_month int
	Exp_year  int
	Cvc       string
	Name      string
}

func NewCard() Card {
	ret := Card{}
	return ret
}

func (c Card) AddParams(params url.Values) url.Values {
	params.Add("card[number]", c.Number)
	params.Add("card[exp_month]", strconv.Itoa(c.Exp_month))
	params.Add("card[exp_year]", strconv.Itoa(c.Exp_year))
	params.Add("card[cvc]", c.Cvc)
	params.Add("card[name]", c.Name)

	return params
}
