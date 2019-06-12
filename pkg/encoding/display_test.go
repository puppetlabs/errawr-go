package encoding_test

import (
	"encoding/json"
	"testing"

	"github.com/puppetlabs/errawr-go/pkg/encoding"
	"github.com/puppetlabs/errawr-go/pkg/errawr"
	"github.com/puppetlabs/errawr-go/pkg/impl"
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
		"domain": "td",
		"section": "ts",
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
				"domain": "td",
				"section": "ts",
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

	var ede encoding.ErrorDisplayEnvelope
	require.NoError(t, json.Unmarshal(b, &ede))

	var expected errawr.Error = &impl.Error{
		Version: errawr.Version,
		ErrorDomain: &impl.ErrorDomain{
			Key: e.Domain().Key(),
		},
		ErrorSection: &impl.ErrorSection{
			Key: e.Section().Key(),
		},
		ErrorCode:  e.Code(),
		ErrorTitle: e.Title(),
		ErrorDescription: &impl.ErrorDescription{
			Friendly:  e.Description().Friendly(),
			Technical: e.Description().Technical(),
		},
		ErrorArguments: impl.ErrorArguments{
			"test": &impl.ErrorArgument{Value: true},
		},
		ErrorMetadata:    &impl.ErrorMetadata{},
		ErrorSensitivity: errawr.ErrorSensitivityEdge,
	}
	expected = expected.WithCause(&impl.Error{
		Version: errawr.Version,
		ErrorDomain: &impl.ErrorDomain{
			Key: e.Causes()[0].Domain().Key(),
		},
		ErrorSection: &impl.ErrorSection{
			Key: e.Causes()[0].Section().Key(),
		},
		ErrorCode:  e.Causes()[0].Code(),
		ErrorTitle: e.Causes()[0].Title(),
		ErrorDescription: &impl.ErrorDescription{
			Friendly:  e.Causes()[0].Description().Friendly(),
			Technical: e.Causes()[0].Description().Technical(),
		},
		ErrorArguments:   impl.ErrorArguments{},
		ErrorMetadata:    &impl.ErrorMetadata{},
		ErrorSensitivity: errawr.ErrorSensitivityEdge,
	})

	require.Equal(t, expected, ede.AsError())

	e = e.Bug()

	b, err = json.Marshal(encoding.ForDisplay(e))
	require.NoError(t, err)
	require.JSONEq(t, `{
		"domain": "td",
		"section": "ts",
		"code": "td_ts_test",
		"title": "Test Error"
	}`, string(b))
}
