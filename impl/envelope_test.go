package impl_test

import (
	"encoding/json"
	"testing"

	errawr "github.com/puppetlabs/errawr-go"
	"github.com/puppetlabs/errawr-go/impl"
	"github.com/stretchr/testify/require"
)

func TestErrorJSONEncodingTransitivity(t *testing.T) {
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
				ErrorStatus: 500,
			},
		},
	}
	ea = ea.WithCause(&impl.Error{ErrorTitle: "hello"})
	ea = ea.Bug()

	b, err := json.Marshal(ea)
	require.NoError(t, err)

	var ec impl.Error
	require.NoError(t, json.Unmarshal(b, &ec))
	require.Equal(t, ea, &ec)
}
