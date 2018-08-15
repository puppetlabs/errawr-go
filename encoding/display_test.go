package encoding_test

import (
	"encoding/json"
	"testing"

	errawr "github.com/puppetlabs/errawr-go"
	"github.com/puppetlabs/errawr-go/encoding"
	"github.com/puppetlabs/errawr-go/impl"
	"github.com/stretchr/testify/require"
)

func TestDisplay(t *testing.T) {
	var e errawr.Error = &impl.Error{
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
			Friendly:  "Am I friendly? {{test}}",
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
	e = e.WithCause(&impl.Error{
		Version: errawr.Version,
		ErrorDomain: &impl.ErrorDomain{
			Key:   "td",
			Title: "Test Domain",
		},
		ErrorSection: &impl.ErrorSection{
			Key:   "ts",
			Title: "Test Section",
		},
		ErrorCode:  "test_cause",
		ErrorTitle: "Test Error Cause",
		ErrorDescription: &impl.ErrorDescription{
			Friendly:  "I am a cause.",
			Technical: "I am a cause.",
		},
	})

	b, err := json.Marshal(encoding.ForDisplay(e))
	require.NoError(t, err)
	require.JSONEq(t, `{
		"code": "td_ts_test",
		"title": "Test Error",
		"description": {
			"friendly": "Am I friendly? {{test}}",
			"technical": "I am not."
		},
		"arguments": {"test": true},
		"formatted": {
			"friendly": "Am I friendly? true",
			"technical": "I am not."
		},
		"causes": [
			{
				"code": "td_ts_test_cause",
				"title": "Test Error Cause",
				"description": {
					"friendly": "I am a cause.",
					"technical": "I am a cause."
				},
				"formatted": {
					"friendly": "I am a cause.",
					"technical": "I am a cause."
				}
			}
		]
	}`, string(b))

	e = e.Bug()

	b, err = json.Marshal(encoding.ForDisplay(e))
	require.NoError(t, err)
	require.JSONEq(t, `{
		"code": "td_ts_test",
		"title": "Test Error"
	}`, string(b))
}
