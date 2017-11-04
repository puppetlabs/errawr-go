package errawr

type Error interface {
	error

	// Code is the full code for this error.
	Code() string

	// Arguments is the read-only argument map for this error.
	Arguments() map[string]interface{}

	// Bug causes this error to become a buggy error. Buggy errors are subject
	// to additional reporting.
	Bug() Error

	// IsBug returns true if this error is buggy.
	IsBug() bool

	WithCause(cause Error) Error
}
