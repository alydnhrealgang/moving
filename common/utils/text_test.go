package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RandomID(t *testing.T) {
	s := EmptyString
	for i := 0; i < 1000; i++ {
		token, err := RandomID()
		assert.Nil(t, err)
		assert.Len(t, token, 32)
		assert.NotEqual(t, token, s)
		s = token
		fmt.Println(s)
	}
}
