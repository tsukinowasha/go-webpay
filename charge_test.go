package webpay

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ChargeCreate(t *testing.T) {
	client := NewWebPayClient(TestAuthToken)
	ret, err := client.Charge.Create(
		400.0,
		"jpy",
		testCard,
	)
	assert.Nil(t, err)

	b, err := GetId(ret)
	assert.Nil(t, err)
	assert.NotEmpty(t, b)
	assert.True(t, strings.HasPrefix(b, "ch_"))
}

func Test_ChargeRetrieve(t *testing.T) {
	client := NewWebPayClient(TestAuthToken)
	ret, err := client.Charge.Retrieve("ch_3u32gEgiy2xjfEf")
	assert.Nil(t, err)

	b, err := GetId(ret)
	assert.Nil(t, err)
	assert.NotEmpty(t, b)
	assert.True(t, strings.HasPrefix(b, "ch_"))

	//	m, _ := ret.MarshalJSON()
	// fmt.Println(string(m))
}
