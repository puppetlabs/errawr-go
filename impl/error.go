package impl

import (
	"fmt"

	errawr "github.com/reflect/errawr-go"
)

type ErrorDomain struct {
	Key   string
	Title string
}

type ErrorSection struct {
	Key   string
	Title string
}

type ErrorDescription struct {
	Friendly  string
	Technical string
}

type Error struct {
	Version          uint64
	ErrorDomain      *ErrorDomain
	ErrorSection     *ErrorSection
	ErrorCode        string
	ErrorTitle       string
	ErrorDescription *ErrorDescription
	ErrorArguments   ErrorArguments

	causes []errawr.Error
	buggy  bool
}

func (e Error) Domain() errawr.ErrorDomain {
	return &ErrorDomainRepr{Delegate: e.ErrorDomain}
}

func (e Error) Section() errawr.ErrorSection {
	return &ErrorSectionRepr{Delegate: e.ErrorSection}
}

func (e Error) Code() string {
	return fmt.Sprintf(`%s_%s_%s`, e.ErrorDomain.Key, e.ErrorSection.Key, e.ErrorCode)
}

func (e Error) Is(code string) bool {
	return e.Code() == code
}

func (e Error) Title() string {
	return e.ErrorTitle
}

func (e Error) Description() errawr.ErrorDescription {
	return &UnformattedErrorDescription{e.ErrorDescription}
}

func (e *Error) FormattedDescription() errawr.ErrorDescription {
	return &FormattedErrorDescription{delegate: e}
}

func (e Error) Arguments() map[string]interface{} {
	m := make(map[string]interface{})
	for k, a := range e.ErrorArguments {
		m[k] = a.Value
	}

	return m
}

func (e Error) Bug() errawr.Error {
	e.buggy = true
	return &e
}

func (e Error) IsBug() bool {
	return e.buggy
}

func (e Error) WithCause(cause errawr.Error) errawr.Error {
	if cause.IsBug() {
		e.buggy = true
	}

	e.causes = append([]errawr.Error{}, e.causes...)
	e.causes = append(e.causes, cause)
	return &e
}

func (e Error) Causes() []errawr.Error {
	return e.causes
}

func (e Error) Error() string {
	var buggy string
	if e.IsBug() {
		buggy = " (BUG)"
	}

	repr := fmt.Sprintf(`%s%s: %s`, e.Code(), buggy, e.FormattedDescription().Technical())
	for _, cause := range e.Causes() {
		repr += fmt.Sprintf("\n%s", cause.Error())
	}

	return repr
}
