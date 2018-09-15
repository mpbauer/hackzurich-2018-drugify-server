package errorhandling

func NewErrNotFound(message string) *ErrNotFound {
	return &ErrNotFound{
		message: message,
	}
}
type ErrNotFound struct{
	message string
}

func (err ErrNotFound) Error() string {
	return err.message
}