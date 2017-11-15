// +build unit

package impl_test

import (
	"testing"

	"github.com/reflect/errawr-go/impl"
	"github.com/stretchr/testify/assert"
)

func TestDescriptionFormattingSubstitution(t *testing.T) {
	desc := &impl.ErrorDescription{
		Friendly:  "An unexpected error occurred in the PostgreSQL driver.",
		Technical: "An unexpected error occurred in the PostgreSQL driver: {{cause}}.",
	}

	args := impl.ErrorArguments{
		"cause": impl.NewErrorArgument("OH NO EVERYTHING IS WRONG", "the error that occurred"),
	}

	err := &impl.Error{
		ErrorArguments:   args,
		ErrorDescription: desc,
	}

	fd := err.FormattedDescription()
	assert.Equal(t, "An unexpected error occurred in the PostgreSQL driver.", fd.Friendly())
	assert.Equal(t, "An unexpected error occurred in the PostgreSQL driver: OH NO EVERYTHING IS WRONG.", fd.Technical())
}

func TestDescriptionFormattingBlocks(t *testing.T) {
	desc := &impl.ErrorDescription{
		Friendly: "An error occurred while attempting to execute a database command.",
		Technical: `PostgreSQL error {{code}} ({{pre name}}) occurred: {{message}}
{{#if detail}}

Detail: {{detail}}
{{/if}}
{{#if hint}}

Hint: {{hint}}
{{/if}}`,
	}

	args := impl.ErrorArguments{
		"code":    impl.NewErrorArgument("40904", "the error code"),
		"name":    impl.NewErrorArgument("unexpected_fun", "the human-readable name of the error code"),
		"message": impl.NewErrorArgument("Shut it down!", "the terse message for this error"),
		"detail":  impl.NewErrorArgument("Too much fun was had in postgres.c:420.", "an optional detailed description for this error"),
		"hint":    impl.NewErrorArgument("", "an optional hint on how to solve this error"),
	}

	err := &impl.Error{
		ErrorArguments:   args,
		ErrorDescription: desc,
	}

	expected := `PostgreSQL error 40904 (` + "`unexpected_fun`" + `) occurred: Shut it down!

Detail: Too much fun was had in postgres.c:420.`

	assert.Equal(t, expected, err.FormattedDescription().Technical())
}

func TestDescriptionFormattingHelpers(t *testing.T) {
	desc := &impl.ErrorDescription{
		Friendly:  "We couldn't find the columns {{#join columns}}{{pre this}}{{/join}}.",
		Technical: "The field {{quote field}} references nonexistent columns {{#join columns}}{{pre this}}{{/join}}.",
	}

	args := impl.ErrorArguments{
		"field":   impl.NewErrorArgument("Field 1", "the field that references the nonexistent columns"),
		"columns": impl.NewErrorArgument([]string{"col1", "col2", "col3"}, "the columns that could not be found"),
	}

	err := &impl.Error{
		ErrorArguments:   args,
		ErrorDescription: desc,
	}

	fd := err.FormattedDescription()
	assert.Equal(t, "We couldn't find the columns `col1`, `col2`, and `col3`.", fd.Friendly())
	assert.Equal(t, "The field \"Field 1\" references nonexistent columns `col1`, `col2`, and `col3`.", fd.Technical())
}
