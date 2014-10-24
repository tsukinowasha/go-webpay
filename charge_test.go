package webpay

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ChargeCreate(t *testing.T) {
	client := NewWebPayClientForTesting(TestMode, TestAuthToken)
	ret, err := client.Charge.Create(
		400.0,
		"jpy",
		TestCard,
	)
	assert.Nil(t, err)

	b, err := GetId(ret)
	assert.Nil(t, err)
	assert.NotEmpty(t, b)
	assert.True(t, strings.HasPrefix(b, "ch_"))
}

func Test_ChargeRetrieve(t *testing.T) {
	client := NewWebPayClientForTesting(TestMode, TestAuthToken)
	ret, err := client.Charge.Retrieve(TestCharge)
	assert.Nil(t, err)

	b, err := GetId(ret)
	assert.Nil(t, err)
	assert.NotEmpty(t, b)
	assert.Equal(t, TestCharge, b)

	//	m, _ := ret.MarshalJSON()
	// fmt.Println(string(m))
}

func Test_ChargeCreateByCustomer(t *testing.T) {
	client := NewWebPayClientForTesting(TestMode, TestAuthToken)
	ret, err := client.Charge.CreateByCustomer(
		400.0,
		"jpy",
		TestCustomer,
	)
	assert.Nil(t, err)

	b, err := GetId(ret)
	assert.Nil(t, err)
	assert.NotEmpty(t, b)
	assert.True(t, strings.HasPrefix(b, "ch_"))
}

func Test_ChargeAll(t *testing.T) {
	client := NewWebPayClientForTesting(TestMode, TestAuthToken)
	ret, err := client.Charge.All(map[string]int{
		"count": 5,
		"gt":    1412751347, // TODO: How to specify?
	},
		"",
		"",
	)
	assert.Nil(t, err)

	m, err := ret.MarshalJSON()
	assert.Nil(t, err)
	assert.NotEmpty(t, m)
	//fmt.Println(string(m))
}
