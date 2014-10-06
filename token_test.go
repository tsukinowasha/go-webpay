package webpay

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TokenCreate(t *testing.T) {
	client := NewWebPayClient(TestAuthToken)
	ret, err := client.Token.Create(
		"", // If UUID can be empty
		testCard,
	)
	assert.Nil(t, err)

	b, err := GetId(ret)
	assert.Nil(t, err)
	assert.NotEmpty(t, b)
	assert.True(t, strings.HasPrefix(b, "tok_"))
}

func Test_TokenRetrieve(t *testing.T) {
	client := NewWebPayClient(TestAuthToken)
	ret, err := client.Token.Retrieve("tok_3ybc93ckR01qeKx")
	assert.Nil(t, err)

	b, err := GetId(ret)
	assert.Nil(t, err)
	assert.NotEmpty(t, b)
	assert.True(t, strings.HasPrefix(b, "tok_"))
}
