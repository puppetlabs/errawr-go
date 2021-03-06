package testutil_test

import (
	"testing"

	"github.com/puppetlabs/errawr-go/v2/pkg/encoding"
	"github.com/puppetlabs/errawr-go/v2/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestStub(t *testing.T) {
	stub := testutil.NewStubError("woo")
	assert.True(t, testutil.IsStubError("woo", stub))
	assert.False(t, testutil.IsStubError("nope", stub))
	assert.NotPanics(t, func() {
		encoding.ForDisplay(stub)
		encoding.ForTransit(stub)
	})
}
