package encoding_test

import (
	"encoding/json"
	"testing"

	"github.com/puppetlabs/errawr-go/v2/pkg/encoding"
	"github.com/puppetlabs/errawr-go/v2/pkg/errawr"
	"github.com/puppetlabs/errawr-go/v2/pkg/impl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransit(t *testing.T) {
	var ea errawr.Error = &impl.Error{
		Version: errawr.Version,
		ErrorDomain: &impl.ErrorDomain{
			Key:   "td",
			Title: "Test Domain",
		},
		ErrorSection: &impl.ErrorSection{
			Key:   "ts",
			Title: "Test Section",
		},
		ErrorCode:  "test",
		ErrorTitle: "Test Error",
		ErrorDescription: &impl.ErrorDescription{
			Friendly:  "I am friendly.",
			Technical: "I am not.",
		},
		ErrorArguments: impl.ErrorArguments{
			"test": &impl.ErrorArgument{
				Value:       true,
				Description: "an argument",
			},
		},
		ErrorItems: impl.ErrorItems{
			"a.b.c": &impl.Error{
				Version: errawr.Version,
				ErrorDomain: &impl.ErrorDomain{
					Key:   "td",
					Title: "Test Domain",
				},
				ErrorSection: &impl.ErrorSection{
					Key:   "ts",
					Title: "Test Section",
				},
				ErrorCode:  "contained",
				ErrorTitle: "Contained Error",
				ErrorDescription: &impl.ErrorDescription{
					Friendly:  "You cannot contain me!",
					Technical: "I am contained.",
				},
				ErrorMetadata: &impl.ErrorMetadata{},
			},
		},
		ErrorMetadata: &impl.ErrorMetadata{
			HTTPErrorMetadata: &impl.HTTPErrorMetadata{
				ErrorStatus:  500,
				ErrorHeaders: map[string][]string{"Test-Header": []string{"Test-Value"}},
			},
		},
	}
	ea = ea.WithCause(&impl.Error{
		Version: errawr.Version,
		ErrorDomain: &impl.ErrorDomain{
			Key:   "td",
			Title: "Test Domain",
		},
		ErrorSection: &impl.ErrorSection{
			Key:   "ts",
			Title: "Test Section",
		},
		ErrorCode:  "hello",
		ErrorTitle: "Hello Error",
		ErrorDescription: &impl.ErrorDescription{
			Friendly:  "Hello!",
			Technical: "Hi.",
		},
		ErrorMetadata: &impl.ErrorMetadata{},
	})
	ea = ea.Bug()

	env := encoding.ForTransit(ea)
	require.Equal(t, ea, env.AsError())

	b, err := json.Marshal(env)
	require.NoError(t, err)

	var ete encoding.ErrorTransitEnvelope
	require.NoError(t, json.Unmarshal(b, &ete))
	require.Equal(t, ea, ete.AsError())
}

func TestTransitConveysContainerTrait(t *testing.T) {
	tests := []struct {
		Error             errawr.Error
		HasContainerTrait bool
	}{
		{
			Error: &impl.Error{
				Version: errawr.Version,
				ErrorDomain: &impl.ErrorDomain{
					Key:   "td",
					Title: "Test Domain",
				},
				ErrorSection: &impl.ErrorSection{
					Key:   "ts",
					Title: "Test Section",
				},
				ErrorCode:  "test",
				ErrorTitle: "Test Error",
				ErrorDescription: &impl.ErrorDescription{
					Friendly:  "I am friendly.",
					Technical: "I am not.",
				},
				ErrorItems: nil,
			},
			HasContainerTrait: false,
		},
	}
	for _, test := range tests {
		b, err := json.Marshal(encoding.ForTransit(test.Error))
		require.NoError(t, err)

		var env encoding.ErrorTransitEnvelope
		require.NoError(t, json.Unmarshal(b, &env))

		_, ok := env.AsError().Items()

		if test.HasContainerTrait {
			assert.NotNil(t, env.Items)
			assert.True(t, ok, "expected error to have container trait")
		} else {
			assert.Nil(t, env.Items)
			assert.False(t, ok, "expected error not to have container trait")
		}
	}
}
