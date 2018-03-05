package errors

type ValidateError struct {
	Message string
}

func (v *ValidateError) AddMessage(message string) {
	v.Message += message + "\n"
}
