package errawr

import "fmt"

type ErrorDomain struct {
	Key   string
	Title string
}

type ErrorSection struct {
	Key   string
	Title string
}

type Error interface {
	error

	WithCause(cause Error) Error
}

type ErrorDescription struct {
	Friendly  string
	Technical string
}

type GeneralError struct {
	Domain      *ErrorDomain
	Section     *ErrorSection
	Code        string
	Title       string
	Description *ErrorDescription
	Arguments   ErrorArguments
	Causes      []Error
}

func (ge *GeneralError) WithCause(cause Error) Error {
	ge.Causes = append(ge.Causes, cause)
	return ge
}

func (ge *GeneralError) Error() string {
	return fmt.Sprintf(`%s_%s_%s: %s`, ge.Domain, ge.Section, ge.Code, ge.Description)
}
