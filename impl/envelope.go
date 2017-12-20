package impl

import errawr "github.com/reflect/errawr-go"

type ErrorEnvelope struct {
	Version     uint64            `json:"version"`
	Domain      *ErrorDomain      `json:"domain"`
	Section     *ErrorSection     `json:"section"`
	Code        string            `json:"code"`
	Title       string            `json:"title"`
	Description *ErrorDescription `json:"description"`
	Arguments   ErrorArguments    `json:"arguments"`
	Metadata    *ErrorMetadata    `json:"metadata,omitempty"`
	Causes      []*ErrorEnvelope  `json:"causes"`
	Buggy       bool              `json:"buggy"`
}

func (ee ErrorEnvelope) AsError() *Error {
	var causes []errawr.Error
	for _, cause := range ee.Causes {
		causes = append(causes, cause.AsError())
	}

	return &Error{
		Version:          ee.Version,
		ErrorDomain:      ee.Domain,
		ErrorSection:     ee.Section,
		ErrorCode:        ee.Code,
		ErrorTitle:       ee.Title,
		ErrorDescription: ee.Description,
		ErrorArguments:   ee.Arguments,
		ErrorMetadata:    ee.Metadata,

		causes: causes,
		buggy:  ee.Buggy,
	}
}

func NewErrorEnvelope(e *Error) *ErrorEnvelope {
	causes := e.Causes()
	ces := make([]*ErrorEnvelope, len(causes))
	for i, cause := range causes {
		ces[i] = NewErrorEnvelope(Copy(cause))
	}

	return &ErrorEnvelope{
		Version:     e.Version,
		Domain:      e.ErrorDomain,
		Section:     e.ErrorSection,
		Code:        e.ErrorCode,
		Title:       e.ErrorTitle,
		Description: e.ErrorDescription,
		Arguments:   e.ErrorArguments,
		Metadata:    e.ErrorMetadata,
		Causes:      ces,
		Buggy:       e.buggy,
	}
}
