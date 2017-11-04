package errawr

type ErrorArgument struct {
	Value       string
	Description string
}

func (ea *ErrorArgument) Set(value string) {
	ea.Value = value
}

func (ea *ErrorArgument) Validate(validator string) {

}

func NewErrorArgument(value, description string) *ErrorArgument {
	return &ErrorArgument{
		Value:       value,
		Description: description,
	}
}

type ErrorArguments map[string]*ErrorArgument
