package webpay

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AccountRetrieve(t *testing.T) {
	client := NewWebPayClientForTesting(TestMode, TestAuthToken)
	ret, err := client.Account.Retrieve()
	assert.Nil(t, err)

	b, err := GetId(ret)
	assert.Nil(t, err)
	assert.NotEmpty(t, b)
	assert.True(t, strings.HasPrefix(b, "acct_"))
}
