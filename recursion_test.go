package webpay

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RecursionCreate(t *testing.T) {
	client := NewWebPayClient(TestAuthToken)
	ret, err := client.Recursion.Create(
		400.0,
		"jpy",
		"cus_45d3MV5phaXJ4uv",
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
	client := NewWebPayClient(TestAuthToken)
	ret, err := client.Recursion.Retrieve("rec_8hla3E10n2T84Ne")
	assert.Nil(t, err)

	b, err := GetId(ret)
	assert.Nil(t, err)
	assert.NotEmpty(t, b)
	assert.True(t, strings.HasPrefix(b, "rec_"))
}
