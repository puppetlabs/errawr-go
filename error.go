package errawr

// ErrorDescription is the accessor type for error descriptions in different
// states.
type ErrorDescription interface {
	// Friendly is an end-user-friendly description of an error.
	Friendly() string

	// Technical is a description of an error suitable for sending to a support
	// person.
	Technical() string
}

// Error is the type of all user-facing errors.
type Error interface {
	error

	// Code is the full code for this error.
	Code() string

	// Title is the short title for this error.
	Title() string

	// Description returns the unformatted descriptions of this error.
	Description() ErrorDescription

	// FormattedDescription returns an ASCII-printable formatted description
	// of this error.
	FormattedDescription() ErrorDescription

	// Arguments is the read-only argument map for this error.
	Arguments() map[string]interface{}

	// Bug causes this error to become a buggy error. Buggy errors are subject
	// to additional reporting.
	Bug() Error

	// IsBug returns true if this error is buggy.
	IsBug() bool

	// WithCause causes this error to be caused by the given error. If it is
	// already caused by another error, it will be caused by both errors.
	WithCause(cause Error) Error

	// Causes returns the list of causes for this error.
	Causes() []Error
}
