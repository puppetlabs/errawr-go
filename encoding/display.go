package encoding

import (
	errawr "github.com/puppetlabs/errawr-go"
)

type ErrorDescription struct {
	Friendly  string `json:"friendly"`
	Technical string `json:"technical"`
}

type ErrorDisplayEnvelope struct {
	Code        string                  `json:"code"`
	Title       string                  `json:"title"`
	Description *ErrorDescription       `json:"description,omitempty"`
	Arguments   map[string]interface{}  `json:"arguments,omitempty"`
	Formatted   *ErrorDescription       `json:"formatted,omitempty"`
	Causes      []*ErrorDisplayEnvelope `json:"causes,omitempty"`
}

func ForDisplay(e errawr.Error) *ErrorDisplayEnvelope {
	return ForDisplayWithSensitivity(e, errawr.ErrorSensitivityEdge)
}

func ForDisplayWithSensitivity(e errawr.Error, sensitivity errawr.ErrorSensitivity) *ErrorDisplayEnvelope {
	if e.Sensitivity() > sensitivity {
		return &ErrorDisplayEnvelope{Code: e.Code(), Title: e.Title()}
	}

	causes := e.Causes()
	ces := make([]*ErrorDisplayEnvelope, len(causes))
	for i, cause := range causes {
		ces[i] = ForDisplay(cause)
	}

	return &ErrorDisplayEnvelope{
		Code:  e.Code(),
		Title: e.Title(),
		Description: &ErrorDescription{
			Friendly:  e.Description().Friendly(),
			Technical: e.Description().Technical(),
		},
		Arguments: e.Arguments(),
		Formatted: &ErrorDescription{
			Friendly:  e.FormattedDescription().Friendly(),
			Technical: e.FormattedDescription().Technical(),
		},
		Causes: ces,
	}
}
