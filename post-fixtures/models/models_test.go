package models

import (
	"blog-examples-go/post-fixtures/testing/fixtures"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Foo(t *testing.T) {
	assert.Equal(t, "bar", "baz")
	_ = fixtures.CreateUser()
}
