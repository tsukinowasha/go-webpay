package webpay

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EventRetrieve(t *testing.T) {
	client := NewWebPayClientForTesting(TestMode, TestAuthToken)
	ret, err := client.Event.Retrieve(TestEvent)
	assert.Nil(t, err)

	b, err := GetId(ret)
	assert.Nil(t, err)
	assert.NotEmpty(t, b)
	assert.Equal(t, TestEvent, b)

	m, _ := ret.MarshalJSON()
	fmt.Println(string(m))
}

func Test_EventAll(t *testing.T) {
	client := NewWebPayClientForTesting(TestMode, TestAuthToken)
	ret, err := client.Event.All(map[string]int{
		"count": 5,
		"gt":    1412751347, // TODO: How to specify?
	})
	assert.Nil(t, err)

	m, _ := ret.MarshalJSON()
	fmt.Println(string(m))
}
