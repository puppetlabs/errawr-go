// +build unit

package impl_test

import (
	"fmt"
	"testing"

	"github.com/puppetlabs/errawr-go/impl"
	"github.com/stretchr/testify/assert"
)

func TestErrorImmutability(t *testing.T) {
	err1 := &impl.Error{ErrorTitle: "Error 1"}
	err2 := err1.Bug()

	assert.NotEqual(t, err2, err1)
	assert.Equal(t, "Error 1", err1.Title())
	assert.Equal(t, "Error 1", err2.Title())
	assert.False(t, err1.IsBug())
	assert.True(t, err2.IsBug())

	cause1 := &impl.Error{ErrorTitle: "Cause 1"}
	cause2 := &impl.Error{ErrorTitle: "Cause 2"}
	err3 := err2.WithCause(cause1)
	assert.NotEqual(t, err2, err3)
	assert.Empty(t, err2.Causes())
	assert.Len(t, err3.Causes(), 1)
	assert.Equal(t, err3.Causes()[0], cause1)

	err4 := err3.WithCause(cause2)
	assert.Len(t, err3.Causes(), 1)
	assert.Equal(t, err3.Causes()[0], cause1)
	assert.Len(t, err4.Causes(), 2)
	assert.Equal(t, err4.Causes()[0], cause1)
	assert.Equal(t, err4.Causes()[1], cause2)
}

func TestErrorWithGoCause(t *testing.T) {
	err := (&impl.Error{ErrorTitle: "Error 1"}).WithCause(fmt.Errorf("oh no!"))
	assert.Len(t, err.Causes(), 1)
	assert.Equal(t, "errors_errorString", err.Causes()[0].Code())
	assert.Equal(t, "*errors.errorString", err.Causes()[0].Title())
	assert.Equal(t, "oh no!", err.Causes()[0].FormattedDescription().Friendly())
}
