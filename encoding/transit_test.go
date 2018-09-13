package encoding_test

import (
	"encoding/json"
	"testing"

	errawr "github.com/puppetlabs/errawr-go"
	"github.com/puppetlabs/errawr-go/encoding"
	"github.com/puppetlabs/errawr-go/impl"
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
