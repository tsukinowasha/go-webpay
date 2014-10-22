package webpay

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RecursionCreate(t *testing.T) {
	client := NewWebPayClientForTesting(TestMode, TestAuthToken)
	ret, err := client.Recursion.Create(
		400.0,
		"jpy",
		TestCustomer,
		"month",
		"説明",
	)
	assert.Nil(t, err)

	b, err := GetId(ret)
	assert.Nil(t, err)
	assert.NotEmpty(t, b)
	assert.True(t, strings.HasPrefix(b, "rec_"))

	//	m, _ := ret.MarshalJSON()
	//fmt.Println(string(m))
}

func Test_RecursionRetrieve(t *testing.T) {
	client := NewWebPayClientForTesting(TestMode, TestAuthToken)
	ret, err := client.Recursion.Retrieve(TestRecursion)
	assert.Nil(t, err)

	b, err := GetId(ret)
	assert.Nil(t, err)
	assert.NotEmpty(t, b)
	assert.Equal(t, TestRecursion, b)
}

func Test_RecursionAll(t *testing.T) {
	client := NewWebPayClientForTesting(TestMode, TestAuthToken)
	ret, err := client.Recursion.All(map[string]int{
		"count": 5,
		"gt":    1412751347, // TODO: How to specify?
	})
	assert.Nil(t, err)

	m, err := ret.MarshalJSON()
	assert.Nil(t, err)
	assert.NotEmpty(t, m)
	//fmt.Println(string(m))
}
