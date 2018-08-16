package golang_test

import (
	"fmt"
	"testing"

	"github.com/puppetlabs/errawr-go/golang"
	"github.com/stretchr/testify/assert"
)

type GoTestError struct {
	i int
}

func (gte GoTestError) Error() string {
	return fmt.Sprintf("error #%d", gte.i)
}

func TestGoErrorCause(t *testing.T) {
	gerr := &GoTestError{10}
	err := golang.NewError(gerr)

	assert.True(t, err.Domain().Is("err"))
	assert.True(t, err.Section().Is("golang"))
	assert.Equal(t, "github_com_puppetlabs_errawr_go_golang_test_GoTestError", err.Code())
	assert.Equal(t, gerr.Error(), err.FormattedDescription().Friendly())
	assert.Equal(t, gerr.Error(), err.FormattedDescription().Technical())
}
