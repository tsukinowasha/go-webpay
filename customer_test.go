package webpay

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CustomerCreate(t *testing.T) {
	client := NewWebPayClientForTesting(TestMode, TestAuthToken)
	ret, err := client.Customer.Create(
		TestCard,
	)
	assert.Nil(t, err)

	b, err := GetId(ret)
	assert.Nil(t, err)
	assert.NotEmpty(t, b)
	assert.True(t, strings.HasPrefix(b, "cus_"))
}

func Test_CustomerDelete(t *testing.T) {
	client := NewWebPayClientForTesting(TestMode, TestAuthToken)
	ret, err := client.Customer.Delete(
		TestCustomer,
	)
	assert.Nil(t, err)
	if TestMode == "real" {
		// not test
	} else {
		b, err := GetId(ret)
		assert.Nil(t, err)
		assert.NotEmpty(t, b)
		assert.True(t, strings.HasPrefix(b, "cus_"))
	}
}

func Test_CustomerAll(t *testing.T) {
	client := NewWebPayClientForTesting(TestMode, TestAuthToken)
	ret, err := client.Customer.All(map[string]int{
		"count": 5,
		"gt":    1412751347, // TODO: How to specify?
	})
	assert.Nil(t, err)

	fmt.Sprintf("%v", ret)
}
