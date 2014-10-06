package webpay

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CustomerCreate(t *testing.T) {
	client := NewWebPayClient(TestAuthToken)
	ret, err := client.Customer.Create(
		testCard,
	)
	assert.Nil(t, err)

	b, err := GetId(ret)
	assert.Nil(t, err)
	assert.NotEmpty(t, b)
	assert.True(t, strings.HasPrefix(b, "cus_"))
}

func Test_CustomerAll(t *testing.T) {
	client := NewWebPayClient(TestAuthToken)
	ret, err := client.Customer.All(
		3,
		0,
		map[string]int{},
	)
	assert.Nil(t, err)

	m, _ := ret.MarshalJSON()
	fmt.Println(string(m))
}
