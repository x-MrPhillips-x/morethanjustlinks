package presentation

type ErrInvalidUsername struct{}
type ErrInvalidEmail struct{}
type ErrInvalidPhone struct{}

func (e *ErrInvalidUsername) Error() string {
	return "please enter 3-25 characters only, no special characters or numbers"
}

func (e *ErrInvalidEmail) Error() string {
	return "please enter a valid email"
}

func (e *ErrInvalidPhone) Error() string {
	return "please enter a valid US 10-digit mobile number"
}
